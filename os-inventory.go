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
func parseYAML(yamlFile string, inv *Inventory) error {
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

// doSanityChecks verifies passed parameters
func doSanityChecks(inv *Inventory) error {

	err := inv.CheckDeploymentType()
	if err != nil {
		return err
	}
	err = inv.CheckInstallVersion()
	if err != nil {
		return err
	}
	err = inv.CheckClusterMethod()
	if err != nil {
		return err
	}
	err = inv.CheckInfraIpv4()
	if err != nil {
		return err
	}
	err = inv.CheckSdnPlugin()
	if err != nil {
		return err
	}

	return nil
}

func main() {

	showDefaults := flag.BoolP("show-defaults", "d", false, "Dump defaults parameters to stdout.")
	loadYAML := flag.StringP("load-yaml", "f", "", "Load configuration from YAML file.")
	dumpFile := flag.StringP("output", "o", "", "Printe generated inventory to file.")

	flag.Parse()

	// Load default values
	inventory := New(defaults)

	// Print defaults to stdout in YAML format
	if *showDefaults {
		d, err := yaml.Marshal(&inventory)
		if err != nil {
			log.Fatal("error: %v", err)
		}
		fmt.Printf("---\n%s", d)
		return
	}

	// Load YAML configuration if passed
	if *loadYAML != "" {
		filePath := *loadYAML
		fmt.Printf("Yaml loaded\n")
		err := parseYAML(filePath, inventory)
		if err != nil {
			log.Fatal("Error opening YAML: %v", err)
			return
		}
	}

	// Create new template and parse content
	t, err := template.New("OpenShiftInventory").Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// Run sanity checks before exporting
	err = doSanityChecks(inventory)
	if err != nil {
		log.Fatal("Sanity check: ", err)
		return
	}

	// Generate the processed inventory
	if *dumpFile != "" {
		f, err := os.Create(*dumpFile)
		if err != nil {
			log.Fatal("Create file: ", err)
			return
		}
		// Print inventory to file
		err = t.Execute(f, inventory)
		if err != nil {
			log.Fatal("Execute: ", err)
			return
		}
	} else {
		// Print inventory to stdout
		err = t.Execute(os.Stdout, inventory)
		if err != nil {
			log.Fatal("Execute: ", err)
			return
		}
	}
}
