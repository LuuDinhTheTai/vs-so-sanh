package http

import (
	_ "vs-so-sanh/internal/model"
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
	return page.HomePage(), nil
}
