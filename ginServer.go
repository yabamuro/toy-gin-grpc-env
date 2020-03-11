package main

import (
	pb "ginApl"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"unicode/utf8"
)

const (
	address string = "localhost:50051"
	logfile string = "/var/log/apl/ginApl.log"
)

// LoggingSetting
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings(logfile)
	log.Println("ginServer Start")
	router := gin.Default()
	router.GET("/gin", func(c *gin.Context) {
		// get the query string form Request URL
		query := c.Request.URL.Query().Encode()
		queryLen := utf8.RuneCountInString(query)
		// Eliminate "=" if the query string ends with it
		if query[queryLen-1] == '=' {
			query = query[:queryLen-1]
		}
		log.Printf("REV(HTTPS): %s", query)
		//Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		con := pb.NewGrpcClient(conn)
		// Contact the server and print out its response.
		log.Printf("SND(gRPC): %s", query)
		r, err := con.Send(context.Background(), &pb.GrpcRequest{Query: query})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("RCV(gRPC): %s", r.Message)
		// Return the response value from grpcServer
		c.String(http.StatusOK, r.Message+"\n")
		log.Printf("SND(HTTPS): %s", r.Message)
	})
	s := &http.Server{
		Addr:           ":9000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
