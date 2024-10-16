package internal

import (
	"github.com/google/uuid"
	"github.com/luongdev/gotasker"
	"time"
)

type taskPlan struct {
	id               uuid.UUID
	name             string
	args             map[string]interface{}
	lastRun          time.Time
	nextRun          time.Time
	startTime        time.Time
	endTime          time.Time
	startImmediately bool
}

func (t *taskPlan) Args() map[string]interface{} {
	return t.args
}

func (t *taskPlan) Id() uuid.UUID {
	return t.id
}

func (t *taskPlan) NextRun() (time.Time, error) {
	return t.nextRun, nil
}

func (t *taskPlan) Name() string {
	return t.name
}

var _ gotasker.TaskPlan = (*taskPlan)(nil)
