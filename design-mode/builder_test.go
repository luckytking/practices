/**
*@Author:qrTang
*@Date:2022/10/15
**/

package design_mode

import "testing"

func TestNewResourcePoolConfig(t *testing.T) {
	resourcePool := NewResourcePool1("dbconnectionpool", 16, 8, 10)
	t.Log(resourcePool)
}

func TestNewResourcePoolConfigSetter(t *testing.T) {
	resourcePool := NewResourcePool2("dbconnectionpool").SetMaxTotal(16).SetMinIdle(8)
	t.Log(resourcePool)
}

func TestBuilder(t *testing.T) {
	resourcePool := Builder().SetName("dbconnectionpool").SetMaxTotal(16).SetMinIdle(8).Build()
	t.Log(resourcePool)
}
