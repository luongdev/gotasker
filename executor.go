package gotasker

type Executor interface {
	Execute(Task, map[string]interface{})
	Queue(TaskPlan) error
}
