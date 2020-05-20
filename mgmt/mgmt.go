package mgmt

import (
	"time"

	"github.com/p-kraszewski/ratatoskr/application"
)

type Mgmt struct{}

func (mgmt *Mgmt) String() string {
	return "MGMT"
}

func (mgmt *Mgmt) Run(app *application.App) error {
	time.Sleep(time.Second * 5)
	return nil
}
