package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/url"

	pb "main/grpcServer"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var ctx = context.Background()
var db *sql.DB

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":3001"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func mysqlConnect() {
	// Cambia las credenciales según tu configuración de MySQL
	dbUser := "root"
	dbPassword := "luE=hEV%s;j&a#op"
	dbHost := "34.86.30.206"
	dbPort := "3306"
	dbName := "tareas"

	// Escapar la contraseña
	escapedPassword := url.QueryEscape(dbPassword)

	// Construir la cadena de conexión DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, escapedPassword, dbHost, dbPort, dbName)

	// Deshacer el escape de la contraseña en la cadena de conexión
	dsn, _ = url.QueryUnescape(dsn)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Conexión a MySQL exitosa")
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí la info de: ", in.GetName())
	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Year:  in.GetYear(),
		Rank:  in.GetRank(),
	}
	insertMySQL(data)
	return &pb.ReplyInfo{Info: "Hola cliente, recibì la info."}, nil
}

func insertMySQL(voto Data) {
	// Prepara la consulta SQL para la inserción en MySQL
	query := "INSERT INTO tarea (name, album, yearr, rankk) VALUES (?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, query, voto.Name, voto.Album, voto.Year, voto.Rank)
	if err != nil {
		log.Println("Error al insertar en MySQL:", err)
	}
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	mysqlConnect()

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
