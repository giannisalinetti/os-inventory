# os-inventory: an OpenShift installation inventory generator

## DISCLAIMER: This project is deprecated. 
Red Hat OpenShift 4 drop the Ansible deployment method and introduces a totally 
new approach in the installation of OpenShift clusters based on Terraform and a 
combination of tecniques which leverage on Red Hat CoreOS **Ignition Configs** 
on boostrap nodes and control plane nodes.
Ansible roles are now only related to worker nodes based on RHEL.

## Why?

Red Hat OpenShift Container Platform is installed using Ansible. The community project
https://github.com/openshift/openshift-ansible provides related playbooks and roles.
To deploy OpenShift a lot of parameters can be set using inventory variables.
To speed up the deployment process, a tool that creates inventories from a template can
be really helpful. Written in Go, si single binary tool can generate custom inventories
starting from a simple yaml file with a minimun set of parameters.
Default values are provided to generate a basic inventory from scratch.
Despite being a very simple tool, os-inventory implements a full working CLI that can be
extended with new features in the future, yet maintaining the KISS phylosophy.
A basic sanity check on data, especially strings, is done to avoid errors on install time.

## Build and Install

To build and install os-inventory:

```
$ make
$ sudo make install
```

This will run tests and compile the binary install it under */usr/local/bin* on your system.

## Usage

To get the generic help and overall description simply run the command without arguments:

```
$ os-inventory
```

To print a specific command help:

```
$ os-inventory <COMMAND> --help
```

To dump all the default values in YAML format to stdout:

```
$ os-inventory defaults
```

To generate an inventory and dump it to stdout:

```
$ os-inventory generate
```

To generate the inventory loading configurations from a YAML file:

```
$ os-inventory generate -f examples/example.yml
```

The generated inventory is redirected by defayult to stdout. To create a custom file
use the options *-o* or *--output*:

```
$ os-inventory generate -f example.yml -o /tmp/myinventory
```

## The YAML config

The YAML configuration file can include a different number of parameters:

```
---
deployment_type: "enterprise"
multi_master: true
haproxy_enabled: true
masters_list:
    - m1.example.com
    - m2.example.com
    - m3.example.com
etcd_list:
    - m1.example.com
    - m2.example.com
    - m3.example.com
lb_list:
    - lb.example.com
nodes_map:
    m1.example.com: ""
    m2.example.com: ""
    m3.example.com: ""
    n1.example.com: "openshift_node_labes=\"{'region': 'infra'}\""
    n2.example.com: "openshift_node_labes=\"{'region': 'infra'}\""
    n3.example.com: "openshift_node_labes=\"{'region': 'infra'}\""
    n4.example.com: "openshift_node_labes=\"{'region': 'primary'}\""
    n5.example.com: "openshift_node_labes=\"{'region': 'primary'}\""
```

A more minimalistic config file can be produces, maybe with just one custom parameter:

```
---
deployment_type: "origin"
```

It is possible to generate CNS ready configurations. This is a small example of Registry 
using Container Native Storage, a very common scenario:

```
---
deployment_type: "enterprise"
multi_master: true
haproxy_enabled: true
nfs_enabled: false
glusterfs_registry_enabled: true
registry_cns: true
registry_native_nfs: false
masters_list:
    - m1.example.com
    - m2.example.com
    - m3.example.com
etcd_list:
    - m1.example.com
    - m2.example.com
    - m3.example.com
lb_list:
    - lb.example.com
glusterfs_registry_map:
    n1.example.com: "glusterfs_devices='[ \"/dev/vdb\", \"/dev/vdc\", \"/dev/vdd\" ]'"
    n2.example.com: "glusterfs_devices='[ \"/dev/vdb\", \"/dev/vdc\", \"/dev/vdd\" ]'"
    n3.example.com: "glusterfs_devices='[ \"/dev/vdb\", \"/dev/vdc\", \"/dev/vdd\" ]'"
nodes_map:
    m1.example.com: ""
    m2.example.com: ""
    m3.example.com: ""
    n1.example.com: "openshift_node_labes=\"{'region': 'infra'}\""
    n2.example.com: "openshift_node_labes=\"{'region': 'infra'}\""
    n3.example.com: "openshift_node_labes=\"{'region': 'infra'}\""
    n4.example.com: "openshift_node_labes=\"{'region': 'primary'}\""
    n5.example.com: "openshift_node_labes=\"{'region': 'primary'}\""
```

## Defaults

Currently, the configured defaults are the following:

```
---
deployment_type: origin     # Choose between "origin" or "enterprise" deployment
ssh_user: root              # The user that connects to the nodes 
nfs_enabled: true           # Use native NFS for storage
glusterfs_enabled: false    # Enable Glusterfs based CNS or CRS
glusterfs_registry_enabled: false   # Use Glusterfs for registry storage
registry_native_nfs: true   # Use NFS for registry storage
registry_cns: false         # Enable Container Native Storage for registy
haproxy_enabled: false      # Enable HAProxy loadbalancer in HA deployments
install_version: v3.9       # Installation version
skip_checks: false          # Skip storage, memory, package checks
multi_master: false         # Multi master HA deployment
cluster_method: native      # Cluster method 
cluster_hostname: osapi.example.com # Private cluster endpoint name for API
cluster_public_hostname: osapi.example.com  # Public cluster endpoint name for API
containerized_deploy: false # Enable containerized deployment
containerized_ovs: false    # Containerize OpenvSwitch
containerized_node: false   # Containerize atomic-openshift-node service
containerized_master: false # Containerize master services (api and controllers)
containerized_etcd: false   # Containerize etcd service
system_images_registry: registry.access.redhat.com  # Default system images registry
openshift_use_crio: false   # Enable CRI-O deployment
openshift_crio_use_rpm: false   # Use rpm basesd installation for CRI-O
multi_infra: false          # Deploy multiple infra nodes
use_xip: false              # Use XIP for wildcard DNS (useful in training labs)
infra_ipv4: ""              # Infra IPv4 address for XIP based wildcard
ext_dns_wildcard: osapps.example.com    # Wildcard address (the common scenario)
sdn_plugin: ovs-multitenant # Cluster SDN plugin
disable_servicecatalog: true    # Disable the Service Catalog (from v3.7)
infra_replicas: 3           # Number of replicas for infra nodes
metrics_enabled: true       # Enable metrics subsystem
deploy_hosa: true           # Deploy Hawkular OpenShift Agent (HOSA)
metrics_native_nfs: true    # Use cluster NFS for metrics storage
prometheus_enabled: false   # Enable Prometheus deployment
prometheus_native_nfs: true # Use cluster NFS for Prometheus
logging_enabled: true       # Enable logging subsystem
logging_native_nfs: true    # Use Cluster NFS for logging
masters_list: []            # List of the master nodes
etcd_list: []               # List of the etcd nodes
lb_list: []                 # List of the load balancers
nfs_list: []                # List of cluster NFS nodes
glusterfs_map: {}           # Key/Value map of GlusterFS nodes
glusterfs_registry_map: {}  # Key/Value map of GlusterFS nodes for registry
nodes_map: {}               # Key/Value map of nodes.
```

## TODO

- Create new entries in the template for Container Native Storage (CNS)
- Create more sanity checks on string and int values
- Implement a good code testing


