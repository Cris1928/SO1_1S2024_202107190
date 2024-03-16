# Usa la imagen oficial de Golang como base
FROM golang:alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del cliente a la imagen
COPY . .

# Compila el código Go
RUN go build -o app

# Expone el puerto necesario
EXPOSE 5200

# Comando para ejecutar la aplicación cuando se inicie el contenedor
CMD ["./app"]