
package clients

import (
	"demyst/config"
	rs "demyst/responseSchema"
	u "demyst/utilities"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TodoClient struct {
	HTTPClient *u.CustomHttpClient
	BaseURL    string
}

var  clientConfig = &http.Client{
	Timeout: 2 * time.Second,
}
func NewTodoClient() *TodoClient {
	return &TodoClient{
		HTTPClient: u.NewCustomHttpClient(clientConfig),
		BaseURL:    config.TODO_BASE_URL,
	}
}

func (tc *TodoClient) FetchTodo(todoID int) (*rs.TodoResponse, error) {
	todoPath := tc.BaseURL + "/todos"
	url := fmt.Sprintf("%s/%d", todoPath, todoID)
	respBody, err := tc.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}

	var todo rs.TodoResponse
	if err := json.Unmarshal(respBody, &todo); err != nil {
		return nil, err
	}

	return &todo, nil
}
