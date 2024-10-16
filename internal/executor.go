package internal

import (
	"fmt"
	"github.com/luongdev/gotasker"
)

type executor struct {
	taskStore gotasker.TaskStore
}

func (e *executor) Queue(plan gotasker.TaskPlan) error {
	t, err := e.taskStore.GetByName(plan.Name())
	if err != nil {
		return err
	}

	internalTask, ok := t.(*task)
	if !ok {
		return fmt.Errorf("task %s is not a valid task", plan.Name())
	}
	internalTask.planId = plan.Id()

	e.Execute(internalTask, plan.Args())

	return nil
}

func (e *executor) Execute(t gotasker.Task, args map[string]interface{}) {
	go func() {
		t.Start(args)
		//internalTask, ok := t.(*task)
		//if ok {
		//
		//}
	}()
}

var _ gotasker.Executor = (*executor)(nil)
