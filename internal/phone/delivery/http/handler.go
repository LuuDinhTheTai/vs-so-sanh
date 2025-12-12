package http

import (
	"vs-so-sanh/internal/brand"
	_ "vs-so-sanh/internal/model"
	"vs-so-sanh/internal/phone"
	"vs-so-sanh/web/page"

	"github.com/gin-gonic/gin"
	"maragu.dev/gomponents"
)

type PhoneHandler struct {
	phoneUseCase phone.UseCase
	brandUseCase brand.UseCase
}

func NewPhoneHandler(useCase phone.UseCase, brandUseCase brand.UseCase) phone.Handler {
	return &PhoneHandler{
		phoneUseCase: useCase,
		brandUseCase: brandUseCase,
	}
}

func (p *PhoneHandler) HomePage(ctx *gin.Context) (gomponents.Node, error) {
	return page.HomePage(p.brandUseCase), nil
}
