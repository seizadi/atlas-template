package svc

import (
	"github.com/seizadi/secops/pkg/pb"
)

type regionsServer struct {
	*pb.RegionsDefaultServer
}
// NewRegionsServer returns an instance of the default Region server interface
func NewRegionsServer() (pb.RegionsServer, error) {
	return &regionsServer{&pb.RegionsDefaultServer{}}, nil
}

type ec2sServer struct {
	*pb.Ec2sDefaultServer
}
// NewEc2sServer returns an instance of the default Ec2 server interface
func NewEc2sServer() (pb.Ec2sServer, error) {
	return &ec2sServer{&pb.Ec2sDefaultServer{}}, nil
}

type applicationsServer struct {
	*pb.ApplicationsDefaultServer
}
// NewApplicationsServer returns an instance of the default Application server interface
func NewApplicationsServer() (pb.ApplicationsServer, error) {
	return &applicationsServer{&pb.ApplicationsDefaultServer{}}, nil
}

type scansServer struct {
	*pb.ScansDefaultServer
}
// NewScansServer returns an instance of the default Scan server interface
func NewScansServer() (pb.ScansServer, error) {
	return &scansServer{&pb.ScansDefaultServer{}}, nil
}

type knowledgeBasesServer struct {
	*pb.KnowledgeBasesDefaultServer
}
// NewKnowledgeBasesServer returns an instance of the default KnowledgeBase server interface
func NewKnowledgeBasesServer() (pb.KnowledgeBasesServer, error) {
	return &knowledgeBasesServer{&pb.KnowledgeBasesDefaultServer{}}, nil
}

type vulnerabilitysServer struct {
	*pb.VulnerabilitysDefaultServer
}
// NewVulnerabilitysServer returns an instance of the default Vulnerability server interface
func NewVulnerabilitysServer() (pb.VulnerabilitysServer, error) {
	return &vulnerabilitysServer{&pb.VulnerabilitysDefaultServer{}}, nil
}

type cloudProvidersServer struct {
	*pb.CloudProvidersDefaultServer
}
// NewCloudProvidersServer returns an instance of the default CloudProvider server interface
func NewCloudProvidersServer() (pb.CloudProvidersServer, error) {
	return &cloudProvidersServer{&pb.CloudProvidersDefaultServer{}}, nil
}

type amisServer struct {
	*pb.AmisDefaultServer
}
// NewAmisServer returns an instance of the default Ami server interface
func NewAmisServer() (pb.AmisServer, error) {
	return &amisServer{&pb.AmisDefaultServer{}}, nil
}

type vpcsServer struct {
	*pb.VpcsDefaultServer
}
// NewVpcsServer returns an instance of the default Vpc server interface
func NewVpcsServer() (pb.VpcsServer, error) {
	return &vpcsServer{&pb.VpcsDefaultServer{}}, nil
}

type registerysServer struct {
	*pb.RegisterysDefaultServer
}
// NewRegisterysServer returns an instance of the default Registery server interface
func NewRegisterysServer() (pb.RegisterysServer, error) {
	return &registerysServer{&pb.RegisterysDefaultServer{}}, nil
}

type containersServer struct {
	*pb.ContainersDefaultServer
}
// NewContainersServer returns an instance of the default Container server interface
func NewContainersServer() (pb.ContainersServer, error) {
	return &containersServer{&pb.ContainersDefaultServer{}}, nil
}

