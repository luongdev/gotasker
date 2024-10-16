package pkg

import (
	"github.com/luongdev/gotasker"
	"github.com/luongdev/gotasker/internal"
)

type SchedulerOption = internal.SchedulerOption

func NewScheduler(opts ...SchedulerOption) (gotasker.Scheduler, error) {
	return internal.NewScheduler(opts...)
}

func WithTaskStore(store gotasker.TaskStore) SchedulerOption {
	return internal.WithTaskStore(store)
}

func WithPlanStore(store gotasker.PlanStore) SchedulerOption {
	return internal.WithPlanStore(store)
}
