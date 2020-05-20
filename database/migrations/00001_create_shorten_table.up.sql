CREATE TABLE IF NOT EXISTS shorten(
	"id" SERIAL NOT NULL,
    "url" varchar(100) NOT NULL,
    "shortcode" varchar(100) NOT NULL,
    "redirect_count" int2 DEFAULT 0 NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "last_seen_date" timestamp
);