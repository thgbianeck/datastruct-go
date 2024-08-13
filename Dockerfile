# Usa a imagem oficial do Go como base
FROM golang:latest

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos do projeto para o contêiner
COPY . .

# Baixa as dependências
# RUN go mod tidy

# Define o comando padrão para manter o contêiner em execução
CMD ["tail", "-f", "/dev/null"]