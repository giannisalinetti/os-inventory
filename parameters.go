package main

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
	ip := net.ParseIP(i.GeneratorInfraIpv4)
	if ip == nil {
		return errors.New("Invalid IPv4 address.")
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
