package services

import (
	"bytes"
	"demyst/responseSchema"
	"errors"
	"strings"

	"testing"
)

// MockTodoClient is a mock implementation of TodoClient interface
type MockTodoSuccessClient struct {}

func (m *MockTodoSuccessClient) FetchTodo(todoID int) (*responseSchema.TodoResponse, error) {
	switch todoID {
	case 2:
		return &responseSchema.TodoResponse{UserID: 1, ID: 2, Title: "Mock Todo 2", Completed: true}, nil
	case 4:
		return &responseSchema.TodoResponse{UserID: 1, ID: 4, Title: "Mock Todo 4", Completed: false}, nil
	default:
		return nil, errors.New("todo not found")
	}
}


func TestFetchTodos_Success(t *testing.T) {
	var buf bytes.Buffer
	todoClient := &MockTodoSuccessClient{}
	fetchService := NewFetchService(&buf, todoClient)

	err := fetchService.FetchTodos()
	if err != nil {
		t.Errorf("FetchTodos returned an unexpected error: %v", err)
	}

	expectedSuccessOutput := "Todo: Mock Todo 2, Has Completed: true\n"
	expectedErrorOutput := "Error fetching todo 8: todo not found"
	if !strings.Contains(buf.String(), expectedSuccessOutput) {
		t.Errorf("FetchTodos returned unexpected output. Expected: %q, Got: %q", expectedSuccessOutput, buf.String())
	}
	
	if !strings.Contains(buf.String(), expectedErrorOutput) {
		t.Errorf("FetchTodos returned unexpected output. Expected: %q, Got: %q", expectedErrorOutput, buf.String())
	}
}

