# Imagem base
FROM golang:latest

# Diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod .
COPY go.sum .

# Baixar e instalar as dependências do projeto
RUN go mod download

# Copiar o código fonte para o diretório de trabalho
COPY . .

# Compilar o código Go dentro do contêiner
RUN go build -o monitor-buckets-s3 .

# Comando padrão a ser executado quando o contêiner for iniciado
CMD ["./monitor-buckets-s3"]
