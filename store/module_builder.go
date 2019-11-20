package store
import (
	grpc "google.golang.org/grpc"
)

type ModuleBuilder interface {
	CreateTable(name string)
	RegisterServer(desc *grpc.ServiceDesc, server interface{})
}

func NewAutoUint64TableBuilder(builder ModuleBuilder) AutoUInt64TableBuilder  {
	panic("TODO")
}
