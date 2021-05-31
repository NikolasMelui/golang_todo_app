# TODO GO APPLICATION

Need to install [migrate](https://github.com/golang-migrate/migrate) first

- CREATE MIGRATION:

```bash
  migrate create -ext sql -dir ./schema -seq init
```

- UP:

```bash
  migrate -path ./schema -database 'postgres://username:password@127.0.0.1:5432/database?sslmode=disable' up
```

- DOWN:

```bash
  migrate -path ./schema -database 'postgres://username:password@127.0.0.1:5432/database?sslmode=disable' down
```
