package logger

import (
  "log/slog"
  "os"
)

func New(env string) {
  var handler slog.Handler
  if env == "local" {
    handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
      Level: slog.LevelDebug,
    })
  } else if env == "dev" {
    handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
      Level: slog.LevelDebug,     
    })
  } else if env == "prod" {
    handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
      Level: slog.LevelInfo,
    })
  }

  slog.SetDefault(slog.New(handler))
}
