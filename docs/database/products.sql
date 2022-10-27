CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    count INTEGER NOT NULL
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE public.products IS 'Продукты';

COMMENT ON COLUMN products.id IS 'Первичный ключ';
COMMENT ON COLUMN products.name IS 'Название';
COMMENT ON COLUMN products.count IS 'Количество';
COMMENT ON COLUMN products.created_at IS 'Дата создания';
COMMENT ON COLUMN products.updated_at IS 'Дата обновления';
COMMENT ON COLUMN products.deleted_at IS 'Дата удаления';