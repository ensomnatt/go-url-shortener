package db

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type Storage struct {
  db *sql.DB
}

func New(connStr string) (*Storage, error) {
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    slog.Error("ошибка подключения к базе данных", "error", err, "db connection string", connStr)
    return nil, err
  } else {    
    slog.Debug("база данных открыта")
  }

  err = db.Ping()
  if err != nil {
    slog.Error("ошибка проверки соединения с базой данных", "error", err)
    return nil, err
  } else {
    slog.Debug("проверка соединения с базой данных прошла успешно")
  }

  return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
  return s.db.Close()
}

func (s *Storage) Init() error {
  query := `CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    alias TEXT NOT NULL
  )`
  _, err := s.db.Exec(query)
  if err != nil {
    slog.Error("ошибка при создании таблицы urls", "error", err)
    return err
  }

  return err
}

func (s *Storage) Save(url string, alias string) error {
  var exist int 
  query := `SELECT 1 FROM urls WHERE alias = $1`
  err := s.db.QueryRow(query, alias).Scan(&exist) 
  if err != nil {
    if err == sql.ErrNoRows {
      slog.Info("ссылка не найдена, сохраняем", "alias", alias, "url", url)
    } else {
      slog.Error("ошибка при проверке на существование ссылки", "error", err, "alias", alias)
    }
  } else {
    slog.Info("ссылка уже существует", "alias", alias, "url", url)
    return nil
  }

  query = `INSERT INTO urls (original_url, alias) VALUES ($1, $2)`
  _, err = s.db.Exec(query, url, alias)
  if err != nil {
    slog.Error("ошибка при сохранении ссылки", "error", err)
    return err
  } else {
    slog.Info("ссылка сохранена", "alias", alias, "url", url)
  }

  return err
}

func (s *Storage) Get(alias string) (string, error) {
  var url string 
  query := `SELECT original_url FROM urls WHERE alias = $1`
  err := s.db.QueryRow(query, alias).Scan(&url)
  if err != nil {
    slog.Error("ошибка при получении ссылки", "error", err, "alias", alias)
  } else {
    slog.Info("ссылка успешно получена", "alias", alias, "url", url)
  }

  return url, err
}

func (s *Storage) Delete(alias string) error {
  query := `DELETE FROM urls WHERE alias = $1`
  _, err := s.db.Exec(query, alias)
  if err != nil {
    slog.Error("ошибка при удалении ссылки", "error", err, "alias", alias)
  } else {
    slog.Info("ссылка удалена", "alias", alias)
  }

  return err
}
