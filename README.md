# Atlas Template
This project is intended to enhance
[Atlas CLI](https://github.com/infobloxopen/atlas-cli)
At some point this type of functionality might be integrated into
Atlas CLI to auto generate more of the application templates.

The use case I had was that for the
[CMDB Project](https://github.com/seizadi/cmdb) which based on
Atlas Toolkit, I manually code 5 services by hand, it was not a
very DRY process and I was doing a lot repetitve functions. I
created this project to complete the remaining 9 services and
see if the process could be automated. This setups up the basic
scafolding so that the developer could modify the proto file to
custom function, but the basic service, CRUD and migrations would
be automated.

In this release you define your resources using ./config/model.yaml
```yaml
resources:
  Application:
  Artifact:
  AwsRdsInstance:
  AwsService:
  Container:
  Deployment:
  Environment:
  KubeCluster:
  Manifest:
  Region:
  Secret:
  Vault:
  VersionTag:
```
In subsequent releases we could look at a more complex golang model
with associations and drive the generation from the model. See the
model definitions in ./model/* and the following output of go-erd
here [model ERD](https://github.com/seizadi/atlas-template/blob/master/doc/db/out.pdf)

## Build
```sh
make build
```

## Run
```
./bin/atlas-template create
```

You will have artifacts created in the ./resolved directory, I copied
them to the CMDB project to corresponding folders there. I had to merge
my migration and proto files since I started to manually update them
but this wont be a problem for a new project. It is recommended you
reorganize your migration files ordered from the leaf nodes up to
the root based your ERD diagram. So that you will alter your tables
the least.
```sh
$ find resolved -type f -print
resolved/cmd/server/endpoints.go
resolved/db/migration/00002_artifacts.down.sql
resolved/db/migration/00002_artifacts.up.sql
resolved/db/migration/00002_manifests.down.sql
resolved/db/migration/00002_manifests.up.sql
resolved/db/migration/00003_artifacts.down.sql
resolved/db/migration/00003_artifacts.up.sql
resolved/db/migration/00003_aws_rds_instances.down.sql
resolved/db/migration/00003_aws_rds_instances.up.sql
resolved/db/migration/00004_aws_rds_instances.down.sql
resolved/db/migration/00004_aws_rds_instances.up.sql
resolved/db/migration/00004_deployments.down.sql
resolved/db/migration/00004_deployments.up.sql
resolved/db/migration/00005_containers.down.sql
resolved/db/migration/00005_containers.up.sql
resolved/db/migration/00005_manifests.down.sql
resolved/db/migration/00005_manifests.up.sql
resolved/db/migration/00006_environments.down.sql
resolved/db/migration/00006_environments.up.sql
resolved/db/migration/00006_secrets.down.sql
resolved/db/migration/00006_secrets.up.sql
resolved/db/migration/00007_kube_clusters.down.sql
resolved/db/migration/00007_kube_clusters.up.sql
resolved/db/migration/00007_vaults.down.sql
resolved/db/migration/00007_vaults.up.sql
resolved/db/migration/00008_applications.down.sql
resolved/db/migration/00008_applications.up.sql
resolved/db/migration/00008_vaults.down.sql
resolved/db/migration/00008_vaults.up.sql
resolved/db/migration/00009_aws_services.down.sql
resolved/db/migration/00009_aws_services.up.sql
resolved/db/migration/00009_version_tags.down.sql
resolved/db/migration/00009_version_tags.up.sql
resolved/db/migration/00010_applications.down.sql
resolved/db/migration/00010_applications.up.sql
resolved/db/migration/00010_containers.down.sql
resolved/db/migration/00010_containers.up.sql
resolved/db/migration/00011_aws_services.down.sql
resolved/db/migration/00011_aws_services.up.sql
resolved/db/migration/00011_environments.down.sql
resolved/db/migration/00011_environments.up.sql
resolved/db/migration/00012_deployments.down.sql
resolved/db/migration/00012_deployments.up.sql
resolved/db/migration/00012_kube_clusters.down.sql
resolved/db/migration/00012_kube_clusters.up.sql
resolved/db/migration/00013_regions.down.sql
resolved/db/migration/00013_regions.up.sql
resolved/db/migration/00014_secrets.down.sql
resolved/db/migration/00014_secrets.up.sql
resolved/db/migration/00014_version_tags.down.sql
resolved/db/migration/00014_version_tags.up.sql
resolved/pkg/pb/cmdb.proto
resolved/pkg/svc/servers.go
```
