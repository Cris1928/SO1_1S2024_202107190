/*package main

import (
	"context"
	"fmt"
	"log"
	pb "main/grpcServer"
	"net"
	"os"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

// thunder client
type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	topic         = "votes-submitted"
	brokerAddress = "my-cluster-kafka-bootstrap.kafka.svc.cluster-p2so1.local:9092"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {

	fmt.Println(in)
	msg := pb.ReplyInfo{Info: "creado"}
	produce(ctx, in)
	return &msg, nil
}
func produce(ctx context.Context, in *pb.RequestId) {
	l := log.New(os.Stdout, "Kafka Escribiendo: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(in.GetRank()),
		Value: []byte("{\"name\": \" \"" + in.GetName() + "\", \"album\":\"" +
			in.GetAlbum() + "\", \"year\": \"" + in.GetYear() +
			"\", \"rank\":" + "\"" + in.GetRank() + "\" }"),
	})
	if err != nil {
		panic("No se pudo enviar el mensaje: " + err.Error())
	}
}
func main() {
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("funcionando")
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
*/
/*
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	pb "main/grpcServer"

	"google.golang.org/grpc"
)

const (
	topic         = "votes-submitted"
	brokerAddress = "my-cluster-kafka-bootstrap.kafka.svc.cluster-p2so1.local:9092"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

func (s *server) ReturnInfo(ctx context.Context, req *pb.RequestId) (*pb.ReplyInfo, error) {
	var topic string = "votes-submitted"
	// Aquí procesarías la información recibida del cliente
	fmt.Printf("Recibido: Name=%s, Album=%s, Year=%s, Rank=%s\n", req.Name, req.Album, req.Year, req.Rank)

	// Crear un productor de Kafka
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokerAddress})
	if err != nil {
		return nil, err
	}
	defer p.Close()

	// Enviar el mensaje al topic de Kafka
	message := fmt.Sprintf("%s, %s, %s, %s", req.Name, req.Album, req.Year, req.Rank)
	deliveryChan := make(chan kafka.Event)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)
	if err != nil {
		return nil, err
	}

	// Esperar a que el mensaje sea entregado correctamente
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return nil, m.TopicPartition.Error
	}

	// Retornar una respuesta al cliente gRPC
	return &pb.ReplyInfo{Info: "Información enviada a Kafka correctamente"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})
	log.Println("Server listening on port :3001")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
*/
/*
package main

import (
	"context"
	"fmt"
	"log"
	pb "main/grpcServer"
	"net"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

const (
	topic         = "votes-submitted"
	brokerAddress = "my-cluster-kafka-bootstrap.kafka.svc.so1-p2024.local:9092"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

func (s *server) ReturnInfo(ctx context.Context, req *pb.RequestId) (*pb.ReplyInfo, error) {
	// Aquí procesas la información recibida del cliente gRPC
	// y la envías al servicio de Kafka
	llave := req.Rank

	data := fmt.Sprintf("Name: %s, Album: %s, Year: %s, Rank: %s",
		req.Name, req.Album, req.Year, req.Rank)

	// Configuración del cliente de Kafka
	config := kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	}

	// Crear un nuevo escritor de Kafka
	writer := kafka.NewWriter(config)

	// Escribir el mensaje en Kafka
	err := writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(llave),
		Value: []byte(data),
	})
	if err != nil {
		log.Println("Error al enviar mensaje a Kafka:", err)
		return nil, err
	}

	log.Println("Mensaje enviado a Kafka:", data)

	return &pb.ReplyInfo{Info: "Información enviada a Kafka"}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("funcionando")
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})
	log.Println("Servidor gRPC escuchando en el puerto 3001")
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
*/
/*
package main

import (
	"context"
	"fmt"
	"log"
	pb "main/grpcServer"
	"net"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

const (
	port          = ":3001"
	topic         = "votes-submitted"
	brokerAddress = "my-cluster-kafka-bootstrap.kafka.svc.so1-p2024.local:9092"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

func (s *server) ReturnInfo(ctx context.Context, req *pb.RequestId) (*pb.ReplyInfo, error) {
	// Enviar la información al servicio de Kafka
	err := sendToKafka(req)
	if err != nil {
		log.Printf("Error al enviar a Kafka: %v", err)
		return nil, err
	}

	return &pb.ReplyInfo{Info: "Información recibida y enviada a Kafka correctamente"}, nil
}

func sendToKafka(req *pb.RequestId) error {
	// Crear una nueva conexión al broker de Kafka
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Crear un escritor para el topic especificado
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	// Enviar el mensaje al topic de Kafka
	err = writer.WriteMessages(context.TODO(),
		kafka.Message{
			Key:   []byte(req.Name),
			Value: []byte(fmt.Sprintf("Nombre: %s, Album: %s, Year: %s, Rank: %s", req.Name, req.Album, req.Year, req.Rank)),
		},
	)
	if err != nil {
		return err
	}

	// Cerrar el escritor
	err = writer.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Crear un listener en el puerto especificado
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al escuchar en el puerto %s: %v", port, err)
	}

	// Crear un nuevo servidor gRPC
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	log.Printf("Servidor gRPC escuchando en el puerto %s", port)

	// Iniciar el servidor gRPC
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
	}
}
*/

package main

import (
	"context"
	"fmt"
	"log"
	pb "main/grpcServer"
	"net"
	"os"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

// thunder client
type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	topic         = "votes-submitted"
	brokerAddress = "my-cluster-kafka-bootstrap:9092"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	//func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí de cliente: ", in.GetName())
	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Year:  in.GetYear(),
		Rank:  in.GetRank(),
	}
	llave := data.Rank
	datas := fmt.Sprintf("Name: %s, Album: %s, Year: %s, Rank: %s",
		data.Name, data.Album, data.Year, data.Rank)
	l := log.New(os.Stdout, "Kafka Escribiendo: ", 0)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})
	err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(llave),
		Value: []byte(datas),
	})
	if err != nil {
		log.Println("No se pudo enviar el mensaje", datas)

		return &pb.ReplyInfo{Info: "No se pudo enviar el mensaje "}, nil
	}
	log.Println("Mensaje enviado a Kafka:", datas)

	return &pb.ReplyInfo{Info: "Información enviada a Kafka"}, nil
}
func main() {
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("funcionando")
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}

/*
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "main/grpcServer"

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

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí la info de: ", in.GetName())
	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Year:  in.GetYear(),
		Rank:  in.GetRank(),
	}
	fmt.Println(data)
	return &pb.ReplyInfo{Info: "Hola cliente, recibì la info."}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})
	log.Println("funcionando")
	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
*/
