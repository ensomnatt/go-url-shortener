package handler

import (
	"encoding/json"
	"net/http"
  "log/slog"
)

func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
  var req Request
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    http.Error(w, "invalid request", http.StatusBadRequest)
    slog.Info("неверный json запрос", "error", err)
    return
  }

  slog.Info("запрос успешно принят", "url", req.URL, "alias", req.Alias)
 
  err = h.db.Save(req.URL, req.Alias)
  if err != nil {
    http.Error(w, "ошибка при сохранении ссылки", http.StatusInternalServerError)
    slog.Error("ошибка при сохранении ссылки", "error", err)
    return
  }

  slog.Info("запрос в базу данных успешно отправлен")

  w.Header().Set("Content-Type", "application/json")
  response := Response{
    Status: "success",
    Message: "url successfully saved",
  }
  json, err := json.Marshal(response)
  if err != nil {
    http.Error(w, "ошибка при кодировании json", http.StatusInternalServerError)
    slog.Error("ошибка при кодировании json", "error", err)
    return
  } 
  slog.Debug("ответ закодирован")
  w.Write(json)

  slog.Info("ответ пользователю отправлен", "status", response.Status, "message", response.Message)
}
