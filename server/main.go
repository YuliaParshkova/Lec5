package main

import (
  p "github.com/YuliaParshkova/Lec5/server/proto/consignment"
)

const (
	port = ":50051"
)


type repository interface {
	Create (p.Command)
}
func main(){

}