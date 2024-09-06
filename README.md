# GIN API REST

## Description
This is a boilerplate and learning project, developed in Go - using compiler version 1.9 - and the Postgres database. The project depends on several libraries and frameworks, which are listed in the go.sum and go.mod files.

## Requirements
- Go 1.9 or higher
- Postgres installed and configured

## Installation
To install the project, follow these steps:

- Clone the repository using the command git clone <repository URL>
- Run the command go run main.go to start the project

## Postgres Configuration
To configure Postgres, you'll need to create a database and a user with appropriate permissions. You can do this by executing the following commands:

```sql
CREATE DATABASE gin_api_rest;
CREATE ROLE {meu_usuario} WITH PASSWORD '{minha_senha}';
GRANT ALL PRIVILEGES ON DATABASE gin_api_rest TO {meu_usuario};
```

## Environment Variables
The project uses the following environment variables:

| Variable      | Desription                        |
|:--------------|:----------------------------------|
| PG_HOST:      | Postgres server address           |
| PG_PORT:      | Postgres server port              |
| PG_USER:      | Postgres user                     |
| PG_PASSWORD:  | Postgres password                 |


## Dependencies
The project depends on the following libraries and frameworks, which are listed in the go.sum and go.mod files.

## DB Setup and Migrations
The `migrations` folder at the root contains the migration files necessary for the updated functionality of the API. The main.go file is responsible for executing these migrations using `AutoMigrate`.

## Testing
The testing module is under development.