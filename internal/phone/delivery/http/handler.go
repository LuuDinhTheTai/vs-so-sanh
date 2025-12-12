package http

import (
	"vs-so-sanh/internal/model"
	"vs-so-sanh/internal/phone"
	"vs-so-sanh/web/page"

	"github.com/gin-gonic/gin"
	"maragu.dev/gomponents"
)

type PhoneHandler struct {
	useCase *phone.UseCase
}

func NewPhoneHandler(useCase phone.UseCase) phone.Handler {
	return &PhoneHandler{
		useCase: &useCase,
	}
}

func (p *PhoneHandler) HomePage(ctx *gin.Context) (gomponents.Node, error) {
	items := []model.Phone{
		{
			Brand:    "Samsung",
			Model:    "",
			Network:  model.Network{},
			Launch:   model.Launch{},
			Body:     model.Body{},
			Display:  model.Display{},
			Platform: model.Platform{},
			Memory:   model.Memory{},
			Camera:   model.Camera{},
			Sound:    model.Sound{},
			Comms:    model.Comms{},
			Features: model.Features{},
			Battery:  model.Battery{},
			Misc:     model.Misc{},
		},
		{
			Brand:    "Apple",
			Model:    "",
			Network:  model.Network{},
			Launch:   model.Launch{},
			Body:     model.Body{},
			Display:  model.Display{},
			Platform: model.Platform{},
			Memory:   model.Memory{},
			Camera:   model.Camera{},
			Sound:    model.Sound{},
			Comms:    model.Comms{},
			Features: model.Features{},
			Battery:  model.Battery{},
			Misc:     model.Misc{},
		},
	}

	return page.ComparePage(items, []string{"Brand"}), nil
}
