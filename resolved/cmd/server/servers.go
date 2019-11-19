package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jinzhu/gorm"
	"github.com/seizadi/secops/pkg/pb"
	"github.com/seizadi/secops/pkg/svc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func CreateServer(logger *logrus.Logger, db *gorm.DB, interceptors []grpc.UnaryServerInterceptor) (*grpc.Server, error) {
	// create new gRPC grpcServer with middleware chain
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)))
	
	// register all of our services into the grpcServer
	s, err := svc.NewBasicServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterSecopsServer(grpcServer, s)

	region, err := svc.NewRegionsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterRegionsServer(grpcServer, region)

	ec_2, err := svc.NewEc2sServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterEc2sServer(grpcServer, ec_2)

	application, err := svc.NewApplicationsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterApplicationsServer(grpcServer, application)

	scan, err := svc.NewScansServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterScansServer(grpcServer, scan)

	knowledge_base, err := svc.NewKnowledgeBasesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterKnowledgeBasesServer(grpcServer, knowledge_base)

	vulnerability, err := svc.NewVulnerabilitysServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterVulnerabilitysServer(grpcServer, vulnerability)

	cloud_provider, err := svc.NewCloudProvidersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterCloudProvidersServer(grpcServer, cloud_provider)

	ami, err := svc.NewAmisServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAmisServer(grpcServer, ami)

	vpc, err := svc.NewVpcsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterVpcsServer(grpcServer, vpc)

	registery, err := svc.NewRegisterysServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterRegisterysServer(grpcServer, registery)

	container, err := svc.NewContainersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterContainersServer(grpcServer, container)

	return grpcServer, nil
}
