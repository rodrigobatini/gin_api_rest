# GIN API REST

## Description
This is a boilerplate and learning project, developed in Go - using compiler version 1.9 - and the Postgres database. The project depends on several libraries and frameworks, which are listed in the go.sum and go.mod files.

## Requirements
- docker-engine and docker-compose

## Installation
To install the project, follow these steps:

- Clone the repository using the command git clone <repository URL>
- Create a .env file with the following content

    | Variable    | Description |
    |:------------|:------------|
    | PG_HOST     | postgres    |
    | PG_PORT     | 5432        |
    | PG_USER     | postgres    |
    | PG_PASSWORD | postgres    |
    | DB_NAME     | gin_api_rest|

    - You can adjust these settings as you prefer, as long as they are reflected in the `docker-compose.yml` under the postgres section. 

- Run the command `docker-compose up` to run the application on foreground (use `-d` for running on background)
    - depending on which version and how you installed docker-compose, you may need to run as `docker compose up`

## DB Setup and Migrations
The `migrations` folder at the root contains the migration files necessary for the updated functionality of the API. The main.go file is responsible for executing these migrations using `AutoMigrate`.

## Testing
The testing module is under development.