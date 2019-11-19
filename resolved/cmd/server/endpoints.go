package main

import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/spf13/viper"
	
	"github.com/seizadi/secops/pkg/pb"
)

func RegisterGatewayEndpoints() gateway.Option {
	return gateway.WithEndpointRegistration(viper.GetString("server.version"),
		pb.RegisterSecopsHandlerFromEndpoint,
		pb.RegisterRegionsHandlerFromEndpoint,
		pb.RegisterEc2sHandlerFromEndpoint,
		pb.RegisterApplicationsHandlerFromEndpoint,
		pb.RegisterScansHandlerFromEndpoint,
		pb.RegisterKnowledgeBasesHandlerFromEndpoint,
		pb.RegisterVulnerabilitysHandlerFromEndpoint,
		pb.RegisterCloudProvidersHandlerFromEndpoint,
		pb.RegisterAmisHandlerFromEndpoint,
		pb.RegisterVpcsHandlerFromEndpoint,
		pb.RegisterRegisterysHandlerFromEndpoint,
		pb.RegisterContainersHandlerFromEndpoint,
	)
}
