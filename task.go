package gotasker

type TaskStatus uint

const (
	StatusPending TaskStatus = 0
	StatusRunning TaskStatus = 1
	StatusDone    TaskStatus = 2
	StatusError   TaskStatus = 3
	StatusTimeout TaskStatus = 4
)

type Task interface {
	Name() string
	Status() TaskStatus
	Start(args map[string]interface{})
}
