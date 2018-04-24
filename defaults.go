package main

var defaults = make(map[string]interface{})

func init() {
	defaults["generatorDeploymentType"] = "origin"
	defaults["generatorSshUser"] = "root"
	defaults["generatorNfsEnabled"] = true
}
