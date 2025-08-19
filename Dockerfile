# Etapa 1: Build da aplicação
FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod  ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o backup-etl .

FROM alpine:3.20

# Adiciona certificado raiz (caso o app faça chamadas HTTPS)
RUN apk add --no-cache ca-certificates

# Define diretório de trabalho
WORKDIR /app

# Copia apenas o binário compilado
COPY --from=builder /app/backup-etl .

# Comando padrão para rodar o app
CMD ["./backup-etl"]
