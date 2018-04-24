package main

const tmpl = `# OpenShift {{ .GeneratorDeploymentType }} advanced installation
[OSEv3:children]
masters
etcd
nodes
{{ if .GeneratorNfsEnabled }}nfs{{ end }}
lb

[OSEv3:vars]
ansible_ssh_user={{ .GeneratorSshUser }}
`
