CREATE TABLE IF NOT EXISTS client (
	"id" CHAR(36) PRIMARY KEY,
	"firstname" VARCHAR(255) NOT NULL,
	"lastname" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
	"type" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP DEFAULT now(),
	"updated_at" TIMESTAMP,
	"deleted_at" TIMESTAMP
 );

	