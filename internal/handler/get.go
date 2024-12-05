package handler

import (
	"log/slog"
	"net/http"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
  alias := r.PathValue("alias")
  slog.Info("запрос успешно принят", "alias", alias)

  url, err := h.db.Get(alias)
  if err != nil {
    slog.Error("ошибка при запросе на получение ссылки из базы данных", "error", err)
    return
  }

  slog.Info("url был получен", "url", url)

  http.Redirect(w, r, url, http.StatusFound)
  slog.Info("пользователь был перенаправлен на оригинальную ссылку", "url", url)
}
