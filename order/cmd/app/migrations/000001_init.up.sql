CREATE TABLE "order" (
    "id" uuid primary key,
    "msg_id" uuid,
    "product_id" int,
    "product_count" int,
    "product_price" float,
    "version" int,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
);

CREATE TABLE "message" (
    "id" uuid UNIQUE,
    "created_at" timestamp not null
);