package main

import (
	"bufio"
	pb "chat/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial("localhost:33244", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connected:%v", err)
	}
	defer conn.Close()
	sender := pb.NewSendingClient(conn)
	s, err := sender.Run(context.Background())
	if err != nil {
		log.Fatalf("SendingClient Run:%v", err)
	}
	bye := false
	fmt.Println("sending start")
	scanner := bufio.NewScanner(os.Stdin)
	for !bye {
		scanner.Scan()
		words := scanner.Text()
		if words == "bye" {
			bye = true
			l, _ := s.CloseAndRecv()
			fmt.Println(l)
		}
		s.Send(&pb.TalkingWords{words})
	}
	fmt.Println("sending done")
	receiver := pb.NewReceivingClient(conn)
	length := 10
	r, err := receiver.Run(context.Background(), &pb.TalkingLength{int32(length)})
	if err != nil {
		log.Fatalf("SendingClient Run:%v", err)
	}
	fmt.Println("receivng start")
	for i := 0; i < length; i++ {
		words, _ := r.Recv()
		fmt.Println("round:", i+1, "receiving:", words.GetWords())
	}
	fmt.Println("receiving done")
	talker := pb.NewTalkingClient(conn)
	t, err := talker.Run(context.Background())
	bye = false
	fmt.Println("talking start")
         bye = false
	for !bye {
		scanner.Scan()
		sending := scanner.Text()
		t.Send(&pb.TalkingWords{sending})
		receiving, err := t.Recv()
                fmt.Println("receiving:", receiving.Words)
		if err != nil {
			log.Fatal("taling receing")
		}
		if "bye" == receiving.GetWords() {
			bye = true
		        t.Send(&pb.TalkingWords{"bye"})
		}
	}
	fmt.Println("talking done")
}
