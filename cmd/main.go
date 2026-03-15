package main

import (
	"auto-pharmacy/internal/config"
	"auto-pharmacy/internal/controllers"
	"auto-pharmacy/internal/repository/pgsql"
	"auto-pharmacy/internal/routes"
	"auto-pharmacy/internal/services"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type CustomValidator struct {
    v *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
    return cv.v.Struct(i)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env:", err)
	}
}

func main() {
	flag.Parse()

	// Загружаем переменные окружения
	loadEnv()

	pg, err := config.InitPgsql()
	if err != nil {
		panic(err)
	}

	// defer database.Disconnect()
	// models.MedMigrate()
	// models.UserMigrate()

	userRepo := pgsql.NewUserRepository(pg.GetDBConn())

	userService := services.NewUserService(userRepo)

	uc := controllers.NewUserController(userService)
	lc := controllers.NewLoginController(userService)

  	e := echo.New()
	e.Validator = &CustomValidator{v: validator.New()}
  	e.Use(middleware.RequestLogger())
	e.Use(middleware.CORS("http://localhost:3000", "http://localhost"))
	routes.RegisterWebRoutes(e.Group(""), uc)
	routes.RegisterAuthRoutes(e.Group(""), lc)

	srv := http.Server{
		Addr:         ":8080",
		Handler:      e,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	wg := sync.WaitGroup{}
	wg.Go(func() {
		log.Print("HTTP-сервер запускается")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			e.Logger.Error("Ошибка запуска HTTP-сервера", "error", err)
		}
	})

	// Ждем сигналов ОС
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Получаем сигнал завершения работы
	sig := <-stopChan
	log.Printf("\nПолучен сигнал завершения (%v)\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Завершаем работу сервера
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка выключения сервера: %v", err)
	}

	wg.Wait()
	log.Println("Завершили работу...")
}
