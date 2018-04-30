package parameters

import (
	"errors"
	"net"
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
	GeneratorClusterMethod         string            `yaml:"cluster_method"`
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

// Check if deployment type is enterprise or origin
func (i *Inventory) CheckDeploymentType() error {
	if i.GeneratorDeploymentType == "origin" || i.GeneratorDeploymentType == "enterprise" {
		return nil
	} else {
		return errors.New("Invalid deployemnt type.")
	}
}

// Check if installation version is valid
func (i *Inventory) CheckInstallVersion() error {
	versions := []string{"v3.4", "v3.5", "v3.6", "v3.7", "v3.9", "v3.10", "v3.11"}
	for _, v := range versions {
		if i.GeneratorInstallVersion == v {
			return nil
		}
	}
	return errors.New("Invalid or unsupported version.")
}

// Check if cluster method is native
func (i *Inventory) CheckClusterMethod() error {
	if i.GeneratorClusterMethod != "native" {
		return errors.New("Invalid cluster method.")
	}
	return nil
}

// Check if string is a valid IPv4 address
func (i *Inventory) CheckInfraIpv4() error {
	if i.GeneratorInfraIpv4 != "" {
		ip := net.ParseIP(i.GeneratorInfraIpv4)
		if ip == nil {
			return errors.New("Invalid IPv4 address.")
		}
	}
	return nil
}

// Check if SND plugin is among the supported ones
func (i *Inventory) CheckSdnPlugin() error {
	plugins := []string{"ovs-subnet", "ovs-multitenant", "ovs-networkpolicy"}
	for _, v := range plugins {
		if i.GeneratorSdnPlugin == v {
			return nil
		}
	}
	return errors.New("Invalid SND plugin.")
}
