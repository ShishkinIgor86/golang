// test.go
package v1

import (
"net/http"
"net/http/httptest"
"strings"
"testing"
)

func TestPostHandler(t *testing.T) {
	reqBody := strings.NewReader({"key": "value"})
	req, err := http.NewRequest("POST", "/post", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PostHandler)

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	expected := `{"message": "Post request successful"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body to be %s, but got %s", expected, recorder.Body.String())
	}
}

func TestDeleteHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delete", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteHandler)

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	expected := `{"message": "Delete request successful"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body to be %s, but got %s", expected, recorder.Body.String())
	}
}