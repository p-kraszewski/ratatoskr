package application

import (
	"context"
	"sync"

	"github.com/p-kraszewski/ratatoskr/config"
)

type ErrorMap map[Task]error

type App struct {
	barrier sync.WaitGroup
	quit    context.CancelFunc
	errors  ErrorMap

	Ctx    context.Context
	Config *config.Config
}

type Task interface {
	Run(app *App) error
	String() string
}

func Run(conf *config.Config, tasks ...Task) *App {
	app := &App{
		errors: map[Task]error{},
		Config: conf,
	}

	app.Ctx, app.quit = context.WithCancel(context.Background())

	for _, t := range tasks {
		app.barrier.Add(1)
		go func(task Task) {
			defer app.barrier.Done()
			err := task.Run(app)
			app.errors[task] = err
			if err != nil {
				app.Quit()
			}
		}(t)
	}

	return app

}

func (app *App) Quit() {
	app.quit()
}

func (app *App) Wait() ErrorMap {
	app.barrier.Wait()
	return app.errors
}

func (em ErrorMap) String() string {
	ans := ""
	for task, err := range em {
		if err == nil {
			ans += task.String() + ": no error\n"
		} else {
			ans += task.String() + ": error " + err.Error() + "\n"
		}
	}
	return ans
}
