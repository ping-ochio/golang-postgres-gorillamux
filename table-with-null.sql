-- -------------- USERS TABLE  ------------------
CREATE TABLE users (
	user_id SERIAL NOT NULL,
	user_name VARCHAR(50),
	name VARCHAR(50),
	surname VARCHAR(50),
	email VARCHAR(50),
	password VARCHAR(100), -- 100 because hash password
	age SMALLINT,
	active BOOLEAN,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP
);

-- -------------- PRODUCTS TABLE -----------------
DROP TABLE IF EXISTS products;

CREATE TABLE products(
	product_id serial primary key,
	product_name VARCHAR(255) NOT NULL,
	description VARCHAR(500),
	price	NUMERIC,
	quantity INTEGER,
	create_date TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP
);



-- ------------- INSERT COMMAND -------------------
insert into users
(user_name, name, surname, password, age, active) 
values('s','s','s','s',d,bool);

insert into products(product_name, description, price, quantity) 
values('s','s',f,d);

-- --------- QUERY FOR A NULL VALUE ---------------

select * from products where quantity is NULL;


-- ------------- UPDATE COMMAND -------------------
update users set user_name = 's' where user_name = 's';

-- --------- SHOW TABLE'S COLUMNS & TYPES --------- 
\d products

-- ------------ CASE INSENSITIVE ------------------
select * from users where name ~* 'an' ;


-- ----- HOW TO CONNECT TO MONGODB IN CONSOLE -----
mongosh --username admin --password password --port 27017