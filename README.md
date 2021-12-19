# test-api-bs

TEST-API-bondstate

please run the both migration file first to setup the db

I am using postgres db for storage.

First Migration File :

```
psql postgres -f migration.sql
```

migration_test.sql

```
psql postgres -f migration_test.sql
```

please note running migration files first is important

to run this test-api

`DBHOST=localhost DBUSER=postgres DBPASS=root DBNAME=bondstate_db DBPORT=5432 PORT=9090 go run main.go`

to test the service
`go test ./... -v`

# end-points

/users -> GetAll users from db

/users/:id -> Get user from db by id

/portfolio/:id/entry -> save a entry in portfolio [id provided]

/user/:id/portfolio -> Get portfolio info of a user

# logic

user and portfolio are 1-1 relations

user can add many entries in portfolio
