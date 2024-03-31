package clients

import (
	"demyst/responseSchema"
	u "demyst/utilities"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)


func createMockClient(url string) TodoClient {
	client := TodoClient{
		HTTPClient: &u.CustomHttpClient{
			Client: &http.Client{},
			},
			BaseURL: url,
			}
	return client			
}

func TestFetchTodoWith200Response(t *testing.T) {
	mockResponse := responseSchema.TodoResponse{
		UserID:    1,
		ID:        2,
		Title:     "quis ut nam facilis et officia qui",
		Completed: false,
	}

	server := u.CreateMockServer(http.StatusOK, mockResponse)
	defer server.Close()

	todoClient := createMockClient(server.URL)

	todo, err := todoClient.FetchTodo(2)
	assert.NoError(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, mockResponse, *todo)
}

func TestFetchTodoWith4xxResponse(t *testing.T) {
	server :=u.CreateMockServer(http.StatusNotFound, nil)
	defer server.Close()

	todoClient := createMockClient(server.URL)

	todo, err := todoClient.FetchTodo(2)
	assert.Error(t, err)
	assert.Nil(t, todo)
}

func TestFetchTodoWith5xxResponse(t *testing.T) {
	server :=u.CreateMockServer(http.StatusInternalServerError, nil)
	defer server.Close()
	
	todoClient := createMockClient(server.URL)

	todo, err := todoClient.FetchTodo(2)
	assert.Error(t, err)
	assert.Nil(t, todo)
}
