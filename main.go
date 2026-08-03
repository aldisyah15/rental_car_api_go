package main

import (
	"net/http"
	handler2 "rental_car/internal/auth/login/handler"
	repository2 "rental_car/internal/auth/login/repository"
	useCase2 "rental_car/internal/auth/login/useCase"
	"rental_car/internal/auth/register/handler"
	repoRegister "rental_car/internal/auth/register/repository"
	useCaseRegister "rental_car/internal/auth/register/useCase"
	handler7 "rental_car/internal/booking/handler"
	repository7 "rental_car/internal/booking/repository"
	useCase7 "rental_car/internal/booking/useCase"
	handler5 "rental_car/internal/brand/handler"
	repository5 "rental_car/internal/brand/repository"
	useCase5 "rental_car/internal/brand/useCase"
	handler4 "rental_car/internal/car/handler"
	repository4 "rental_car/internal/car/repository"
	useCase4 "rental_car/internal/car/useCase"
	handler6 "rental_car/internal/favorite/handler"
	repository6 "rental_car/internal/favorite/repository"
	useCase6 "rental_car/internal/favorite/useCase"
	handler3 "rental_car/internal/user/handler"
	repository3 "rental_car/internal/user/repository"
	useCase3 "rental_car/internal/user/useCase"
	platform2 "rental_car/platform"
	"rental_car/platform/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	db := platform2.ConnectDB()
	platform2.ConnectGoogleStorage()

	repo := repoRegister.NewAuthRepo(db)
	UseCase := useCaseRegister.NewAuthCase(repo)
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

	repoLogo := repository5.NewRepositoryLogo(db)
	useCaseLogo := useCase5.NewUseCaseLogo(repoLogo)
	handlerLogo := handler5.NewCarHandlerLogo(useCaseLogo)

	repoFavorite := repository6.NewRepositoryFavorite(db)
	useCasefavorite := useCase6.NewUseCaseFavorite(repoFavorite)
	handlerFavorite := handler6.NewHandlerFavorite(useCasefavorite)

	repoBooking := repository7.NewRepositoryBooking(db, &http.Client{Timeout: 15 * time.Second})
	useCaseBooking := useCase7.NewUseCaseBooking(repoBooking, repoCar)
	handlerBooking := handler7.NewHandlerBooking(useCaseBooking)
	r := gin.Default()

	r.POST("/login", HandlerLogin.LoginUser)
	r.POST("/register", Handler.CreateUser)
	r.GET("/car", handlerCar.GetAllCars)

	r.Use(middleware.AuthMiddleware())

	r.GET("/user", handlerUser.GetUser)
	r.PATCH("/user", handlerUser.UpdateUser)
	r.DELETE("/user", handlerUser.DeleteUser)
	r.POST("/user/favorite", handlerFavorite.Favorite)
	r.GET("/user/favorite", handlerFavorite.GetFavorite)

	r.POST("/car", handlerCar.UploadRentalCar)
	r.GET("/car/:id", handlerCar.GetCarById)

	r.POST("/logo", handlerLogo.UploadLogo)
	r.GET("/logo", handlerLogo.GetAllLogo)

	r.POST("/proses-checkout", handlerBooking.ProsesToCheckout)
	r.GET("/detail-checkout/:idOrder", handlerBooking.GetDetailCheckout)
	r.POST("/booking/:idOrder", handlerBooking.BookingNow)
	err := r.Run(":8080")
	if err != nil {
		return
	}

	err = db.Close()
}
