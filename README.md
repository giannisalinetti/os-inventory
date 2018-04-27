# os-inventory: an OpenShift installation inventory generator

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

## TODO

- Create new entries in the template for Container Native Storage (CNS)
- Create more sanity checks on string and int values
- Implement a good code testing


