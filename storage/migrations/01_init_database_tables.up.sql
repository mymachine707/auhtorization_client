CREATE TABLE IF NOT EXISTS product (
   "id" CHAR(36) PRIMARY KEY,
   "category_id"  CHAR(36),
   "product_name" VARCHAR(255) UNIQUE NOT NULL,
   "description" VARCHAR(255) NOT NULL,
   "price" VARCHAR(255) NOT NULL,
   "created_at" TIMESTAMP DEFAULT now(),
   "updated_at" TIMESTAMP,
   "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS category (
	"id" CHAR(36) PRIMARY KEY,
	"category_name" VARCHAR(255) NOT NULL,
	"description" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP DEFAULT now(),
	"updated_at" TIMESTAMP,
	"deleted_at" TIMESTAMP
 );

ALTER TABLE product ADD CONSTRAINT fk_product_category FOREIGN KEY ("category_id") REFERENCES category ("id");

