package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/kost/application/handler/user"
	userRepo "github.com/umardev500/kost/application/repository/user"
	userUc "github.com/umardev500/kost/application/usecase/user"
	"github.com/umardev500/kost/config"
)

func (a *Api) User(r fiber.Router) {
	tr := config.NewTransaction(a.db)
	ur := userRepo.NewUserRepository(tr)
	uuc := userUc.NewUserUsecase(ur, tr)
	uh := user.NewUserhandler(uuc)

	r.Delete("/:id", uh.Delete)
	r.Get("/", uh.Find)
	r.Get("/:id", uh.FindByID)
}
