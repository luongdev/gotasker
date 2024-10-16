package internal

import (
	"context"
	"github.com/google/uuid"
	"github.com/luongdev/gotasker"
	"sync"
	"time"
)

type SchedulerOption func(*scheduler) error

func WithTaskStore(store gotasker.TaskStore) SchedulerOption {
	return func(s *scheduler) error {
		s.taskStore = store
		return nil
	}
}

func WithPlanStore(store gotasker.PlanStore) SchedulerOption {
	return func(s *scheduler) error {
		s.planStore = store
		return nil
	}
}

type scheduler struct {
	shutdownCtx    context.Context
	shutdownCancel context.CancelFunc

	executor executor
	logger   gotasker.Logger

	inProgressPlans map[uuid.UUID]taskPlan
	mu              sync.RWMutex

	queuedChan    chan *taskPlan
	completedChan chan struct{}
	failedChan    chan struct{}

	running     bool
	startedChan chan struct{}
	stopChan    chan struct{}

	taskStore gotasker.TaskStore
	planStore gotasker.PlanStore
}

func (s *scheduler) NewTask(task gotasker.Task) error {
	//TODO implement me
	panic("implement me")
}

func (s *scheduler) Start() {
	s.startSelector()
}

func (s *scheduler) Shutdown() error {
	//TODO implement me
	panic("implement me")
}

func (s *scheduler) Tasks() map[string]gotasker.Task {
	//TODO implement me
	panic("implement me")
}

func (s *scheduler) startSelector() {
	s.running = true

	plans := s.planStore.Plans()
	for _, plan := range plans {
		if s.isProgressing(plan.Id()) {
			continue
		}
		tPlan, ok := plan.(*taskPlan)
		if !ok {
			continue
		}

		if tPlan.startImmediately {
			_ = s.executor.Queue(tPlan)
			continue
		}

		if !tPlan.nextRun.IsZero() && tPlan.nextRun.Before(time.Now()) {
			s.queuedChan <- tPlan
			continue
		}

		if tPlan.startTime.Before(time.Now()) {
			s.queuedChan <- tPlan
			continue
		}
	}
}

func (s *scheduler) stopScheduler() {

}

func (s *scheduler) isProgressing(planId uuid.UUID) bool {
	if planId == uuid.Nil {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.inProgressPlans[planId]
	return ok
}

func NewScheduler(opts ...SchedulerOption) (gotasker.Scheduler, error) {
	ctx, cancel := context.WithCancel(context.Background())
	s := &scheduler{
		shutdownCtx:    ctx,
		shutdownCancel: cancel,

		mu:              sync.RWMutex{},
		inProgressPlans: make(map[uuid.UUID]taskPlan),

		queuedChan:    make(chan *taskPlan),
		startedChan:   make(chan struct{}),
		completedChan: make(chan struct{}),
		failedChan:    make(chan struct{}),
		stopChan:      make(chan struct{}),

		logger:   &noOpLogger{},
		executor: executor{},
	}

	for _, optionFn := range opts {
		if err := optionFn(s); err != nil {
			return nil, err
		}
	}

	if s.taskStore != nil {
		s.executor.taskStore = s.taskStore
	}

	go func() {
		for {
			select {
			case plan := <-s.queuedChan:
				err := s.executor.Queue(plan)
				if err != nil {

				}
			case <-s.stopChan:
				s.stopScheduler()
			case <-s.shutdownCtx.Done():
				s.stopScheduler()
				return
			}
		}
	}()

	return s, nil
}

var _ gotasker.Scheduler = (*scheduler)(nil)
