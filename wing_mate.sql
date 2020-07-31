CREATE TABLE user_profile (
	username VARCHAR ( 50 ) PRIMARY KEY,
	password VARCHAR ( 255 ) NOT NULL,
	first_name VARCHAR (50) NOT NULL,
	last_name VARCHAR (50) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	app_use_status TEXT CHECK (app_use_status IN ('dater', 'wing_mate', 'relative','other')),
	created_on timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE wing_profile (
	username VARCHAR ( 50 ) PRIMARY KEY,
	dater_username VARCHAR ( 50 ),
	relationship_to_dater TEXT CHECK (relationship_to_dater IN ('friend', 'relative','other')),
	date_of_birth DATE NOT NULL,
	intro_wing_line VARCHAR (255),
	current_city VARCHAR (50),
	last_modified_on timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE dater_profile (
	username VARCHAR ( 50 ) PRIMARY KEY,
	wing_username VARCHAR ( 50 ),
	search_permission BOOLEAN NOT NULL DEFAULT false,
	date_of_birth DATE NOT NULL,
	current_city VARCHAR (50),
	job_role VARCHAR (50),
	employment_status VARCHAR (50),
	study_college VARCHAR (255),
	short_intro VARCHAR (255),
	last_modified_on timestamptz NOT NULL DEFAULT now()
);