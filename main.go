package main

import (
	"auto-pharmacy/controllers"
	"auto-pharmacy/database"
	"auto-pharmacy/models"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

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

	db, err := database.RegisterMysql()
	if err != nil {
		panic(err)
	}

	database.MysqlDB = db
	// defer database.Disconnect()
	models.MedMigrate()

	r := mux.NewRouter()
	r.HandleFunc("/medicines", controllers.MedicineIndex).Methods("GET")
	r.HandleFunc("/medicines", controllers.MedicineSet).Methods("POST")
	r.HandleFunc("/medicines/{medicine}", controllers.MedicineGet).Methods("GET")
	r.HandleFunc("/medicines/{medicine}", controllers.MedicineDelete).Methods("DELETE")
	r.HandleFunc("/medicines/{medicine}/restore", controllers.MedicineRestore).Methods("PUT")
	r.HandleFunc("/medicines/{medicine}/force", controllers.MedicineForceDelete).Methods("DELETE")

	r.HandleFunc("/supply", controllers.SupplyIndex).Methods("GET")
	r.HandleFunc("/supply", controllers.SupplySet).Methods("POST")
	r.HandleFunc("/supply/{supply}", controllers.SupplyGet)

	srv := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Print("HTTP-сервер запускается")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска HTTP-сервера: %v", err)
		}
	}()

	// Ждем сигналов ОС
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Получаем сигнал завершения работы
	sig := <-stopChan
	fmt.Printf("\nПолучен сигнал завершения (%v)\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Завершаем работу сервера
	srv.Shutdown(ctx)
	// if err := ; err != nil {
	// 	log.Fatalf("Ошибка выключения сервера: %v", err)
	// }

	wg.Wait()
	fmt.Println("Завершили работу...")
}
