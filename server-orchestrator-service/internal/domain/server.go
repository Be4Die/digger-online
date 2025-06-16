package domain

type Server struct {
	ID      string `json:"id" binding:"required"`      // Уникальный идентификатор
	Address string `json:"address" binding:"required"` // Адрес сервера
	Port    int    `json:"port" binding:"required"`    // Порт сервера
}