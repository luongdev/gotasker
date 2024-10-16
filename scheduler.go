package gotasker

type Scheduler interface {
	Tasks() map[string]Task
	NewTask(task Task) error
	Start()
	Shutdown() error
}
