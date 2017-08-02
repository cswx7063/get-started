package main

import(
    "fmt"
    "net"
    "log"
    "io"
    crand "math/rand"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    pb "chat/proto"
)
type sendingServer struct{}
func (s *sendingServer)Run(stream pb.Sending_RunServer)error{
        length := 0
        for{
           words, err := stream.Recv()
           if err == io.EOF{
                return stream.SendAndClose(&pb.TalkingLength{int32(length)})
           }
           length += len(words.GetWords())
           fmt.Println(words.GetWords())
      }
}

type receivingServer struct{}
func (s *receivingServer)Run(in *pb.TalkingLength, stream pb.Receiving_RunServer)error{
     batch := 10
     words := make([]byte, batch)
     block := int(in.Length) / batch
     for i:=0; i<block; i++{
          crand.Read(words)
          fmt.Println(string(words))
          stream.Send(&pb.TalkingWords{string(words)})
     }   
     return nil
}

type talkingServer struct{}

func (s *talkingServer)Run(stream pb.Talking_RunServer)error{
     for{
           words, err := stream.Recv()
           if err != nil{
              return nil
           }
           if "bye" == words.GetWords(){
                 stream.Send(&pb.TalkingWords{"bye!"}) 
                 return nil
           }
           fmt.Println("receing:", words.GetWords()) 
           var resp string
           fmt.Scanf("sending:%s", &resp)
           stream.Send(&pb.TalkingWords{resp}) 
     }  
     return nil 
}
func main(){
        fmt.Println("starting")
        lis, err := net.Listen("tcp", ":33244")
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterSendingServer(s, &sendingServer{})
        pb.RegisterReceivingServer(s, &receivingServer{})
        pb.RegisterTalkingServer(s, &talkingServer{})
        // Register reflection service on gRPC server.
        reflection.Register(s)
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}

