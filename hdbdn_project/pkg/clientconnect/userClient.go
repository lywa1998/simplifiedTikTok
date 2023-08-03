package clientconnect

import (
	"github.com/hdbdn77/simplifiedTikTok/pkg/userService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserRegisterChan chan userService.RegisterServiceClient
var UserLoginChan chan userService.LoginServiceClient
var UserChan chan userService.UserServiceClient
var addr = ":8002"

func init() {
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	UserRegisterChan = make(chan userService.RegisterServiceClient, 10)
	UserLoginChan = make(chan userService.LoginServiceClient, 10)
	UserChan = make(chan userService.UserServiceClient, 10)
	for i := 0; i < 10; i++ {
		UserRegisterChan <- userService.NewRegisterServiceClient(conn)
		UserLoginChan <- userService.NewLoginServiceClient(conn)
		UserChan <- userService.NewUserServiceClient(conn)
	}
}