package defaults

var DefaultCfg = make(map[string]interface{})

func init() {
	DefaultCfg["generatorDeploymentType"] = "origin"
	DefaultCfg["generatorSshUser"] = "root"
	DefaultCfg["generatorNfsEnabled"] = true
	DefaultCfg["generatorGlusterfsEnabled"] = false
	DefaultCfg["generatorGlusterfsRegEnabled"] = false
	DefaultCfg["generatorRegistryNativeNfs"] = true
	DefaultCfg["generatorRegistryCNS"] = false
	DefaultCfg["generatorHaproxyEnabled"] = false
	DefaultCfg["generatorInstallVersion"] = "v3.9"
	DefaultCfg["generatorSkipChecks"] = false
	DefaultCfg["generatorMultiMaster"] = false
	DefaultCfg["generatorClusterMethod"] = "native"
	DefaultCfg["generatorClusterHostname"] = "osapi.example.com"
	DefaultCfg["generatorClusterPublicHostname"] = "osapi.example.com"
	DefaultCfg["generatorContainerizedDeploy"] = false
	DefaultCfg["generatorContainerizedOvs"] = false
	DefaultCfg["generatorContainerizedNode"] = false
	DefaultCfg["generatorContainerizedMaster"] = false
	DefaultCfg["generatorContainerizedEtcd"] = false
	DefaultCfg["generatorSystemImagesRegistry"] = "registry.access.redhat.com"
	DefaultCfg["generatorOpenshiftUseCrio"] = false
	DefaultCfg["generatorOpenshiftCrioUseRpm"] = false
	DefaultCfg["generatorMultiInfra"] = false
	DefaultCfg["generatorUseXip"] = false
	DefaultCfg["generatorInfraIpv4"] = ""
	DefaultCfg["generatorExtDnsWildcard"] = "osapps.example.com"
	DefaultCfg["generatorSdnPlugin"] = "ovs-multitenant"
	DefaultCfg["generatorDisableServiceCatalog"] = true
	DefaultCfg["generatorInfraReplicas"] = 3 // TODO: Calculate the value based on the infra nodes provided
	DefaultCfg["generatorMetricsEnabled"] = true
	DefaultCfg["generatorDeployHosa"] = true
	DefaultCfg["generatorMetricsNativeNfs"] = true
	DefaultCfg["generatorPrometheusEnabled"] = false
	DefaultCfg["generatorPrometheusNativeNfs"] = true
	DefaultCfg["generatorLoggingEnabled"] = true
	DefaultCfg["generatorLoggingNativeNfs"] = true
	DefaultCfg["generatorMastersList"] = make([]string, 0)
	DefaultCfg["generatorEtcdList"] = make([]string, 0)
	DefaultCfg["generatorLbList"] = make([]string, 0)
	DefaultCfg["generatorNfsList"] = make([]string, 0)
	DefaultCfg["generatorGlusterfsMap"] = make(map[string]string)
	DefaultCfg["generatorGlusterfsRegMap"] = make(map[string]string)
	DefaultCfg["generatorNodesMap"] = make(map[string]string)
}
