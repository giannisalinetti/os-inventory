package parameters

import (
	"errors"
	"github.com/giannisalinetti/os-inventory/pkg/defaults"
	"testing"
)

func TestCheckDeploymentType(t *testing.T) {
	i := New(defaults.DefaultCfg)
	checkErr := errors.New("Invalid deployment type.")
	var tests = []struct {
		args        string
		expectedErr error
	}{
		{"origin", nil},
		{"enterprise", nil},
		{"", checkErr},
		{"dummy", checkErr},
		{"enter prise", checkErr},
		{"Origin", checkErr},
	}
	for _, test := range tests {
		i.GeneratorDeploymentType = test.args
		err := i.CheckDeploymentType()
		if test.expectedErr != nil {
			if err.Error() != test.expectedErr.Error() {
				t.Error("CheckDeploymentType testing error.")
			}
		} else {
			if err != test.expectedErr {
				t.Error("CheckDeploymentType testing error.")
			}
		}
	}
}

func TestCheckInstallVersion(t *testing.T) {
	i := New(defaults.DefaultCfg)
	checkErr := errors.New("Invalid or unsupported version.")
	var tests = []struct {
		args        string
		expectedErr error
	}{
		{"v3.4", nil},
		{"v3.5", nil},
		{"v3.6", nil},
		{"v3.7", nil},
		{"v3.9", nil},
		{"v3.10", nil},
		{"v3.11", nil},
		{"v1.2", checkErr},
		{"3.9", checkErr},
		{"v3.0", checkErr},
		{"", checkErr},
	}
	for _, test := range tests {
		i.GeneratorInstallVersion = test.args
		err := i.CheckInstallVersion()
		if test.expectedErr != nil {
			if err.Error() != test.expectedErr.Error() {
				t.Error("CheckInstallVersion testing error.")
			}
		} else {
			if err != test.expectedErr {
				t.Error("CheckInstallVersion testing error.")
			}
		}
	}
}

func TestCheckClusterMethod(t *testing.T) {
	i := New(defaults.DefaultCfg)
	checkErr := errors.New("Invalid cluster method.")
	var tests = []struct {
		args        string
		expectedErr error
	}{
		{"native", nil},
		{"parallel", checkErr},
		{"NATIVE", checkErr},
		{"", checkErr},
		{"pcs", checkErr},
	}
	for _, test := range tests {
		i.GeneratorClusterMethod = test.args
		err := i.CheckClusterMethod()
		if test.expectedErr != nil {
			if err.Error() != test.expectedErr.Error() {
				t.Error("CheckClusterMethod testing error.")
			}
		} else {
			if err != test.expectedErr {
				t.Error("CheckClusterMethod testing error.")
			}
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
