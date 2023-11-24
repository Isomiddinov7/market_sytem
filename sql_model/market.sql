CREATE DATABASE market;

CREATE TABLE "category" (
    "id" UUID NOT NULL PRIMARY KEY,
    "title" VARCHAR(46) NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "parent_id" UUID REFERENCES "category" ("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "product" (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "title" VARCHAR(46) NOT NULL,
    "description" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "category_id" UUID NOT NULL REFERENCES "category"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "clients"(
    "id" UUID NOT NULL PRIMARY KEY,
    "first_name" VARCHAR(64) NOT NULL,
    "last_name" VARCHAR(128) NOT NULL,
    "phone" VARCHAR(13) NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "data_of_birth" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "branches"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "phone" VARCHAR(13) NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "address" VARCHAR(128) NOT NULL,
    "work_start_hour" TIME NOT NULL,
    "work_end_hour" TIME NOT NULL,
    "delivery_price" NUMERIC DEFAULT 10000,
    "active" BOOLEAN NOT NULL DEFAULT false,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "orders"(
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "client_id" UUID REFERENCES "clients"("id"),
    "branch_id" UUID REFERENCES "branches"("id"),
    "address" VARCHAR(128) NOT NULL,
    "delivery_price" NUMERIC,
    "total_count" INT,
    "total_price" NUMERIC,
    "status" VARCHAR NOT NULL DEFAULT 'new',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "order_products"(
    "product_id" VARCHAR REFERENCES "product"("id"),
    "discount_type" VARCHAR,
    "discount_amount" INT,
    "quantity" INT,
    "price" NUMERIC,
    "sum" NUMERIC

);
