package main

import (
	"log/slog"
	"net/http"
	"os"
	"urlshortener/internal/logger"
	"urlshortener/internal/storage/db"
  "urlshortener/internal/handler"

	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")
  if err != nil {
    panic(err)
  }

  env := os.Getenv("ENV")
  connStr := os.Getenv("DB_CONN_STR")
  port := os.Getenv("PORT")
  
  logger.New(env)
  slog.With("env", env)
  slog.Info("логгер инициализирован")
  slog.Debug("дебаг логи подключены")
  slog.Error("логи ошибок подключены")
  slog.Warn("логи предупреждений подключены")

  db, err := db.New(connStr) 
  if err != nil {
    slog.Error("не удалось создать базу данных", "error", err)
  }
  defer db.Close()
  db.Init()

  r := http.NewServeMux()
  handler := handler.New(db)
  handler.CreateRoutes(r)

  http.ListenAndServe(":" + port, r)
}
