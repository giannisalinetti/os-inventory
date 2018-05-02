package parameters

import (
	"errors"
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
	checkErr := errors.New("Invalid IPv4 address.")
	var tests = []struct {
		args        string
		expectedErr error
	}{
		{"192.168.1.20", nil},
		{"127.0.0.1", nil},
		{"172.25.250.10", nil},
		{"327.0.0.1", checkErr},
		{"302.200.1", checkErr},
		{"0.0,12", checkErr},
		{"a string", checkErr},
	}
	for _, test := range tests {
		i.GeneratorInfraIpv4 = test.args
		err := i.CheckInfraIpv4()
		if test.expectedErr != nil {
			if err.Error() != test.expectedErr.Error() {
				t.Error("CheckInfraIpv4 testing error.")
			}
		} else {
			if err != test.expectedErr {
				t.Error("CheckInfraIpv4 testing error.")
			}
		}
	}
}

func TestCheckSdnPlugin(t *testing.T) {
	i := New(defaults.DefaultCfg)
	checkErr := errors.New("Invalid SDN plugin.")
	var tests = []struct {
		args        string
		expectedErr error
	}{
		{"ovs-subnet", nil},
		{"ovs-multitenant", nil},
		{"ovs-networkpolicy", nil},
		{"ovs-vxlan", checkErr},
		{"dummy", checkErr},
		{"Ovs-MultiTenant", checkErr},
		{"ovs_networkpolicy", checkErr},
	}
	for _, test := range tests {
		i.GeneratorSdnPlugin = test.args
		err := i.CheckSdnPlugin()
		if test.expectedErr != nil {
			if err.Error() != test.expectedErr.Error() {
				t.Error("CheckSdnPlugin testing error.")
			}
		} else {
			if err != test.expectedErr {
				t.Error("CheckSdnPlugin testing error.")
			}
		}
	}
}

func TestCheckRegistryStorage(t *testing.T) {
	i := New(defaults.DefaultCfg)

	// Test wrong combination
	i.GeneratorRegistryNativeNfs = true
	i.GeneratorRegistryCNS = true
	err := i.CheckRegistryStorage()
	if err == nil {
		t.Error("CheckRegistryStorage testing error.")
	}

	// Test good combinations
	okCombinations := [][]bool{{true, false}, {false, true}, {false, false}}
	for c, _ := range okCombinations {
		i.GeneratorRegistryNativeNfs = okCombinations[c][0]
		i.GeneratorRegistryCNS = okCombinations[c][1]
		err := i.CheckRegistryStorage()
		if err != nil {
			t.Error("CheckRegistryStorage testing error.")
		}
	}
}
