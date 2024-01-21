package main

import (
	"howell/howell_rpc/infra/db"
	howell_rpc "howell/howell_rpc/kitex_gen/coder/hao/howell_rpc/howellrpcservice"
	"log"
)

func init() {
	db.Init()
}

func main() {
	svr := howell_rpc.NewServer(new(HowellRpcServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
