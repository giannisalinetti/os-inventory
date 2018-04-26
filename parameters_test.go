package main

import "testing"

func TestCheckDeploymentType(t *testing.T) {
	i := New(defaults)
	tests := []string{"dummy", "Origin", "enter prise", "origin", ""}
	for _, testValue := range tests {
		i.GeneratorDeploymentType = testValue
		err := i.CheckDeploymentType()
		if (testValue != "origin" && testValue != "enterprise") && err == nil {
			t.Error("CheckDeploymentType testing error.")
		}
	}
}

func TestCheckInstallVersion(t *testing.T) {
	i := New(defaults)
	validVersions := []string{"v3.4", "v3.5", "v3.6", "v3.7", "v3.9", "v3.10", "v3.11"}
	tests := []string{"v1.2", "3.9", "v3.0", "v3.6"}
	for _, testValue := range tests {
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
