package link_updater

import (
	"context"
)

func New(repository repository, consumer amqpConsumer) *Story {
	return &Story{repository: repository, consumer: consumer}
}

type Story struct {
	repository repository
	consumer   amqpConsumer
}

func (s *Story) Run(ctx context.Context) error {
	messages, err := s.consumer.Consume(ctx, "link_queue")
	if err != nil {
		return err
	}

	for {
		select {
		case msg := <-messages:
			var m message
			err := json.Unmarshal(msg.Body, &m)
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			link, err := s.repository.GetLinkByID(m.ID)
			if err != nil {
				log.Printf("Error getting link from repository: %v", err)
				continue
			}

			newTitle, err := scrape.GetTitle(link.URL)
			if err != nil {
				log.Printf("Error getting title from scrape: %v", err)
			} else {
				link.Title = newTitle
			}

			err = s.repository.UpdateLink(link)
			if err != nil {
				log.Printf("Error updating link in repository: %v", err)
				continue
			}

			err = msg.Ack()
			if err != nil {
				log.Printf("Error acknowledging message: %v", err)
			}

		case <-ctx.Done():
			return nil
		}
	}
}
