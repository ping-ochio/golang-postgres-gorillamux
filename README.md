## API-REST -- GORILLA MUX -- POSTGRESQL (PLANTEAMIENTO DEL PROYECTO )

- In this *API* we implement the use of ***GORILLA MUX PACKAGE*** as route handler
and a database **POSRGESQL** to store the data of ***USERS*** and ***PRODUCTS***.

- These are the tables we use:

```
-- -------------- USERS TABLE  ------------------
CREATE TABLE users (
    user_id SERIAL NOT NULL,
    user_name VARCHAR(50),
    name VARCHAR(50),
    surname VARCHAR(50),
    password VARCHAR(50),
    age SMALLINT,
    active BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- -------------- PRODUCTS TABLE -----------------
CREATE TABLE products(
    product_id serial primary key,
    product_name VARCHAR(255) NOT NULL,
    description VARCHAR(500),
    price	NUMERIC,
    quantity INTEGER,
    create_date TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);
```

- The endpoints handle the values that are not entered, as null values. Since if an age is not entered, for example, said value cannot be ***"0"***, but it will be ***"NULL"***.

- The configuration of the connection to the server and database is saved in a file ***".env"*** to load that configuration we use the   ***"github.com/joho/godotenv"*** package.

- The database run in a ***docker container*** on an ***Ubuntu*** server in a virtual machine (VirtualBox Headless) whose startup and shutdown we automate through two files...

1 - ***"makefile"***
``` SHELL := /bin/bash # Use bash syntax
SERVER := <server_IP>
USER := <VM_User>
SEVERNAME := <VM_Name>
PATHDOCKERCOMPOSE := <path_to_docker-compose>.yml
ISUP := $$( ssh -o BatchMode=yes -o ConnectTimeout=1 $(SERVER) echo ok 2>&1; )

run: 
	@cd cmd && go run .

## Start the server
poserver:
	@echo "Initializing the server and postgresql container, please wait "
	@VBoxManage startvm $(SEVERNAME) --type headless
	
## Waiting until the server responds.
	@while [[ $(ISUP) != "ok" ]]; do \
		printf "#"; \
	done \

	@echo -e '\n'; \

	@echo "Server is ready...."
## Set up postgresql container	
	@ssh $(USER)@$(SERVER) 'cd $(PATHDOCKERCOMPOSE) && docker-compose up -d'

## Stop postgresql database, and poweroff the server
shutdown:
	@echo "Shutting down the server and postgresql container"
	@ssh $(USER)@$(SERVER) 'cd $(PATHDOCKERCOMPOSE) && docker-compose stop'
	@sleep 5
	@VBoxManage controlvm "$(SEVERNAME)" poweroff

```


2 - ***"docker compose.yml"***


``` 
version: "3.8"

services:

  postgres:
    image: postgres:14.5
      #restart: always
    ports:
      - "5432:5432"
    environment:
      - DATABASE_HOST=127.0.0.1
      - POSTGRES_USER=root
      - POSTGRES_PASSWORD=root
      - POSTGRES_DB=root
 
  pgadmin:
    image: dpage/pgadmin4
    environment:
      PGADMIN_DEFAULT_EMAIL: "user@user.com"
      PGADMIN_DEFAULT_PASSWORD: "password"
    ports:
      - "80:80"
    depends_on:
      - postgres    
```

- we also have a docker container that runs a ***"pgadmin4 "*** instance on localhost:80.

- Entry and visualization of the data we use templates html, css and js.



## MORE TO DO 
- testing
- midellware
- jwt
- authentication
- add redis and mongodb to manage nosql data
- problems & solutions
  - handle null values in Insert & Update 
- make a retry limit to try connect to the database(course storage)
