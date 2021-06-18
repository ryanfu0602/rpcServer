package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"rpcServer/controllers"
	"rpcServer/model"
)

type Article struct {
}

type GetByIdRequest struct {
	Id int
}

type GetByIdResponse struct {
	Ariticles *model.Article
}

func (*Article) GetById(req GetByIdRequest, res *GetByIdResponse) error {
	res.Ariticles = controllers.GetArticleById(req.Id)
	fmt.Println(res.Ariticles)
	return nil
}

func main() {
	rpc.Register(new(Article))
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection")

	http.Serve(lis, nil)
}
