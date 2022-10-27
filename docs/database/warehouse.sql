CREATE TABLE warehouse (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    cells jsonb NOT NULL
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE public.warehouse IS 'Склад';

COMMENT ON COLUMN warehouse.id IS 'Первичный ключ';
COMMENT ON COLUMN warehouse.name IS 'Название склада в системе';
COMMENT ON COLUMN warehouse.cell IS 'Ячейки';
COMMENT ON COLUMN warehouse.created_at IS 'Дата создания';
COMMENT ON COLUMN warehouse.updated_at IS 'Дата обновления';
COMMENT ON COLUMN warehouse.deleted_at IS 'Дата удаления';

CREATE TABLE cell (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    row INTEGER NOT NULL,
    column INTEGER NOT NULL,
    num INTEGER NOT NULL,
    content jsonb NOT NULL
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE public.cell IS 'Ячейки';

COMMENT ON COLUMN cell.id IS 'Первичный ключ';
COMMENT ON COLUMN cell.name IS 'Строковое представление адреса ячейки';
COMMENT ON COLUMN cell.row IS 'Номер ряда';
COMMENT ON COLUMN cell.column IS 'Номер секции';
COMMENT ON COLUMN cell.num IS 'Номер ячейки';
COMMENT ON COLUMN cell.content IS 'Содержимое';
COMMENT ON COLUMN cell.created_at IS 'Дата создания';
COMMENT ON COLUMN cell.updated_at IS 'Дата обновления';
COMMENT ON COLUMN cell.deleted_at IS 'Дата удаления';