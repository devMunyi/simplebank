# 1. Design DB schema and generate SQL code with dbdiagram.io
- site url: https://dbdiagram.io/

# 2. Setup development environment on Linux: Go + VSCode + Docker + Make + Sqlc

- sudo apt install make, => for development purpose
- sudo apt snap install go --classic, => installing golang
- go version, => confirming golang installation
- sudo snap install sqlc, => for generating golang code from sql queries
- sqlc version, => confirming sqlc installed version
- sqlc help, => to get helpful commands we can use

# 3. Install & use Docker + Postgres + TablePlus to create DB schema
- docker pull postgres:16-alpine, => install postgres from docker image, we go with alpine version as it's relatively small.
- start a postgres instance or container:
    docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

    syntax used:

    docker run --name <container_name> -p <host_ports:container_ports> -e  <environment_variable>  -d <image>:<tag>

    i) --name: used to specify container name
    ii) -e: used when declaring environment variable
    iii) -d: instruct docker to start container in the background or   terminal detach mode
    - mapping ports don't necessary be the same but it's common to use same ports for easy rememberance.


    - Run command in container:
    => docker exec -it <container_name_or_id> <command> [args]

    e.g, logging into postgresql
    => docker exec -it postgres16 psql -U root

    - view container logs
    => docker logs <container_name_or_id>

    - Listing all docker containers:
    => dokcer ps -a

    - Stoping a container:
    => docker stop <container_name_or_id>

    - Starting a container:
    => docker start <container_name_or_id>

    - Remove a container completely:
    => docker rm <container_name_or_id>

    - Accessing container shell which allows doing all standard linux commands:
    => docker exec -it postgres16 /bin/sh

    - Creating simplebank database from cli using postgres createdb CLI  tool:
    => createdb --user=root --owner=root simplebank

    - Checking history for certain command from terminal:
    => history | grep "docker run"

# 4. How to write & run database migration in Golang
- Install golang-migrate using following guide:
    => https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- Once installed check version as follows:
    => migrate -version
- Create migration as follows:
    => migrate create -ext sql -dir ./db/migration -seq init_schema
- Paste previously generated sql schema into 000001_init_schema.up.sql file

- Running migration up into simple_bank database:
    => migrate -path db/migration -database
    "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

- Access the simple_bank db via docker command to check created tables/schemas:
    => docker exec -it postgres16 psql -U root simple_bank

- Checking for created tables:
 \dt
              List of relations
 Schema |       Name        | Type  | Owner 
--------+-------------------+-------+-------
 public | accounts          | table | root
 public | entries           | table | root
 public | schema_migrations | table | root
 public | transfers         | table | root
 public | users             | table | root
(5 rows)

- Running migration down from simple_bank database:
    => migrate -path db/migration -database
    "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

# 5. Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc
    a) Database/SQL
    => Very fast and straightforward
    => Manual mappings SQL fields to variables
    => Easy to make mistakes, not caught until runtime

    b) GORM
    => CRUD functions already implemented, very short production code
    => Must learn how to write queries using gorm's function
    => Run slowly on high load, benchmarks shows gets slower 3 to 5 times slow than standard library

    c) SQLX
    => Quite fast & easy to use
    => Fields mapping is via query text and struct tags
    => Failure won't occur until runtime

    d) SQLC
    => Very fast & easy to use
    => Automatic code generation
    => Catch SQL query errors before generating code

- SQLC Preferred Option:
    => Docs - sqlc.dev
    => Github - https://github.com/sqlc-dev/sqlc
    => Intallation - https://github.com/sqlc-dev/sqlc/blob/e623dc13679d0d0d7ebdf47ec6a30e22dd7e5a08/docs/overview/install.md#L4

- Create sqlc.yaml file on the project root or use command:
    => sqlc init
    
- Get startup configs settings section:
    => https://github.com/sqlc-dev/sqlc/tree/v1.4.0
    
- Settings look as follows:
version: "1"
packages:
  - name: "db"
    path: "internal/db"
    queries: "./sql/query/"
    schema: "./sql/schema/"
    engine: "postgresql"
    emit_json_tags: true
    emit_prepared_queries: true
    emit_interface: false
    emit_exact_table_names: false

- Getting started with with SQLC crud can be found from Getting Started Section of this README.md:
 => https://github.com/sqlc-dev/sqlc/tree/v1.4.0

 - Run following command to generate golang code:
   sqlc generate => thre files get generated and shouldn't be edited manually

# 6. Write Golang unit tests for database CRUD with random data
- create two files:
a) account_test.go - write tests for accounts functionalities
b) main_test.go - tests entry point

- You will also need connection to db so 
- for postgres install using:
=> go get github.com/lib/pq
 
- for mysql install using:
=> go get -u github.com/go-sql-driver/mysql

- Install Assertions and Mocks Package:
=> go get -u github.com/stretchr/testify

- Run tests written for all your packages:
=> go test -v -cover ./...

# 7. A clean way to implement database transaction in Golang
- Db Transaction: A single unit of work often made up of multiple db        operations.

- Why it matters?
 => provide reliable and consistent unit of work even in case of system failure
 => provide isolation for concurrent db operations

# 8. DB transaction lock & How to handle deadlock in Golang


# 9. 