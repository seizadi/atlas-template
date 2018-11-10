package main

import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/spf13/viper"
	
	"github.com/seizadi/cmdb/pkg/pb"
)

func RegisterGatewayEndpoints() gateway.Option {
	return gateway.WithEndpointRegistration(viper.GetString("server.version"),
		pb.RegisterCmdbHandlerFromEndpoint,
		pb.RegisterManifestsHandlerFromEndpoint,
		pb.RegisterArtifactsHandlerFromEndpoint,
		pb.RegisterAwsRdsInstancesHandlerFromEndpoint,
		pb.RegisterContainersHandlerFromEndpoint,
		pb.RegisterEnvironmentsHandlerFromEndpoint,
		pb.RegisterKubeClustersHandlerFromEndpoint,
		pb.RegisterVaultsHandlerFromEndpoint,
		pb.RegisterVersionTagsHandlerFromEndpoint,
		pb.RegisterApplicationsHandlerFromEndpoint,
		pb.RegisterAwsServicesHandlerFromEndpoint,
		pb.RegisterDeploymentsHandlerFromEndpoint,
		pb.RegisterRegionsHandlerFromEndpoint,
		pb.RegisterSecretsHandlerFromEndpoint,
	)
}
