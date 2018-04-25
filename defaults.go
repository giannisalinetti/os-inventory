package main

var defaults = make(map[string]interface{})

func init() {
	defaults["generatorDeploymentType"] = "origin"
	defaults["generatorSshUser"] = "root"
	defaults["generatorNfsEnabled"] = true
	defaults["generatorRegistryNativeNfs"] = true
	defaults["generatorHaproxyEnabled"] = false
	defaults["generatorInstallVersion"] = "v3.9"
	defaults["generatorSkipChecks"] = false
	defaults["generatorMultiMaster"] = false
	defaults["generatorClusterMethod"] = "native"
	defaults["generatorClusterHostname"] = "osapi.example.com"
	defaults["generatorClusterPublicHostname"] = "osapi.example.com"
	defaults["generatorContainerizedDeploy"] = false
	defaults["generatorContainerizedOvs"] = false
	defaults["generatorContainerizedNode"] = false
	defaults["generatorContainerizedMaster"] = false
	defaults["generatorContainerizedEtcd"] = false
	defaults["generatorSystemImagesRegistry"] = "registry.access.redhat.com"
	defaults["generatorOpenshiftUseCrio"] = false
	defaults["generatorOpenshiftCrioUseRpm"] = false
	defaults["generatorMultiInfra"] = false
	defaults["generatorUseXip"] = false
	defaults["generatorInfraIpv4"] = ""
	defaults["generatorExtDnsWildcard"] = "osapps.example.com"
	defaults["generatorSdnPlugin"] = "ovs-multitenant"
	defaults["generatorDisableServiceCatalog"] = true
	defaults["generatorInfraReplicas"] = 3 // TODO: Calculate the value based on the infra nodes provided
	defaults["generatorMetricsEnabled"] = true
	defaults["generatorDeployHosa"] = true
	defaults["generatorMetricsNativeNfs"] = true
	defaults["generatorPrometheusEnabled"] = true
	defaults["generatorPrometheusNativeNfs"] = true
	defaults["generatorLoggingEnabled"] = true
	defaults["generatorLoggingNativeNfs"] = true
	defaults["generatorMastersList"] = make([]string, 0)
	defaults["generatorEtcdList"] = make([]string, 0)
	defaults["generatorLbList"] = make([]string, 0)
	defaults["generatorNodesMap"] = make(map[string]string)
}
