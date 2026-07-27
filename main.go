package main

import (
	handler2 "rental_car/internal/auth/login/handler"
	repository2 "rental_car/internal/auth/login/repository"
	useCase2 "rental_car/internal/auth/login/useCase"
	"rental_car/internal/auth/register/handler"
	"rental_car/internal/auth/register/repository"
	"rental_car/internal/auth/register/useCase"
	handler4 "rental_car/internal/car/handler"
	repository4 "rental_car/internal/car/repository"
	useCase4 "rental_car/internal/car/useCase"
	"rental_car/internal/platform"
	"rental_car/internal/platform/middleware"
	handler3 "rental_car/internal/user/handler"
	repository3 "rental_car/internal/user/repository"
	useCase3 "rental_car/internal/user/useCase"

	"github.com/gin-gonic/gin"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	db := platform.ConnectDB()
	platform.ConnectGoogleStorage()

	repo := repository.NewAuthRepo(db)
	UseCase := useCase.NewAuthCase(repo)
	Handler := handler.NewRegisterHandler(UseCase)

	repoLogin := repository2.NewLoginRepository(db)
	UseCaseLogin := useCase2.NewUseCaseLogin(repoLogin)
	HandlerLogin := handler2.NewLoginHandler(UseCaseLogin)

	repoUser := repository3.NewUserRepository(db)
	useCaseUser := useCase3.NewUserUseCase(repoUser)
	handlerUser := handler3.NewUserHandler(useCaseUser)

	repoCar := repository4.NewRepositoryCar(db)
	useCaseCar := useCase4.NewCarUseCase(repoCar)
	handlerCar := handler4.NewHandlerCar(useCaseCar)

	r := gin.Default()

	r.POST("/login", HandlerLogin.LoginUser)
	r.POST("/register", Handler.CreateUser)
	r.GET("/car", handlerCar.GetAllCars)

	r.Use(middleware.AuthMiddleware())

	r.GET("/user", handlerUser.GetUser)
	r.PATCH("/user", handlerUser.UpdateUser)
	r.DELETE("/user", handlerUser.DeleteUser)

	r.POST("/car", handlerCar.UploadRentalCar)
	r.GET("/car/:id", handlerCar.GetCarById)

	err := r.Run(":8080")
	if err != nil {
		return
	}

	err = db.Close()
	if err != nil {
		return
	}

}
