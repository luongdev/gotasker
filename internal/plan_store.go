package internal

import (
	"github.com/google/uuid"
	"github.com/luongdev/gotasker"
)

type inMemoryPlanStore struct {
	taskPlans map[uuid.UUID]*taskPlan
}

func (i *inMemoryPlanStore) Plans() map[string]gotasker.TaskPlan {
	plans := make(map[string]gotasker.TaskPlan)

	for _, v := range i.taskPlans {
		plans[v.Name()] = v
	}

	return plans
}

func MockPlan(name string) gotasker.TaskPlan {
	return &taskPlan{
		name: name,
		id:   uuid.New(),
	}
}

func NewInMemoryPlanStore(plans ...gotasker.TaskPlan) gotasker.PlanStore {

	store := &inMemoryPlanStore{
		taskPlans: make(map[uuid.UUID]*taskPlan),
	}

	for _, t := range plans {
		if internalPlan, ok := t.(*taskPlan); ok {
			store.taskPlans[t.Id()] = internalPlan
		}

	}

	return store
}

var _ gotasker.PlanStore = (*inMemoryPlanStore)(nil)
