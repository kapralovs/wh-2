package models

import "time"

type User struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Фамилия
	Surname string `json:"surname,omitempty"`
	// Имя
	Name string `json:"name,omitempty"`
	// Ник пользователя
	Username string `json:"username,omitempty"`
	// Пароль
	Password string `json:"password,omitempty"`
	// Идентификатор роли
	RoleID int64 `json:"role_id,omitempty"`
	// Дата создания
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Дата обновления
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Дата удаления
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type Roles struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Название
	Name string `json:"name,omitempty"`
	// Описание
	Description string `json:"description,omitempty"`
}
