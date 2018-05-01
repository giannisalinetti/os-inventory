package main

const tmpl = `# OpenShift {{ .GeneratorDeploymentType }} advanced installation
[OSEv3:children]
masters
etcd
nodes
{{ if .GeneratorNfsEnabled }}nfs{{"\n"}}{{ else }}{{ end -}}
{{ if .GeneratorHaproxyEnabled }}lb{{"\n"}}{{ else }}{{"\n"}}{{ end -}}
{{ if .GeneratorGlusterfsEnabled }}glusterfs{{"\n"}}{{ else }}{{ end -}}
{{ if .GeneratorGlusterfsRegEnabled }}glusterfs_registry{{"\n"}}{{ else }}{{ end -}}

{{"\n"}}[OSEv3:vars]
ansible_ssh_user={{ .GeneratorSshUser }}
deployment_type={{ .GeneratorDeploymentType }}
openshift_release={{ .GeneratorInstallVersion -}}

{{ if .GeneratorSkipChecks }}
{{"\n"}}# Disable package, disk and memory checks
openshift_disable_check=memory_availability,disk_availability,docker_storage,docker_storage_driver,docker_image_availability,package_version,package_availability,package_update
{{- end -}}

{{- if .GeneratorMultiMaster }}
{{"\n"}}# Configure cluster
openshift_master_cluster_method={{ .GeneratorClusterMethod }}
openshift_master_cluster_hostname={{ .GeneratorClusterHostname }}
openshift_master_cluster_public_hostname={{ .GeneratorClusterPublicHostname }}
{{- end }}

{{- if .GeneratorContainerizedDeploy }}
{{"\n"}}# Deploy Containerized components
{{ if .GeneratorContainerizedOvs }}
openshift_use_openvswitch_system_container=True
{{- end }}
{{ if .GeneratorContainerizedNode }}
openshift_use_node_system_container=True
{{- end }}
{{ if .GeneratorContainerizedMaster }}
openshift_use_master_system_container=True
{{- end }}
{{ if .GeneratorContainerizedEtcd }}
openshift_use_etcd_system_container=True
{{- end }}
system_images_registry={{ .GeneratorSystemImagesRegistry }}
{{- end }}

{{- if and (.GeneratorOpenshiftUseCrio) ( eq .GeneratorInstallVersion "v3.9") }}
{{"\n"}}# CRI-O configuration
openshift_use_crio=True
{{ if .GeneratorOpenshiftCrioUseRpm }}
openshift_crio_use_rpm=True
#openshift_docker_systemcontainer_image_override="registry.example.com/container-engine:latest"
#openshift_crio_systemcontainer_image_override="registry.example.com/cri-o:latest"
#openshift_crio_enable_docker_gc=True
#openshift_crio_docker_gc_node_selector={'runtime': 'cri-o'}
{{- end }}
{{- end }}

{{"\n"}}# Configure authentication with basic HTPasswdPassowordIdentityProvider
# Users must be defined on master node using htpasswd command
# TODO: Add more identity providers templates
openshift_master_identity_providers=[{'name': 'htpasswd_auth', 'login': 'true', 'challenge': 'true', 'kind': 'HTPasswdPasswordIdentityProvider', 'filename': '/etc/origin/master/htpasswd'}]

# Configure default subdomain
{{- if and (.GeneratorMultiInfra) (.GeneratorUseXip) }}{{"\n"}}openshift_master_default_subdomain={{ .GeneratorInfraIpv4 }}.xip.io
{{- else }}{{"\n"}}openshift_master_default_subdomain={{ .GeneratorExtDnsWildcard }}
{{- end }}

{{- if eq .GeneratorSdnPlugin "ovs-multitenant" }}
{{"\n"}}# Configure Network SDN plugin
os_sdn_network_plugin_name='redhat/openshift-ovs-multitenant'
{{ else if eq .GeneratorSdnPlugin "ovs-networkpolicy" }}os_sdn_network_plugin_name='redhat/openshift-ovs-networkpolicy'
{{- else }}
{{- end }}

{{- if .GeneratorDisableServiceCatalog }}
{{"\n"}}# Disable service catalog
openshift_enable_service_catalog=false
template_service_broker_install=false
{{- end }}

{{"\n"}}# Configure default pod selectors
openshift_router_selector='region=infra'
openshift_registry_selector='region=infra'
osm_default_node_selector='region=primary'
{{- if .GeneratorMultiInfra  }}
{{"\n"}}openshift_hosted_registry_replicas={{ .GeneratorInfraReplicas }}
openshift_hosted_router_replicas={{ .GeneratorInfraReplicas }}
{{- end }}

{{- if and (.GeneratorNfsEnabled) (.GeneratorRegistryNativeNfs) }}
{{"\n"}}# Configure Registry storage
openshift_hosted_registry_storage_kind=nfs
openshift_hosted_registry_storage_access_modes=['ReadWriteMany']
openshift_hosted_registry_storage_nfs_directory=/exports
openshift_hosted_registry_storage_nfs_options='*(rw,root_squash)'
openshift_hosted_registry_storage_volume_name=registry
openshift_hosted_registry_storage_volume_size=20Gi
{{- end }}

{{- if and (.GeneratorGlusterfsRegEnabled) (.GeneratorRegistryCNS) }}
{{"\n"}}# Configure Registry storage
openshift_hosted_registry_storage_kind=glusterfs
{{- end }}

{{- if .GeneratorMetricsEnabled }}
{{"\n"}}# Configure Cluster Metrics
openshift_metrics_install_metrics=true
{{ if .GeneratorDeployHosa }}openshift_metrics_install_hawkular_agent=true{{ end }}
openshift_metrics_cassandra_nodeselector={"region":"infra"}
openshift_metrics_hawkular_nodeselector={"region":"infra"}
openshift_metrics_heapster_nodeselector={"region":"infra"}

{{- if and (.GeneratorNfsEnabled) (.GeneratorMetricsNativeNfs) }}
{{"\n"}}# Configure Metrics Storage
openshift_metrics_storage_kind=nfs
openshift_metrics_storage_access_modes=['ReadWriteOnce']
openshift_metrics_storage_nfs_directory=/exports
openshift_metrics_storage_nfs_options='*(rw,root_squash)'
openshift_metrics_storage_volume_name=metrics
openshift_metrics_storage_volume_size=10Gi
{{- end }}
{{- end }}

{{- if .GeneratorPrometheusEnabled }}
{{"\n"}}# Deploy Prometheus
openshift_hosted_prometheus_deploy=true

{{- if and (.GeneratorNfsEnabled) (.GeneratorPrometheusNativeNfs) }}
{{"\n"}}# Prometheus storage config 
openshift_prometheus_storage_kind=nfs 
openshift_prometheus_storage_access_modes=['ReadWriteOnce']
openshift_prometheus_storage_nfs_directory=/exports
openshift_prometheus_storage_nfs_options='*(rw,root_squash)'
openshift_prometheus_storage_volume_name=prometheus
openshift_prometheus_storage_volume_size=10Gi
openshift_prometheus_storage_labels={'storage': 'prometheus'}
openshift_prometheus_storage_type='pvc'
openshift_prometheus_storage_class=glusterfs-storage

# Storage config for prometheus-alertmanager
openshift_prometheus_alertmanager_storage_kind=nfs
openshift_prometheus_alertmanager_storage_access_modes=['ReadWriteOnce']
openshift_prometheus_alertmanager_storage_nfs_directory=/exports
openshift_prometheus_alertmanager_storage_nfs_options='*(rw,root_squash)'
openshift_prometheus_alertmanager_storage_volume_name=prometheus-alertmanager
openshift_prometheus_alertmanager_storage_volume_size=10Gi
openshift_prometheus_alertmanager_storage_labels={'storage': 'prometheus-alertmanager'}
openshift_prometheus_alertmanager_storage_type='pvc'
openshift_prometheus_alertmanager_storage_class=glusterfs-storage

# Storage config for prometheus-alertbuffer
openshift_prometheus_alertbuffer_storage_kind=nfs
openshift_prometheus_alertbuffer_storage_access_modes=['ReadWriteOnce']
openshift_prometheus_alertbuffer_storage_nfs_directory=/exports
openshift_prometheus_alertbuffer_storage_nfs_options='*(rw,root_squash)'
openshift_prometheus_alertbuffer_storage_volume_name=prometheus-alertbuffer
openshift_prometheus_alertbuffer_storage_volume_size=10Gi
openshift_prometheus_alertbuffer_storage_labels={'storage': 'prometheus-alertbuffer'}
openshift_prometheus_alertbuffer_storage_type='pvc'
openshift_prometheus_alertbuffer_storage_class=glusterfs-storage
{{- end }}
{{- end }}

{{- if .GeneratorLoggingEnabled }}
{{"\n"}}# Configure Cluster logging
openshift_logging_install_logging=false
openshift_logging_kibana_nodeselector={"region":"infra"}
openshift_logging_es_nodeselector={"region":"infra"}

{{- if and (.GeneratorNfsEnabled) (.GeneratorLoggingNativeNfs) }}
{{"\n"}}# Configure Logging Storage
openshift_logging_storage_kind=nfs
openshift_logging_storage_access_modes=['ReadWriteOnce']
openshift_logging_storage_nfs_directory=/exports
openshift_logging_storage_nfs_options='*(rw,root_squash)'
openshift_logging_storage_volume_name=logging
openshift_logging_storage_volume_size=2Gi
{{- end }}
{{- end }}

{{"\n"}}[masters]
{{- range .GeneratorMastersList }}
{{ . }}
{{- end }}

{{"\n"}}[etcd]
{{- range .GeneratorEtcdList }}
{{ . }}
{{- end }}

{{- if and (.GeneratorHaproxyEnabled) (.GeneratorMultiMaster) }}
{{"\n"}}[lb]
{{- range .GeneratorLbList }}
{{ . }}
{{- end }}
{{- end }}

{{- if .GeneratorNfsEnabled }}
{{"\n"}}[nfs]
{{- range .GeneratorNfsList }}
{{ . }}
{{- end }}
{{- end }}

{{- if .GeneratorGlusterfsEnabled }}
{{"\n"}}[glusterfs]
{{- range $key, $value := .GeneratorGlusterfsMap }}
{{ $key }} {{ $value }}
{{- end }}
{{- end }}

{{- if .GeneratorGlusterfsRegEnabled }}
{{"\n"}}[glusterfs_registry]
{{- range $key, $value := .GeneratorGlusterfsRegMap }}
{{ $key }} {{ $value }}
{{- end }}
{{- end }}

{{"\n"}}[nodes]
{{- range $key, $value := .GeneratorNodesMap }}
{{ $key }} {{ $value }}
{{- end }}
`
