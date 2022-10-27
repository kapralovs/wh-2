CREATE TABLE deliveries (
    id BIGSERIAL PRIMARY KEY,
    client_id BIGINT NOT NULL,
    address VARCHAR NOT NULL,
    content jsonb NOT NULL
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE public.deliveries IS 'Поставки';

COMMENT ON COLUMN deliveries.id IS 'Первичный ключ';
COMMENT ON COLUMN deliveries.client_id IS 'Идентификатор заказчика';
COMMENT ON COLUMN deliveries.address IS 'Адрес';
COMMENT ON COLUMN deliveries.content IS 'Содержимое';
COMMENT ON COLUMN deliveries.created_at IS 'Дата создания';
COMMENT ON COLUMN deliveries.updated_at IS 'Дата обновления';
COMMENT ON COLUMN deliveries.deleted_at IS 'Дата удаления';

CREATE TABLE shipping_units (
    id BIGSERIAL PRIMARY KEY,
    task_id BIGINT NOT NULL,
    employee_id BIGINT NOT NULL,
    content jsonb NOT NULL
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE public.shipping_units IS 'Едининцы отгрузки';

COMMENT ON COLUMN shipping_units.id IS 'Первичный ключ';
COMMENT ON COLUMN shipping_units.client_id IS 'Идентификатор задачи';
COMMENT ON COLUMN shipping_units.employee_id IS 'Идентификатор исполнителя';
COMMENT ON COLUMN shipping_units.content IS 'Содержимое';
COMMENT ON COLUMN shipping_units.created_at IS 'Дата создания';
COMMENT ON COLUMN shipping_units.updated_at IS 'Дата обновления';
COMMENT ON COLUMN shipping_units.deleted_at IS 'Дата удаления';