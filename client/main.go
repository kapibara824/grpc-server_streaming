package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"context"

	"io"

	counterproto "github.com/kapibara824/grpc-server_streaming/pb/counter"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8240", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect:%v\n", err)
	}
	fmt.Printf("Running gRPC client...\n")

	defer conn.Close()

	c := counterproto.NewCounterServiceClient(conn)
	for {
		fmt.Printf("Please input number:")
		num, err := input()
		if err != nil {
			log.Printf("invalid format. Please input number only.")
		}
		req := &counterproto.CounterRequest{
			Num: num,
		}
		stream, err := c.Counter(context.Background(), req)
		if err != nil {
			log.Println(err)
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)

			}
			fmt.Println(res.GetResult())
		}
	}
}

func input() (int64, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	num, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		return int64(num), err
	}
	return int64(num), nil

}
