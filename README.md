# GIN API REST

## Descrição

Este é um projeto boilerplate e de aprendizado, feito em Go - que utiliza a versão 1.9 do compilador - e o banco de dados Postgres. O projeto depende de várias bibliotecas e frameworks, que estão listadas no arquivo go.sum e go.mod.

## Requisitos

- Go 1.9 ou superior
- Postgres instalado e configurado

## Instalação

Para instalar o projeto, siga os passos abaixo:

- Clone o repositório utilizando o comando git clone <URL do repositório>
- Execute o comando ```go run main.go``` para iniciar o projeto

## Configuração do Postgres

Para configurar o Postgres, você precisará criar um banco de dados e um usuário com permissões adequadas. Você pode fazer isso executando os seguintes comandos:

```sql
CREATE DATABASE gin_api_rest;
CREATE ROLE {meu_usuario} WITH PASSWORD '{minha_senha}';
GRANT ALL PRIVILEGES ON DATABASE gin_api_rest TO {meu_usuario};
```

## Variáveis de Ambiente
O projeto utiliza as seguintes variáveis de ambiente:

| Variável      | Descrição                         |
|:--------------|:----------------------------------|
| PG_HOST:      | o endereço do servidor Postgres   |
| PG_PORT:      | a porta do servidor Postgres      |
| PG_USER:      | o nome do usuário Postgres        |
| PG_PASSWORD:  | a senha do usuário Postgres       |


## Dependências
O projeto depende das seguintes bibliotecas e frameworks, que estão listadas no arquivo go.sum e go.mod

## DB Setup e Migrations
Atualmente precisa criar o banco na mão, setup e migrations são WiP