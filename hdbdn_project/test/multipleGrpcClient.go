package test

// // 定义客户端接口
// type Client interface {
// 	Call(ctx context.Context, req interface{}) (resp interface{}, err error)
// }

// // User服务客户端
// type UserClient struct {
// 	conn   *grpc.ClientConn
// 	client pb.UserServiceClient
// }

// func NewUserClient() *UserClient {
// 	conn, err := grpc.Dial("userservice.example.com")
// // 如果err不为空,返回错误处理

// return &UserClient{conn: conn, client: pb.NewUserServiceClient(conn)}
// }

// func (c *UserClient) Call(ctx context.Context, req interface{}) (resp interface{}, err error) {
// 	// 调用gRPC服务方法
// 	return c.client.SomeMethod(ctx, req)
// }

// // Order服务客户端
// type OrderClient struct {
// //...
// }

// func NewOrderClient() *OrderClient {
// //...
// }

// // 初始化客户端
// var (
// 	userClient = NewUserClient()
// 	orderClient = NewOrderClient()
// )

// // 调用
// // resp, err := userClient.Call(ctx, req)
