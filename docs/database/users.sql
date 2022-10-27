CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    surname VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    role_id BIGINT NOT NULL
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)

COMMENT ON TABLE public.users IS 'Пользователи';

COMMENT ON COLUMN users.id IS 'Первичный ключ';
COMMENT ON COLUMN users.name IS 'Имя';
COMMENT ON COLUMN users.surname IS 'Фамилия';
COMMENT ON COLUMN users.username IS 'Ник пользователя';
COMMENT ON COLUMN users.password IS 'Пароль';
COMMENT ON COLUMN users.role_id IS 'Роль';
COMMENT ON COLUMN users.created_at IS 'Дата создания';
COMMENT ON COLUMN users.updated_at IS 'Дата обновления';
COMMENT ON COLUMN users.deleted_at IS 'Дата удаления';

CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name name NOT NULL,
    description VARCHAR NOT NULL
)

COMMENT ON TABLE public.roles IS 'Склад';

COMMENT ON COLUMN roles.id IS 'Первичный ключ';
COMMENT ON COLUMN roles.name IS 'Название';
COMMENT ON COLUMN roles.content IS 'Содержимое';

INSERT INTO roles (name,description) VALUES ('admin','Администратор всего и вся')
INSERT INTO roles (name,description) VALUES ('order_picker','Комплектовщик')
INSERT INTO roles (name,description) VALUES ('dispatcher','Диспетчер')
INSERT INTO roles (name,description) VALUES ('client','Клиент')

