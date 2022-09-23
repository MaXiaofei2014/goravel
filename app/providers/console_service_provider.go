package providers

import (
	"github.com/goravel/framework/support/facades"
	"goravel/app/console"
)

type ConsoleServiceProvider struct {
}

func (receiver *ConsoleServiceProvider) Boot() {

}

func (receiver *ConsoleServiceProvider) Register() {
	kernel := console.Kernel{}
	facades.Schedule.Register(kernel.Schedule())
	facades.Artisan.Register(kernel.Commands())
}
