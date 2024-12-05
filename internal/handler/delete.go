package handler

import (
	"log/slog"
	"net/http"
  "encoding/json"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
  alias := r.PathValue("alias")
  slog.Info("запрос успешно принят", "alias", alias)

  err := h.db.Delete(alias)
  if err != nil {
    http.Error(w, "ошибка при удалении ссылки", http.StatusInternalServerError)
    slog.Error("ошибка при запросе в базу данных на удаление ссылки")
    return
  }
  slog.Info("ссылка была удалена", "alias", alias)

  w.Header().Set("Content-Type", "application/json")
  response := Response{
    Status: "success",
    Message: "url successfully deleted",
  }
  json, err := json.Marshal(response)
  if err != nil {
    http.Error(w, "ошибка при кодировании json", http.StatusInternalServerError)
    slog.Error("ошибка при кодировании json", "error", err)
    return
  }
  w.Write(json)
}
