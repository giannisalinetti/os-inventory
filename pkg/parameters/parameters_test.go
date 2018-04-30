package parameters

import (
	"github.com/giannisalinetti/os-inventory/pkg/defaults"
	"testing"
)

func TestCheckDeploymentType(t *testing.T) {
	i := New(defaults.DefaultCfg)
	badTests := []string{"dummy", "Origin", "enter prise", "origin", ""}
	for _, testValue := range badTests {
		i.GeneratorDeploymentType = testValue
		err := i.CheckDeploymentType()
		if (testValue != "origin" && testValue != "enterprise") && err == nil {
			t.Error("CheckDeploymentType testing error.")
		}
	}
}

func TestCheckInstallVersion(t *testing.T) {
	i := New(defaults.DefaultCfg)
	validVersions := []string{"v3.4", "v3.5", "v3.6", "v3.7", "v3.9", "v3.10", "v3.11"}
	badTests := []string{"v1.2", "3.9", "v3.0", "v3.6"}
	for _, testValue := range badTests {
		i.GeneratorInstallVersion = testValue
		err := i.CheckInstallVersion()
		if err == nil {
			for _, valid := range validVersions {
				if valid != testValue {
					continue
				} else {
					return
				}
			}
			t.Error("CheckInstallVersion testing error.")
		}
	}
}

func TestCheckClusterMethod(t *testing.T) {
	i := New(defaults.DefaultCfg)
	badTests := []string{"parallel", "NATIVE", "Native", "pcs"}
	for _, testValue := range badTests {
		i.GeneratorClusterMethod = testValue
		err := i.CheckClusterMethod()
		if testValue != "native" && err == nil {
			t.Error("CheckClusterMethod testing error.")
		}
	}
}

func TestCheckInfraIpv4(t *testing.T) {
	i := New(defaults.DefaultCfg)
	validAddr := []string{"192.168.1.20", "127.0.0.1", "172.25.250.10"}
	badAddr := []string{"327.0.0.1", "302.200.1", "0.0,12", "a string"}
	for _, testValue := range badAddr {
		i.GeneratorInfraIpv4 = testValue
		err := i.CheckInfraIpv4()
		if err == nil {
			for _, valid := range validAddr {
				if valid != testValue {
					continue
				} else {
					return
				}
			}
			t.Error("CheckInfraIpv4 tesing error.")
		}
	}
}

func TestCheckSdnPlugin(t *testing.T) {
	i := New(defaults.DefaultCfg)
	validPlugins := []string{"ovs-subnet", "ovs-multitenant", "ovs-networkpolicy"}
	badTests := []string{"ovs-vxlan", "dummy", "Ovs-MultiTenant", "ovs_networkpolicy"}
	for _, testValue := range badTests {
		i.GeneratorSdnPlugin = testValue
		err := i.CheckSdnPlugin()
		if err == nil {
			for _, valid := range validPlugins {
				if valid != testValue {
					continue
				} else {
					return
				}
			}
			t.Error("CheckSdnPlugin testing error.")
		}
	}
}
