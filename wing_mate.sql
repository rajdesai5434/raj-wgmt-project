CREATE TABLE user_profile (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	first_name VARCHAR (50) NOT NULL,
	last_name VARCHAR (50) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	short_intro VARCHAR ( 255 ),
	app_use_status TEXT CHECK (app_use_status IN ('dater', 'friend', 'relative','other')),
	created_on TIMESTAMP NOT NULL,
	last_login TIMESTAMP
);