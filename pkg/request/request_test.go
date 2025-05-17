package request

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestData struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func TestDecodeSuccess(t *testing.T) {
	jsonStr := `{"name":"John Doe","email":"john.doe@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(jsonStr))

	data, err := Decode[TestData](req)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}

	if data.Name != "John Doe" {
		t.Errorf("Expected Name to be 'John Doe', got '%s'", data.Name)
	}

	if data.Email != "john.doe@example.com" {
		t.Errorf("Expected Email to be 'john.doe@example.com', got '%s'", data.Email)
	}
}

func TestDecodeInvalidJSON(t *testing.T) {
	invalidJSON := `{"name": "Jane Doe", "email": "not-an-email"`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(invalidJSON))

	_, err := Decode[TestData](req)
	if err == nil {
		t.Fatal("Expected error due to invalid JSON but got nil")
	}
}

func TestValidateSuccess(t *testing.T) {
	valid := TestData{Name: "Alice", Email: "alice@example.com"}

	err := Validate(valid)
	if err != nil {
		t.Fatalf("Validate returned error for valid data: %v", err)
	}
}

func TestValidateFailure(t *testing.T) {
	invalid := TestData{Name: "", Email: "invalid-email"}

	err := Validate(invalid)
	if err == nil {
		t.Fatal("Expected validation error but got nil")
	}
}

func TestHandleBodySuccess(t *testing.T) {
	jsonStr := `{"name":"Bob","email":"bob@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(jsonStr))

	result, err := HandleBody[TestData](req)
	if err != nil {
		t.Fatalf("HandleBody returned error: %v", err)
	}

	if result.Name != "Bob" {
		t.Errorf("Expected Name to be 'Bob', got '%s'", result.Name)
	}
	if result.Email != "bob@example.com" {
		t.Errorf("Expected Email to be 'bob@example.com', got '%s'", result.Email)
	}
}

func TestHandleBodyInvalidJSON(t *testing.T) {
	invalidJSON := `{"name":"Eve","email":"eve@example.com"`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(invalidJSON))

	_, err := HandleBody[TestData](req)
	if err == nil {
		t.Fatal("Expected error due to invalid JSON but got nil")
	}
}

func TestHandleBodyValidationError(t *testing.T) {
	jsonStr := `{"name":"","email":"not-an-email"}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(jsonStr))

	_, err := HandleBody[TestData](req)
	if err == nil {
		t.Fatal("Expected validation error but got nil")
	}
}
