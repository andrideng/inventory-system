-- PRODUCTS
DROP TABLE IF EXISTS "products";

CREATE TABLE "products"
(
    `sku` TEXT,
    `name` TEXT,
    `amount` INTEGER,
    `created_at` TEXT,
    `updated_at` TEXT,
    PRIMARY KEY(`sku`)
);

INSERT INTO "products" (sku, name, amount, created_at, updated_at) VALUES('SSI-D00791015-LL-BWH', 'Zalekia Plain Casual Blouse (L,Broken White)', 154, datetime(), datetime());
INSERT INTO "products" (sku, name, amount, created_at, updated_at) VALUES('SSI-D00791077-MM-BWH', 'Zalekia Plain Casual Blouse (M,Broken White)', 138, datetime(), datetime());
INSERT INTO "products" (sku, name, amount, created_at, updated_at) VALUES('SSI-D00791091-XL-BWH', 'Zalekia Plain Casual Blouse (XL,Broken White)', 137, datetime(), datetime());

-- INCOMING GOODS
DROP TABLE IF EXISTS "incmoing_goods";

CREATE TABLE "incoming_goods"
(
    `receipt_number` TEXT,
    `sku` TEXT,
    `amount_orders` INTEGER,
    `amount_received` INTEGER,
    `purchase_price` NUMERIC,
    `total` NUMERIC,
    `note` TEXT,
    `status` TEXT,
    `created_at` TEXT,
    `updated_at` TEXT,
    PRIMARY KEY(`receipt_number`)
);

-- OUTGOING GOODS
DROP TABLE IF EXISTS "outgoing_goods";

CREATE TABLE "outgoing_goods"
(
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `sku` TEXT,
    `amount` INTEGER,
    `price` NUMERIC,
    `total` NUMERIC,
    `created_at` TEXT,
    `updated_at` TEXT
);
