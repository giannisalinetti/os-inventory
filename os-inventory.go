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
	GeneratorDeploymentType string `yaml: "generatorDeploymentType"`
	GeneratorSshUser        string `yaml: "generatorSshUser"`
	GeneratorNfsEnabled     bool   `yaml: "generatorNfsEnabled"`
}

// New loads default values and returns *Inventory
func New(defaultValues map[string]interface{}) *Inventory {
	inv := &Inventory{
		GeneratorDeploymentType: defaultValues["generatorDeploymentType"].(string),
		GeneratorSshUser:        defaultValues["generatorSshUser"].(string),
		GeneratorNfsEnabled:     defaultValues["generatorNfsEnabled"].(bool),
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
