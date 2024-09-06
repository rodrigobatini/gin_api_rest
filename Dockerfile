# Use uma imagem base do Golang
FROM golang:1.19.8 AS builder

# Define o diretório de trabalho
WORKDIR /

# Copia os arquivos go mod e go sum
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod download

# Copia o código-fonte da aplicação
COPY . .

# Compila a aplicação
RUN go build -o main .

# Cria uma nova imagem base para a aplicação
FROM debian:bullseye-slim

# Define o diretório de trabalho
WORKDIR /

# Instala dependências e dockerize
RUN apt-get update && apt-get install -y \
    curl \
    && curl -sSL https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz \
    | tar -C /usr/local/bin/ -xzv dockerize \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Copia o binário compilado da imagem builder
COPY --from=builder /main .

# Define o comando para executar a aplicação com dockerize
CMD ["dockerize", "-wait", "tcp://postgres:5432", "-timeout", "60s", "./main"]