# Etapa 1: Build da aplicação
FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod  ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o backup-etl .

# Etapa 2: Container final
FROM alpine:3.20

# Adiciona certificado raiz e timezone
RUN apk add --no-cache ca-certificates tzdata

# Define timezone de São Paulo
ENV TZ=America/Sao_Paulo

# Define diretório de trabalho
WORKDIR /app

# Copia apenas o binário compilado
COPY --from=builder /app/backup-etl .

# Comando padrão para rodar o app
CMD ["./backup-etl"]
