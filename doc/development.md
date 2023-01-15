# Development

## Database

Migrations are managed using [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate), with the CLI tool.
If this is your first time using golang-migrate, check out the [Getting Started guide](https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md).

To install the golang-migrate CLI, follow the instructions in the [migrate CLI README](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md).

### Local development database

1. Install Mysql 5 on your machine for local development. It should use the default Postgres port of 3306.
You can use mysql database locally with docker container as follows:

```sh
docker run
   -p 3306:3306
   --env MYSQL_DATABASE=test
   --env MYSQL_ALLOW_EMPTY_PASSWORD=yes
   --name mysql
   --detach
   --rm
   mysql:5
   --character-set-server=utf8mb4
   --collation-server=utf8mb4_unicode_ci
```

### Setting up for tests

If you followed the docker setup in step 1 in local development database above, then you have one.

When running `go test ./...`, database tests will not run if you don't have mysql running

Tests use the following environment variables:

- `DATABASE_DATASOURCE` the mysql connection string, such as `root@tcp(localhost:3306)/test?parseTime=true`

### Creating a migration

To create migration:

```
devtools/create_migration.sh <title>
```

This creates two empty files in `/migrations`:

```
{version}_{title}.up.sql
{version}_{title}.down.sql
```

The two migration files are used to migrate "up" to the specified version from the previous version, and to migrate "down" to the previous version. See [golang-migrate/migrate/MIGRATIONS.md](https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md) for details.

If you are migrating for the first time. You just runs the command, make sure the connection is matched with your local development setup.

```
migrate -source file:migrations -database "mysql://root@tcp(localhost:3306)/test?parseTime=true" up
```