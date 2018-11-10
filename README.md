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
here [model ERD](https://github.com/seizadi/app-template/blob/master/doc/db/out.pdf)
