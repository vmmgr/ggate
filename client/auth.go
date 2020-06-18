package client

import (
	"context"
	pb "github.com/vmmgr/controller/proto/proto-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func GenerateTokenClient(user, pass string) *AuthResult {
	log.Println(getServerAddress())
	conn, err := grpc.Dial(getServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Not connect; %v\n", err)
	}
	defer conn.Close()
	c := pb.NewControllerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GenerateToken(ctx, &pb.UserData{Name: user, Pass: pass})
	if err != nil {
		log.Printf("Not connect; %v\n", err)
	}

	return &AuthResult{Result: r.Status, Token: r.Data1}
}

func CheckTokenClient(token string) Result {
	conn, err := grpc.Dial(getServerAddress(), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Printf("Not connect; %v\n", err)
	}
	defer conn.Close()
	c := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	r, err := c.GetToken(ctx, &pb.TokenData{Token: token})
	if err != nil {
		log.Printf("Not connect; %v\n", err)
	}

	return Result{
		Result: r.Status,
		Info:   r.Info,
	}
}

func DeleteTokenClient(token string) Result {
	conn, err := grpc.Dial(getServerAddress(), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Printf("Not connect; %v\n", err)
	}
	defer conn.Close()
	c := pb.NewControllerClient(conn)
	header := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	r, err := c.DeleteToken(ctx, &pb.Null{})
	if err != nil {
		log.Printf("Not connect; %v\n", err)
	}
	return Result{
		Result: r.Status,
		Info:   r.Info,
	}
}

/*
func GetAllTokenClient(token string) bool {
	conn, err := grpc.Dial(getServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//r, err := c.GetAllToken(ctx, &pb.Base{Token: token})
	//if err != nil {
	//	log.Println("could not greet: %v", err)
	//}
	//if r.Status {
	//	return true
	//} else {
	//	return false
	//}
}
*/
