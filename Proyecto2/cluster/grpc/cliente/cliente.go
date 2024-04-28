/*package main

import (
	"context"
	"fmt"
	"log"
	pb "main/grpcCliente"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func insertData(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	rank := Data{
		Name:  data["name"],
		Album: data["album"],
		Year:  data["year"],
		Rank:  data["rank"],
	}

	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Name:  rank.Name,
		Album: rank.Album,
		Year:  rank.Year,
		Rank:  rank.Rank,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Respuesta del server " + ret.GetInfo())

	return nil
}

func main() {
	app := fiber.New()
	app.Post("/insert", insertData)

	fmt.Println("Server started successfully.")
	fmt.Println("Listening on port :3000...")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}


package main

import (
	"context"
	"fmt"
	"log"
	pb "main/grpcCliente" // nombre_proyecto/carpeta

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

var ctx = context.Background()

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func insertData(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	voto := Data{
		Name:  data["name"],
		Album: data["album"],
		Year:  data["year"],
		Rank:  data["rank"],
	}
	fmt.Println(voto)
	go sendServer(voto)
	return nil
}

func sendServer(voto Data) {
	//conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Name:  voto.Name,
		Album: voto.Album,
		Year:  voto.Year,
		Rank:  voto.Rank,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Respuesta del server " + ret.GetInfo())
}

func main() {
	app := fiber.New()

	app.Get("/grpc", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"res": "todo bien",
		})
	})
	app.Post("/grpc/insert", insertData)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
*/

package main

import (
	"context"
	"fmt"
	"log"
	pb "main/grpcCliente"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

var ctx = context.Background()

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func insertData(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	voto := Data{
		Name:  data["name"],
		Album: data["album"],
		Year:  data["year"],
		Rank:  data["rank"],
	}
	fmt.Println(voto)
	go sendServer(voto)
	return nil
}

func sendServer(voto Data) {
	//conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("1")
		log.Fatalln(err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("2")
			log.Fatalln(err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Name:  voto.Name,
		Album: voto.Album,
		Year:  voto.Year,
		Rank:  voto.Rank,
	})
	if err != nil {
		fmt.Println("3")
		log.Fatalln(err)
	}

	fmt.Println("Respuesta del server " + ret.GetInfo())
}

func main() {
	app := fiber.New()

	app.Get("/grpc", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"res": "todo bien",
		})
	})
	app.Post("/grpc/insert", insertData)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
