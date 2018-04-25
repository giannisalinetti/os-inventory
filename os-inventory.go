package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Inventory is the basic definition of all parameters
type Inventory struct {
	GeneratorDeploymentType        string            `yaml:"deployment_type"`
	GeneratorSshUser               string            `yaml:"ssh_user"`
	GeneratorNfsEnabled            bool              `yaml:"nfs_enabled"`
	GeneratorRegistryNativeNfs     bool              `yaml:"registry_native_nfs"`
	GeneratorHaproxyEnabled        bool              `yaml:"haproxy_enabled"`
	GeneratorInstallVersion        string            `yaml:"install_version"`
	GeneratorSkipChecks            bool              `yaml:"skip_checks"`
	GeneratorMultiMaster           bool              `yaml:"multi_master"`
	GeneratorClusterMethod         string            `yaml:"cluste_rmethod"`
	GeneratorClusterHostname       string            `yaml:"cluster_hostname"`
	GeneratorClusterPublicHostname string            `yaml:"cluster_public_hostname"`
	GeneratorContainerizedDeploy   bool              `yaml:"containerized_deploy"`
	GeneratorContainerizedOvs      bool              `yaml:"containerized_ovs"`
	GeneratorContainerizedNode     bool              `yaml:"containerized_node"`
	GeneratorContainerizedMaster   bool              `yaml:"containerized_master"`
	GeneratorContainerizedEtcd     bool              `yaml:"containerized_etcd"`
	GeneratorSystemImagesRegistry  string            `yaml:"system_images_registry"`
	GeneratorOpenshiftUseCrio      bool              `yaml:"openshift_use_crio"`
	GeneratorOpenshiftCrioUseRpm   bool              `yaml:"openshift_crio_use_rpm"`
	GeneratorMultiInfra            bool              `yaml:"multi_infra"`
	GeneratorUseXip                bool              `yaml:"use_xip"`
	GeneratorInfraIpv4             string            `yaml:"infra_ipv4"`
	GeneratorExtDnsWildcard        string            `yaml:"ext_dns_wildcard"`
	GeneratorSdnPlugin             string            `yaml:"sdn_plugin"`
	GeneratorDisableServiceCatalog bool              `yaml:"disable_servicecatalog"`
	GeneratorInfraReplicas         int               `yaml:"infra_replicas"`
	GeneratorMetricsEnabled        bool              `yaml:"metrics_enabled"`
	GeneratorDeployHosa            bool              `yaml:"deploy_hosa"`
	GeneratorMetricsNativeNfs      bool              `yaml:"metrics_native_nfs"`
	GeneratorPrometheusEnabled     bool              `yaml:"prometheus_enabled"`
	GeneratorPrometheusNativeNfs   bool              `yaml:"prometheus_native_nfs"`
	GeneratorLoggingEnabled        bool              `yaml:"logging_enabled"`
	GeneratorLoggingNativeNfs      bool              `yaml:"logging_native_nfs"`
	GeneratorMastersList           []string          `yaml:"masters_list"`
	GeneratorEtcdList              []string          `yaml:"etcd_list"`
	GeneratorLbList                []string          `yaml:"lb_list"`
	GeneratorNodesMap              map[string]string `yaml:"nodes_map"`
}

