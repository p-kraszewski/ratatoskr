package mgmt

import (
	"os"
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

type MgmtB struct{}

func (mgmt *MgmtB) String() string {
	return "MGMT broken"
}

func (mgmt *MgmtB) Run(app *application.App) error {
	time.Sleep(time.Second * 3)
	return os.ErrInvalid
}
