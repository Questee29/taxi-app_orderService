package main

import (
	"context"
	"errors"
	"io"
	"log"

	config "github.com/Questee29/taxi-app_orderService/configs"
	"github.com/Questee29/taxi-app_orderService/database"
	_ "github.com/Questee29/taxi-app_orderService/migrations"
	model "github.com/Questee29/taxi-app_orderService/models/order"
	server "github.com/Questee29/taxi-app_orderService/pkg/grpcServer"
	handlers "github.com/Questee29/taxi-app_orderService/pkg/grpcServer/handler"
	"github.com/Questee29/taxi-app_orderService/pkg/repository"
	service "github.com/Questee29/taxi-app_orderService/pkg/service"
	pb "github.com/Questee29/taxi-app_orderService/proto/protob"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	business = make(chan int32, 50) // id int
	comfort  = make(chan int32, 50)
	economy  = make(chan int32, 50)
)

func main() {
	config, err := config.LoadConfig("app", ".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	db, err := database.New()
	if err != nil {
		log.Fatalln(errors.New(`failed to load database`))
	}
	repository := repository.New(db)
	service := service.New(repository)
	grpcOrderHandler := handlers.NewOrderHandler(service)
	grpcServ := server.NewServer(server.Deps{
		OrderHandler: grpcOrderHandler,
	})

	go func() {
		if err := grpcServ.ListenAndServe(config.Server.Port); err != nil {
			log.Printf("grpc ListenAndServe err %s", err)
		}
	}()

	conn, err := grpc.Dial(handlers.DriverAdress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("grpc Client Dial err %s", err)
	}
	defer conn.Close()
	client := pb.NewOrderGrpcServiceClient(conn)
	in := &pb.FindDriverRequest{Userid: 2}
	stream, err := client.FindDriver(context.Background(), in)
	if err != nil {
		log.Printf("grpc Cstream err %s", err)
	}

	done := make(chan bool)
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Printf("can not recieve %v", err)
				return
			}

			driversPB := resp.GetDriver()
			for _, driverPB := range driversPB {
				driver := model.FreeDriver{
					DriverID: driverPB.GetDriverID(),
					TaxiType: driverPB.GetType().String(),
				}
				log.Printf("recieved from driverService %v", driver)
				go func(driver model.FreeDriver) {
					switch driver.TaxiType {
					case "comfort":
						log.Println("added comfort")
						comfort <- driver.DriverID
					case "business":
						log.Println("added business")
						business <- driver.DriverID
					case "economy":
						log.Println("added economy")
						economy <- driver.DriverID
					}

				}(driver)

			}

		}
	}()

	<-done
	log.Println("finished")

}
