CREATE TABLE tasks (
    id BIGSERIAL PRIMARY KEY,
    type_of VARCHAR NOT NULL,
    assigned_to BIGINT,
    event_id BIGINT NOT NULL,
    event jsonb NOT NULL
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE public.tasks IS 'Задачи';

COMMENT ON COLUMN tasks.id IS 'Первичный ключ';
COMMENT ON COLUMN tasks.type_of IS 'Идентификатор исполнителя';
COMMENT ON COLUMN tasks.assigned_to IS 'Идентификатор исполнителя';
COMMENT ON COLUMN tasks.event IS 'Содержимое задачи';
COMMENT ON COLUMN tasks.created_at IS 'Дата создания';
COMMENT ON COLUMN tasks.updated_at IS 'Дата обновления';
COMMENT ON COLUMN tasks.deleted_at IS 'Дата удаления';

CREATE TABLE order_tasks (
    id BIGSERIAL PRIMARY KEY,
    delivery_id BIGINT NOT NULL,
    client_id BIGINT NOT NULL,
    shipping_units jsonb NOT NULL
);

COMMENT ON TABLE public.order_task IS 'Заказы';

COMMENT ON COLUMN order_task.id IS 'Первичный ключ';
COMMENT ON COLUMN order_task.delivery_id IS 'Идентификатор поставки';
COMMENT ON COLUMN order_task.client_id IS 'Идентификатор клиента';
COMMENT ON COLUMN order_task.shipping_units IS 'Единицы отгрузки';

CREATE TABLE refill_tasks (
    id BIGSERIAL PRIMARY KEY,
    releasing_cell VARCHAR NOT NULL,
    product_id BIGINT NOT NULL,
    product_count INTEGER NOT NULL,
    host_cell VARCHAR NOT NULL
);

COMMENT ON TABLE public.refill IS 'Пополнение пик-зоны';

COMMENT ON COLUMN refill.id IS 'Первичный ключ';
COMMENT ON COLUMN refill.releasing_cell IS 'Отпускающее складское место';
COMMENT ON COLUMN refill.product_id IS 'Идентификатор клиента';
COMMENT ON COLUMN refill.product_count IS 'Количество продукта';
COMMENT ON COLUMN refill.host_cell IS 'Принимающее складское место';

CREATE TABLE transport_tasks (
    id BIGSERIAL PRIMARY KEY,
    releasing_warehouse VARCHAR NOT NULL,
    delivery_id BIGINT NOT NULL,
    shipping_unit_id BIGINT NOT NULL
);

COMMENT ON TABLE public.transport_task IS 'Отгрузка';

COMMENT ON COLUMN transport_task.id IS 'Первичный ключ';
COMMENT ON COLUMN transport_task.releasing_warehouse IS 'Отпускающий склад';
COMMENT ON COLUMN transport_task.delivery_id IS 'Идентификатор поставки';
COMMENT ON COLUMN transport_task.shipping_unit_id IS 'Идентификатор единицы отгрузки';

CREATE TABLE replacing_tasks (
    id BIGSERIAL PRIMARY KEY,
    releasing_cell VARCHAR NOT NULL,
    product_id BIGINT NOT NULL,
    product_count INTEGER NOT NULL,
    host_cell VARCHAR NOT NULL
);

COMMENT ON TABLE public.replacing_tasks IS 'Движение товара';

COMMENT ON COLUMN replacing_tasks.id IS 'Первичный ключ';
COMMENT ON COLUMN replacing_tasks.releasing_cell IS 'Отпускающее складское место';
COMMENT ON COLUMN replacing_tasks.product_id IS 'Товар';
COMMENT ON COLUMN replacing_tasks.product_count IS 'Количество единиц товара';
COMMENT ON COLUMN replacing_tasks.host_cell IS 'Принимающее складское место';

CREATE TABLE inventory_tasks (
    id BIGSERIAL PRIMARY KEY,
    cell VARCHAR NOT NULL,
    product_id BIGINT NOT NULL,
    declared_quantity INTEGER NOT NULL,
    actual_quantity INTEGER NOT NULL,
);

COMMENT ON TABLE public.inventory_tasks IS 'Инвентаризация';

COMMENT ON COLUMN inventory_tasks.id IS 'Первичный ключ';
COMMENT ON COLUMN inventory_tasks.cell IS 'Ячейка';
COMMENT ON COLUMN inventory_tasks.product_id IS 'Продукт';
COMMENT ON COLUMN inventory_tasks.declared_quantity IS 'Заявленное количество';
COMMENT ON COLUMN inventory_tasks.actual_quantity IS 'Фактическое количество';

CREATE TABLE product_placement_tasks (
    id BIGSERIAL PRIMARY KEY,
    releasing_cell VARCHAR NOT NULL,
    product_id BIGINT NOT NULL,
    product_count INTEGER NOT NULL,
    host_cell VARCHAR NOT NULL
);

COMMENT ON TABLE public.product_placement_tasks IS 'Размещение товара на складе';

COMMENT ON COLUMN product_placement_tasks.id IS 'Первичный ключ';
COMMENT ON COLUMN product_placement_tasks.releasing_cell IS 'Отпускающее складское место';
COMMENT ON COLUMN product_placement_tasks.product_id IS 'Идентификатор клиента';
COMMENT ON COLUMN product_placement_tasks.product_count IS 'Единицы отгрузки';
COMMENT ON COLUMN product_placement_tasks.host_cell IS 'Принимающее складское место';

CREATE TABLE special_situation (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL
);

COMMENT ON TABLE public.special_situation IS 'Особые ситуации';

COMMENT ON COLUMN special_situation.id IS 'Первичный ключ';
COMMENT ON COLUMN special_situation.name IS 'Название';
COMMENT ON COLUMN special_situation.description IS 'Описание';