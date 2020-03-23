CREATE TABLE IF NOT EXISTS shorten(
	id serial PRIMARY KEY,
	url VARCHAR (100) NOT NULL,
	shortcode VARCHAR (100) NOT NULL,
	created_at timestamp NOT NULL
 ); 