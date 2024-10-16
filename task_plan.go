package gotasker

import (
	"github.com/google/uuid"
	"time"
)

type TaskPlan interface {
	Id() uuid.UUID
	Name() string
	Args() map[string]interface{}
	NextRun() (time.Time, error)
}
