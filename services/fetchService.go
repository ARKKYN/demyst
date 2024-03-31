package services

import (
	"demyst/responseSchema"
	"fmt"
	"io"
	"sync"
)

type TodoClient interface {
	FetchTodo(todoID int) (*responseSchema.TodoResponse, error)
}

type FetchService struct {
	out        io.Writer
	todoClient TodoClient
}

func NewFetchService(out io.Writer, todoClient TodoClient) *FetchService {
	return &FetchService{
		out:        out,
		todoClient: todoClient,
	}
}

func (s *FetchService) FetchTodos() error {
	todoIds := s.generateEvenTodoIds()

	var wg sync.WaitGroup
	resultsChannel := make(chan string)

	for id := range todoIds {
		wg.Add(1)
		go s.worker(id, &wg, resultsChannel)
	}

	go func() {
		wg.Wait()
		close(resultsChannel)
	}()
	
	s.printTodos(resultsChannel)
	
	return nil
}

func (s *FetchService) printTodos(resultsChannel chan string) { 
	for result := range resultsChannel {
		fmt.Fprintln(s.out, result)
	}
} 


func (s *FetchService) worker(todoID int, wg *sync.WaitGroup, resultsChannel chan<- string) {
	defer wg.Done()
	todo, err := s.todoClient.FetchTodo(todoID)
	if err != nil {
		resultsChannel <- fmt.Sprintf("Error fetching todo %d: %v", todoID, err)
		return
	}
	resultsChannel <- fmt.Sprintf("Todo: %s, Has Completed: %t", todo.Title, todo.Completed)
}

func (s *FetchService) generateEvenTodoIds() <-chan int {
	todoIds := make(chan int)
	go func() {
		defer close(todoIds)
		for p := 2; p < 41; p = p + 2 {
			todoIds <- p
		}
	}()
	return todoIds
}
