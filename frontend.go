package main

import (
	"fmt"
	"log"
	"net/http"
        "time"
        "encoding/json"
        

	"golang.org/x/net/context"
	grpc "github.com/Romero027/grpc-go"

	echo "github.com/Romero027/echo-server-grpc/pb"
)


func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	start := time.Now()
 	defer func() {
  		in, _ := json.Marshal(req)
  		out, _ := json.Marshal(reply)
  		inStr, outStr := string(in), string(out)
  		duration := int64(time.Since(start).Microseconds())

  		log.Println("grpc", method, "in", inStr, "out", outStr, "err", err, "duration/us", duration)

 }()

	return invoker(ctx, method, req, reply, cc, opts...)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s\n", request.URL.String())

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure(),grpc.WithUnaryInterceptor(UnaryClientInterceptor),)
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := echo.NewEchoServiceClient(conn)

	message := echo.Msg{
		Body: request.URL.String(),
		// Body: "Hello",
	}

	response, err := c.Echo(context.Background(), &message)
	if err != nil {
		log.Fatalf("Erro when calling echo: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
	fmt.Fprintf(writer, "Echo request finished! Length of the request is %d\n", len(response.Body))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
