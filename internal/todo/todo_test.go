package todo_test

import (
	"github.com/nathanfabio/api-usingGoland/internal/todo"
	"reflect"
	"testing"
)

func TestService_Search(t *testing.T) {
	type fields struct {
		todos []todo.Item
	}
	type args struct {
		query string
	}
	tests := []struct {
		name      string
		todoToAdd []string
		query     string
		expect    []string
	}{
		{
			name:      "I've put in a part of the word and I need a whole word that matches it",
			todoToAdd: []string{"supermarket"},
			query:     "market",
			expect:    []string{"supermarket"},
		},
		{
			name:      "still returns the word, even if doesnt match",
			todoToAdd: []string{"SuperMarket"},
			query:     "market",
			expect:    []string{"SuperMarket"},
		},
		{
			name:      "space",
			todoToAdd: []string{"go Supermarket"},
			query:     "go",
			expect:    []string{"go Supermarket"},
		},
		{
			name:      "space at start of word",
			todoToAdd: []string{" Some space"},
			query:     "space",
			expect:    []string{" Some space"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := todo.NewService()
			for _, toAdd := range tt.todoToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Fatal(err)
				}
			}
			if got := svc.Search(tt.query); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("Search() = %v, want %v", got, tt.expect)
			}
		})
	}
}
