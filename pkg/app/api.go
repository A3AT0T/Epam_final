package app

import (
	"context"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"Epam_final/pkg/db/repository"
	"Epam_final/pkg/middleware"
)

type Api struct {
	router          *gin.Engine
	accountRepo     *repository.AccountRepo
	userRepo        *repository.UserRepo
	userRequestRepo *repository.UserRequestRepo
	paymentRepo     *repository.PaymentRepo
	logRepo         *repository.LogRepo
	cardRepo        *repository.CardRepo
}

func New(conn *gorm.DB) (*Api, error) {
	accountRepo := repository.NewAccountRepo(conn)
	userRepo := repository.NewUserRepo(conn)
	userRequestRepo := repository.NewUserRequestRepo(conn)
	paymentRepo := repository.NewPaymentRepo(conn)
	logRepo := repository.NewLogRepo(conn)
	cardRepo := repository.NewCardRepo(conn)

	api := &Api{
		accountRepo:     accountRepo,
		userRepo:        userRepo,
		userRequestRepo: userRequestRepo,
		paymentRepo:     paymentRepo,
		logRepo:         logRepo,
		cardRepo:        cardRepo,
	}
	srv := gin.Default()
	srv.Use(middleware.DefaultStructuredLogger())
	gin.DefaultWriter = os.Stdout

	srv.POST("/users", api.CreateUserHandler)
	srv.GET("/users/:id", api.GetUserHandler)
	srv.DELETE("/users/:id", api.DeleteUserHandler)

	srv.PUT("/admin/users/:id/block", api.BlockUserHandler)
	srv.PUT("/admin/users/:id/unblock", api.UnblockUserHandler)
	srv.PUT("/admin/request/:id/approve", api.UserRequestApproveHandler)

	srv.POST("/users/:id/accounts", api.CreateAccountHandler)
	srv.GET("/users/:id/accounts", api.ListAccountHandler)

	srv.POST("/accounts/:id/card", api.CreateCardHandler)
	srv.GET("/card/:id", api.GetCardHandler)

	srv.POST("/payments", api.CreatePaymentHandler)
	srv.GET("/accounts/:id/payments", api.ListPaymentsHandler)

	srv.POST("/userRequest", api.CreateUserRequestHandler)

	api.router = srv

	return api, nil
}

func (a *Api) Run(ctx context.Context, port string) error {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: a.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	return srv.Shutdown(ctx)
}
