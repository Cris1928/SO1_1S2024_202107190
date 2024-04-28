package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"strconv"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"

	//"github.com/KromDaniel/rejonson"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	topic         = "votes-submitted"
	brokerAddress = "my-cluster-kafka-bootstrap:9092"
)

var (
	usr      = "admin"
	pwd      = "HQ03ilp&dOtj"
	host     = "34.172.49.7"
	port     = 27017
	database = "database"
)
var goRedisClient = redis.NewClient(&redis.Options{
	Addr: "34.121.60.197:6379",
	//Password: "",
	//DB: 0,
})

//defer goRedisClient.Close()

type Data struct {
	Name  string `json:"Name"`
	Album string `json:"Album"`
	Year  string `json:"Year"`
	Rank  string `json:"Rank"`
}

func consume(ctx context.Context) {
	l := log.New(os.Stdout, "Kafka escuchando: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		Logger:  l,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("No se pudo leer el mensaje: " + err.Error())
		}

		fmt.Println("Mensaje recibido: ", string(msg.Value))

		go insertRedis(string(msg.Value))
		go insertMongo(ctx, string(msg.Value))
	}
}

func insertRedis(message string) {
	fmt.Println(message)
	var data Data

	song := parseSong(message)
	jsonData, q := json.Marshal(song)
	if q != nil {
		fmt.Println("Error al convertir a JSON:", q)
		return
	}

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Printf("Error al decodificar el JSON: %v\n", err)
		return
	}

	// Convertir el campo "Rank" a un valor numérico
	rank, err := strconv.Atoi(data.Rank)
	if err != nil {
		log.Printf("Error al convertir el campo 'Rank' a un valor numérico: %v\n", err)
		return
	}

	// Insertar en Redis utilizando el campo "Album" como clave y "Rank" como valor numérico
	err = goRedisClient.HSet("kafka_message", data.Album, rank).Err()
	if err != nil {
		log.Printf("Error al insertar en Redis: %v\n", err)
		return
	}
	log.Println("Mensaje insertado en Redis correctamente.")
}

/*unc insertRedis(body string) {






	client := rejonson.ExtendClient(goRedisClient)
	client.JsonArrAppend("redis-service", ".", body)
	fmt.Println("dato insertado en redis correctamente")

}*/

func insertMongo(ctx context.Context, body string) {
	var data Data
	json.Unmarshal([]byte(body), &data)
	var collection = GetCollection("data")
	var err error
	_, err = collection.InsertOne(ctx, data)

	if err != nil {
		fmt.Println("1")
		fmt.Println(err.Error())
	}
	fmt.Println("dato insertado en mongo correctamente")
}

func GetCollection(collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}

func parseSong(input string) Data {
	song := Data{}
	fields := strings.Split(input, ", ")
	for _, field := range fields {
		parts := strings.Split(field, ": ")
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		switch key {
		case "Name":
			song.Name = value
		case "Album":
			song.Album = value
		case "Year":
			fmt.Sscanf(value, "%d", &song.Year)
		case "Rank":
			fmt.Sscanf(value, "%d", &song.Rank)
		}
	}
	return song
}

func main() {
	ctx := context.Background()
	consume(ctx)
}
