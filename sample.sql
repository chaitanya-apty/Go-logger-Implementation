CREATE TABLE accounts (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP without time zone default (now() at time zone 'utc'),
    last_login TIMESTAMP 
);

INSERT INTO accounts VALUES (1,'Chaitanya','c@gmail.com'),
							(2,'Kiran','k@gmail.com'),
							(3,'Srini','s@gmail.com')