# os-inventory: An OpenShift installation inventory generator

## Why?

Red Hat OpenShift Container Platform is installed using Ansible. The community project
https://github.com/openshift/openshift-ansible provides related playbooks and roles.
To deploy OpenShift a lot of parameters can be set using inventory variables.
To speed up the deployment process, a tool that creates inventories from a template can
be really helpful. Written in Go, si single binary tool can generate custom inventories
starting from a simple yaml file with a minimun set of parameters.
Default values are provided to generate a basic inventory from scratch.


## Usage

To create a basic inventory simply run the command without flags:

```
$ os-inventory
```

To print an help:

```
$ os-inventory --help
```

To dump all the default values in YAML format to stdout:

```
$ os-inventory --show-defaults
```

To load configurations from a YAML file:

```
$ os-inventory --load-yaml test.yml
```

The generated inventory is always redirected to stdout, thus, to create a custom file,
just redirect the output:

```
$ os-inventory --load-yaml test.yml > /tmp/myinventory
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

## TODO

- Create new entries in the template for Container Native Storage (CNS)
- Implement spf13/viper to manage configuration (also by custom flags)
- Create quality checks on string and int values


