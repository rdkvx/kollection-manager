# ===============================
# Etapa 1: Build da aplicação Go
# ===============================
FROM golang:1.22-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /src

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código-fonte
COPY . .

# Compila o binário (sem CGO e para Linux amd64)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/kollection-manager main.go


# ===============================
# Etapa 2: Imagem final mínima
# ===============================
FROM alpine:3.19

# Cria diretório de trabalho
WORKDIR /app

# Instala curl e Bash
RUN apk add --no-cache bash curl

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/kollection-manager .

# Define a porta que o container vai expor
EXPOSE 8080

# Comando de inicialização
ENTRYPOINT ["./kollection-manager"]
CMD ["--port=8080"]
