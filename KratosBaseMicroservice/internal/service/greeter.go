package service

import (
	"KratosBaseMicroservice/internal/myfunction"
	"context"

	hellov1 "KratosBaseMicroservice/api/helloworld/v1"
	"KratosBaseMicroservice/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	hellov1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *hellov1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	var a float64 = 178.012
	b := 8956.124
	var HelloMassage *string
	HelloMassage, _ = myfunction.DataTrafficSizeCalculator(&a, &b)

	return &hellov1.HelloReply{Message: *HelloMassage + " " + g.Hello}, nil
}
