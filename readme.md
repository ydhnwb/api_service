## Web Service

Lorem ipsum dolor sit amet consectetur.

## Setup repo

### Branching

Since this might be me alone, well please use the itemId of github project as branch when develop its task, check url address bar for it. Example itemId is 1230, well then do this:

- git checkout main
- git pull origin main
- git checkout -b "#1230"
- start working
- push your change using git push origin '#1230'

There are more branching strategy with pool but I don't want to talk about it right now

### Dependencies

Most dependencies that needed inside this repo is available on go.mod.  
You can install it using this command:

```
go get
```

and

```
go mod install
```

### Configure hot reload

I am using AIR to make it hot reload when developing. If you are using linux, please take a look at `.air.toml` file and change some parameters there.

### Migration tool

I am using https://github.com/golang-migrate/ package.
You can install it using the command:

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
```

after success, try to run it with `migrate` and then enter. It should show all the commands
of using this migrate cli tool. Please refer to this article `https://medium.com/@aldinofrizal/golang-migrate-8990135cdc6`

### How I add migration?

Here is the cli command template:

```
migrate create -ext sql -dir migrations create_table_users_and_devices
```

`migrate` is cli-app we installed. `create` means we will create a migration, it should create at leats two type of file (UP and DOWN), `-ext` is extension, we use postgresql so it will be sql (If you use mongodb, it migth be `json`). `-dir` is where we want to generate the migration file, please put it under `migrations` folder AND DO NOT CHANGE IT!. `create-table-users` is a name what migration we gonna do, we do not have convention for that so please take a look at previous data.

## Up the migration

The url for database can be variative, but for now, since I don't use containerized thing like docker, so this is the db url and change it with your local environment:

```
postgres://user:password@host:port/dbname?YOUR_QUERY_OPTIONAL
```

example of url database: `postgres://postgres:yudhanewbie@localhost:5432/api_service?sslmode=disable`

and run this command

```
migrate -database "YOU_URL_TO_DATABASE" -path migrations up
```

### How do I cancel the migration?

```
migrate -database "YOU_URL_TO_DATABASE" -path migrations down
```

## Personal target

- Please and please and please don't use GORM or any other ORM :v
- maintain the singleton of the db instance. (Leads too much database connection)
- easy, readable error response
- use HTTP status code as main source of truth rather than other field like 'status' etc
- less magic is better

## Mandatory guide.

- test coverage at least 80%
- lorem
- ipsum
