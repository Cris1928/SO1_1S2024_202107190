#Utilizar Alpine
FROM golang:alpine 

#Directorio de trabajo
WORKDIR /app

#Copiar el código fuente
COPY . .

#Compilar el código fuente
RUN go build -o myapp

#Exponer el puerto 8000
EXPOSE 5200

#Ejecutar el binario
CMD ["./myapp"]

#Construir la imagen
#docker build -t app_golang .