package internal

import (
	"github.com/luongdev/gotasker"
	"log"
)

type inMemoryTaskStore struct {
	tasks map[string]gotasker.Task
}

func (s *inMemoryTaskStore) GetByName(name string) (gotasker.Task, error) {
	return s.tasks[name], nil
}

func (s *inMemoryTaskStore) Tasks() map[string]gotasker.Task {
	return s.tasks
}

func MockTask(name string) gotasker.Task {
	return &task{
		name: name,
		taskFn: func(args map[string]any) {
			log.Printf("task %s is running", name)
		},
		onError: func(err error) {
			log.Printf("task %s has error: %s", name, err.Error())
		},
		onEnded: func() {
			log.Printf("task %s has ended", name)
		},
	}
}

func NewInMemoryTaskStore(tasks ...gotasker.Task) gotasker.TaskStore {

	store := &inMemoryTaskStore{
		tasks: make(map[string]gotasker.Task),
	}

	for _, t := range tasks {
		store.tasks[t.Name()] = t
	}

	return store
}

var _ gotasker.TaskStore = (*inMemoryTaskStore)(nil)
