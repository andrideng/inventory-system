BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `products` (
	`sku`	TEXT,
	`name`	TEXT,
	`size`	TEXT,
	`color`	TEXT,
	`amount`	INTEGER,
	`created_at`	TEXT,
	`updated_at`	TEXT,
	`total_incoming_amount`	INTEGER,
	`total_incoming_price`	NUMERIC,
	`average_purchase_price`	INTEGER,
	PRIMARY KEY(`sku`)
);

CREATE TABLE IF NOT EXISTS `outgoing_goods` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`sku`	TEXT,
	`amount`	INTEGER,
	`price`	NUMERIC,
	`total`	NUMERIC,
	`note`	TEXT,
	`created_at`	TEXT,
	`updated_at`	TEXT,
	`order_id`	TEXT
);

CREATE TABLE IF NOT EXISTS `incoming_goods` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`receipt_number`	TEXT,
	`sku`	TEXT,
	`amount_orders`	INTEGER,
	`amount_received`	INTEGER,
	`purchase_price`	NUMERIC,
	`total`	NUMERIC,
	`note`	TEXT,
	`status`	TEXT,
	`created_at`	TEXT,
	`updated_at`	TEXT
);
COMMIT;
