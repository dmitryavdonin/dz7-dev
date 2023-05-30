CREATE TABLE "order" (
    "id" uuid primary key,
    "product_id" int,
    "product_count" int,
    "product_price" float,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
)