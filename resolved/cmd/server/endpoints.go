package main

import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/spf13/viper"
	
	"github.com/seizadi/cmdb/pkg/pb"
)

func RegisterGatewayEndpoints() gateway.Option {
	return gateway.WithEndpointRegistration(viper.GetString("server.version"),
		pb.RegisterCmdbHandlerFromEndpoint,
		pb.RegisterArtifactsHandlerFromEndpoint,
		pb.RegisterVaultsHandlerFromEndpoint,
		pb.RegisterVersionTagsHandlerFromEndpoint,
		pb.RegisterDeploymentsHandlerFromEndpoint,
		pb.RegisterEnvironmentsHandlerFromEndpoint,
		pb.RegisterKubeClustersHandlerFromEndpoint,
		pb.RegisterManifestsHandlerFromEndpoint,
		pb.RegisterApplicationsHandlerFromEndpoint,
		pb.RegisterAwsRdsInstancesHandlerFromEndpoint,
		pb.RegisterAwsServicesHandlerFromEndpoint,
		pb.RegisterContainersHandlerFromEndpoint,
		pb.RegisterRegionsHandlerFromEndpoint,
		pb.RegisterSecretsHandlerFromEndpoint,
	)
}
