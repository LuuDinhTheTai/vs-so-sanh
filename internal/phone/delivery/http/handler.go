package http

import (
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
	items := []string{"Super", "Duper", "Nice"}
	return page.HomePage(items), nil
}
