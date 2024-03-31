# Use a imagem base do Go
FROM golang:latest

# Define as variáveis de ambiente CGO_ENABLED e GOOS
ENV CGO_ENABLED=1
ENV GOOS=linux

# Copie o código fonte para dentro do contêiner
COPY . /app

# Define o diretório de trabalho
WORKDIR /app

# Compila o código Go
RUN go build -o main .

# Comando padrão a ser executado quando o contêiner for iniciado
CMD ["./main"]
