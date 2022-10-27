package models

type User struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Фамилия
	Surname string `json:"surname,omitempty"`
	// Имя
	Name string `json:"name,omitempty"`
	// Идентификатор роли
	RoleID int64 `json:"role_id,omitempty"`
	// Идентификатор должности
	JobID int64 `json:"job_id,omitempty"`
	// Ник пользователя
	Username string `json:"username,omitempty"`
	// Пароль
	Password string `json:"password,omitempty"`
}

type Roles struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Название
	Name string `json:"name,omitempty"`
	// Описание
	Description string `json:"description,omitempty"`
}
