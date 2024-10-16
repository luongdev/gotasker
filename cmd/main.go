package main

import (
	"github.com/luongdev/gotasker/internal"
	"github.com/luongdev/gotasker/pkg"
)

func main() {
	taskStore := pkg.NewInMemoryTaskStore(
		pkg.MockTask("task1"),
	)

	planStore := pkg.NewInMemoryPlanStore(
		pkg.MockPlan("task1"),
	)
	s, err := pkg.NewScheduler(pkg.WithTaskStore(taskStore), internal.WithPlanStore(planStore))

	if err != nil {
		panic(err)
	}

	s.Start()

	for {

	}
}
