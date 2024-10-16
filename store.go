package gotasker

type TaskStore interface {
	Tasks() map[string]Task

	GetByName(string) (Task, error)
}

type PlanStore interface {
	Plans() map[string]TaskPlan
}
