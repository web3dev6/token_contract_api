# Token Transaction

This is a Go REST API service that manages users, sessions and blockchain-write-transaction entries on a Postgres DB.

Note: Make sure you have docker and golang-migrate (https://github.com/golang-migrate) on your system.

Run dockerized postgres service:
```shell
make postgres
```

Create DB in postgres:
```shell
make createdb
```

Run migration to create tables in db:
```shell
make migrateup
```

Run migration to create tables in db:
```shell
make migrateup
```

Run tests:
```shell
make test
```

Run tests:
```shell
make test
```

Run server:
```shell
make server
```
