package pkg

import (
	"github.com/luongdev/gotasker"
	"github.com/luongdev/gotasker/internal"
)

func MockTask(name string) gotasker.Task {
	return internal.MockTask(name)
}

func NewInMemoryTaskStore(tasks ...gotasker.Task) gotasker.TaskStore {
	return internal.NewInMemoryTaskStore(tasks...)
}

func MockPlan(name string) gotasker.TaskPlan {
	return internal.MockPlan(name)
}

func NewInMemoryPlanStore(plans ...gotasker.TaskPlan) gotasker.PlanStore {
	return internal.NewInMemoryPlanStore(plans...)
}
