package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
    // Вставьте ваш код здесь
	var letterNumber int
	var newLowerString string
	for _, l := range text{
        if unicode.IsLetter(l) {
			letterNumber++
			newLowerString += string(unicode.ToLower(l))
		}
	}
    var perOneLet float32 = float32(letterNumber) / 100
	var countRune int
	runeOne := make(map[rune] int)
	for _, l :=range newLowerString{
		for _, r := range newLowerString{
			if l == r{
				countRune++
			}
		}
		runeOne[l] = countRune
		countRune = 0
	}
	for key, value := range runeOne {
        fmt.Printf("%c - %d %0.2f\n", key,value,float32(value)*perOneLet)
    }
}

