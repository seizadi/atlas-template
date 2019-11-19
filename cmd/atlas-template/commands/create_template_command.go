package commands

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"text/template"
	
	"github.com/codegangsta/cli"
	"github.com/iancoleman/strcase"
)

type confModel struct {
	Resources map[string]string `yaml:"resources"`
}

type templateResource struct {
	NameCamel string
	NameCamels string
	NameLowerCamel string
	NameLowerCamels string
	NameSnake string
	NameSnakes string
	MigrateVer string
	Path string
}


func createTemplate(c *cli.Context) {
	
	// TODO - Read Config from ./model/*.go files
	// Read config
	var config confModel
	config.getConf()
	
	r := []templateResource{}
	migrateNum := 2
	migrateStr := ""
	
	for k, _ := range config.Resources {
		if (migrateNum > 9) {
			migrateStr = "000" + strconv.Itoa(migrateNum)
		} else {
			migrateStr = "0000" + strconv.Itoa(migrateNum)
		}
		resource := templateResource {
			NameCamel: strcase.ToCamel(k),
			NameCamels: strPlural(strcase.ToCamel(k)),
			NameLowerCamel: strcase.ToLowerCamel(k),
			NameLowerCamels: strPlural(strcase.ToLowerCamel(k)),
			NameSnake: strcase.ToSnake(k),
			NameSnakes: strPlural(strcase.ToSnake(k)),
			MigrateVer: migrateStr,
		}
		r = append(r, resource)
		migrateNum += 1
	}
	
	err := runTemplate(r,
		"template/pkg/pb/template.proto.template",
		"resolved/pkg/pb/secops.proto" )
	
	if err != nil {
		log.Fatalf("failed to create pkg/pb/secops.proto\n%s\n", err)
	}
	
	err = runTemplate(r,
		"template/pkg/svc/servers.go.template",
		"resolved/pkg/svc/servers.go" )
	
	if err != nil {
		log.Fatalf("failed to create pkg/svc/servers.go\n%s\n", err)
	}
	
	err = runTemplate(r,
		"template/cmd/server/servers.go.template",
		"resolved/cmd/server/servers.go" )
	
	if err != nil {
		log.Fatalf("failed to create cmd/server/servers.go\n%s\n", err)
	}
	
	err = runTemplate(r,
		"template/cmd/server/endpoints.go.template",
		"resolved/cmd/server/endpoints.go" )
	
	if err != nil {
		log.Fatalf("failed to create cmd/server/endpoints.go\n%s\n", err)
	}
	
	for _ , res := range r {
		err = runTemplate([]templateResource{res},
			"template/db/migration/down.sql.template",
			"resolved/db/migration/" + res.MigrateVer + "_" + res.NameSnakes +  ".down.sql" )
		
		if err != nil {
			log.Fatalf("failed to create resolved/db/migration/" + res.MigrateVer + "_" + res.NameSnakes +  ".down.sql\n%s\n", err)
		}
		
		err = runTemplate([]templateResource{res},
			"template/db/migration/up.sql.template",
			"resolved/db/migration/" + res.MigrateVer + "_" + res.NameSnakes +  ".up.sql" )
		
		if err != nil {
			log.Fatalf("failed to create resolved/db/migration/" + res.MigrateVer + "_" + res.NameSnakes +  ".up.sql\n%s\n", err)
		}
	}
		
	return
}

func strPlural(s string) string {
	// Had some use cases like AwsRds or Kubernetes and tried to
	// fix the plural since plural is not right like: Kubernetess
	// This is fine for URLs or Table names but breaks proto since
	// message and service definitions have collision:
	// sc-l-seizadi:cmdb seizadi$ make protobuf
	// github.com/seizadi/cmdb/pkg/pb/cmdb.proto:144:9: "AwsRds" is already defined in "api.cmdb".
	// github.com/seizadi/cmdb/pkg/pb/cmdb.proto:1087:9: "Kubernetes" is already defined in "api.cmdb".
	// So we need to have people pick names that have natural plural so in the above case
	// AwsRds -> AwsRdsInstance or Kubernetes => KubeCluster
	//if (s[len(s)-1:] == "s") {
	//	return s
	//}
	return s + "s"
}

func runTemplate(r []templateResource, src string, dst string ) error {
	// Create a new template and parse the file into it
	name := path.Base(src)
	t, err := template.New(name).ParseFiles(src)
	if err != nil {
		log.Fatalf("parsing template: %s\n", err)
	}
	// Create Template
	f, err := os.Create(dst)
	if err != nil {
		log.Fatalf("create file %s failed: %s\n", dst, err)
	}
	err = t.Execute(f, r)
	if err != nil {
		log.Fatalf("failed executing template: %s\n", err)
	}
	
	return nil
}

func (c *confModel) getConf() *confModel {
	
	yamlFile, err := ioutil.ReadFile("config/model.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v\n ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v\n", err)
	}
	
	return c
}
