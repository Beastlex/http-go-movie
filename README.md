## Setup env variable

```console

export GREENLIGHT_DB_DSN='postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable'

```

## Run temp postgresql

```console

docker-compose up -d

```

## Create migration

Create migration with golang-migrate tool

```console
migrate create -seq -ext=.sql -dir=./migrations create_movies_table
```