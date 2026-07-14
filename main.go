package main

import (
	handler2 "rental_car/internal/auth/login/handler"
	repository2 "rental_car/internal/auth/login/repository"
	useCase2 "rental_car/internal/auth/login/useCase"
	"rental_car/internal/auth/register/handler"
	"rental_car/internal/auth/register/repository"
	"rental_car/internal/auth/register/useCase"
	"rental_car/internal/platform"

	"github.com/gin-gonic/gin"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	db := platform.ConnectDB()
	repo := repository.NewAuthRepo(db)
	UseCase := useCase.NewAuthCase(repo)
	Handler := handler.NewRegisterHandler(UseCase)

	repoLogin := repository2.NewLoginRepository(db)
	UseCaseLogin := useCase2.NewUseCaseLogin(repoLogin)
	HandlerLogin := handler2.NewLoginHandler(UseCaseLogin)

	r := gin.Default()
	r.POST("/login", HandlerLogin.LoginUser)
	r.POST("/register", Handler.CreateUser)

	

	err := r.Run(":8080")
	if err != nil {
		return
	}

	err = db.Close()
	if err != nil {
		return
	}
}