// New loads default values and returns *Inventory
func New(defaultValues map[string]interface{}) *Inventory {
	inv := &Inventory{
		GeneratorDeploymentType:        defaultValues["generatorDeploymentType"].(string),
		GeneratorSshUser:               defaultValues["generatorSshUser"].(string),
		GeneratorNfsEnabled:            defaultValues["generatorNfsEnabled"].(bool),
		GeneratorRegistryNativeNfs:     defaultValues["generatorRegistryNativeNfs"].(bool),
		GeneratorHaproxyEnabled:        defaultValues["generatorHaproxyEnabled"].(bool),
		GeneratorInstallVersion:        defaultValues["generatorInstallVersion"].(string),
		GeneratorSkipChecks:            defaultValues["generatorSkipChecks"].(bool),
		GeneratorMultiMaster:           defaultValues["generatorMultiMaster"].(bool),
		GeneratorClusterMethod:         defaultValues["generatorClusterMethod"].(string),
		GeneratorClusterHostname:       defaultValues["generatorClusterHostname"].(string),
		GeneratorClusterPublicHostname: defaultValues["generatorClusterPublicHostname"].(string),
		GeneratorContainerizedDeploy:   defaultValues["generatorContainerizedDeploy"].(bool),
		GeneratorContainerizedOvs:      defaultValues["generatorContainerizedOvs"].(bool),
		GeneratorContainerizedNode:     defaultValues["generatorContainerizedNode"].(bool),
		GeneratorContainerizedMaster:   defaultValues["generatorContainerizedMaster"].(bool),
		GeneratorContainerizedEtcd:     defaultValues["generatorContainerizedEtcd"].(bool),
		GeneratorSystemImagesRegistry:  defaultValues["generatorSystemImagesRegistry"].(string),
		GeneratorOpenshiftUseCrio:      defaultValues["generatorOpenshiftUseCrio"].(bool),
		GeneratorOpenshiftCrioUseRpm:   defaultValues["generatorOpenshiftCrioUseRpm"].(bool),
		GeneratorMultiInfra:            defaultValues["generatorMultiInfra"].(bool),
		GeneratorUseXip:                defaultValues["generatorUseXip"].(bool),
		GeneratorInfraIpv4:             defaultValues["generatorInfraIpv4"].(string),
		GeneratorExtDnsWildcard:        defaultValues["generatorExtDnsWildcard"].(string),
		GeneratorSdnPlugin:             defaultValues["generatorSdnPlugin"].(string),
		GeneratorDisableServiceCatalog: defaultValues["generatorDisableServiceCatalog"].(bool),
		GeneratorInfraReplicas:         defaultValues["generatorInfraReplicas"].(int),
		GeneratorMetricsEnabled:        defaultValues["generatorMetricsEnabled"].(bool),
		GeneratorDeployHosa:            defaultValues["generatorDeployHosa"].(bool),
		GeneratorMetricsNativeNfs:      defaultValues["generatorMetricsNativeNfs"].(bool),
		GeneratorPrometheusEnabled:     defaultValues["generatorPrometheusEnabled"].(bool),
		GeneratorPrometheusNativeNfs:   defaultValues["generatorPrometheusNativeNfs"].(bool),
		GeneratorLoggingEnabled:        defaultValues["generatorLoggingEnabled"].(bool),
		GeneratorLoggingNativeNfs:      defaultValues["generatorLoggingNativeNfs"].(bool),
		GeneratorMastersList:           defaultValues["generatorMastersList"].([]string),
		GeneratorEtcdList:              defaultValues["generatorEtcdList"].([]string),
		GeneratorLbList:                defaultValues["generatorLbList"].([]string),
		GeneratorNodesMap:              defaultValues["generatorNodesMap"].(map[string]string),
	}
	return inv
}

// ParseYAML parses a YAML input file for custom paramters values
func ParseYAML(yamlFile string, inv *Inventory) error {
	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &inv)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	showDefaults := flag.Bool("show-defaults", false, "Dump defaults to desired format.")
	loadYAML := flag.String("load-yaml", "", "Load configuration from YAML file.")

	flag.Parse()

	// Load default values
	inventory := New(defaults)

	if *showDefaults {
		d, err := yaml.Marshal(&inventory)
		if err != nil {
			log.Fatal("error: %v", err)
		}
		fmt.Printf("---\n%s", d)
		return
	}

	if *loadYAML != "" {
		filePath := *loadYAML
		fmt.Printf("Yaml loaded\n")
		err := ParseYAML(filePath, inventory)
		if err != nil {
			log.Fatal("Error opening YAML: %v", err)
			return
		}
	}

	t := template.New("OpenShiftInventory")
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse:", err)
		return
	}

	err = t.Execute(os.Stdout, inventory)
	if err != nil {
		log.Fatal("Execute:", err)
		return
	}
}
