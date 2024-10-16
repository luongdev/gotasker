package internal

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/luongdev/gotasker"
	"time"
)

type task struct {
	name string

	planId uuid.UUID

	ctx      context.Context
	cancelFn context.CancelFunc

	status gotasker.TaskStatus

	taskFn  func(args map[string]any)
	onError func(err error)
	onEnded func()

	timeout time.Duration

	locker gotasker.Locker
}

func (t *task) Start(args map[string]interface{}) {
	var lock gotasker.Lock = nil
	var err error
	if t.locker != nil {
		lock, err = t.locker.Lock(t.ctx, t.name)
		if err != nil {
			t.onError(fmt.Errorf("failed to lock task %s: %w", t.name, err))
			return
		}
	}

	if t.timeout <= 0 {
		t.timeout = time.Second * 10
	}

	time.AfterFunc(t.timeout, func() {
		t.cancelFn()
		if lock != nil {
			_ = lock.Unlock(t.ctx)
		}
	})

	t.taskFn(args)
	if lock != nil {
		_ = lock.Unlock(t.ctx)
	}
}

func (t *task) Status() gotasker.TaskStatus {
	return t.status
}

func (t *task) Name() string {
	return t.name
}

var _ gotasker.Task = (*task)(nil)
