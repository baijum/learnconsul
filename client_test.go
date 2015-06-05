package main

import (
	"net/http"
	"testing"

	consulapi "github.com/hashicorp/consul/api"
)

func TestDefaultConfig(t *testing.T) {
	config := consulapi.DefaultConfig()
	if config.Address == "127.0.0.1:8500" {
		t.Log("The default Address is 127.0.0.1:8500")
	} else {
		t.Error("The default Address changed:", config.Address)
	}
	if config.Scheme == "http" {
		t.Log("The default Scheme is http")
	} else {
		t.Error("The default Scheme changed:", config.Scheme)
	}
	if config.Datacenter == "" {
		t.Log("The default Datacenter is empty")
	} else {
		t.Error("The default Datacenter is not empty:", config.Datacenter)
	}
	if config.HttpClient == http.DefaultClient {
		t.Log("The default http client is http.DefaultClient")
	} else {
		t.Error("The default http client is not http.DefaultClient")
	}
	if config.HttpAuth == nil {
		t.Log("The default HTTP Basic Auth is empty")
	} else {
		t.Error("The default HTTP Basic Auth is not empty", config.HttpAuth)
	}
	if config.WaitTime == 0 {
		t.Log("The default WaitTime is 0")
	} else {
		t.Error("The default WaitTime is not 0", config.WaitTime)
	}

}

func TestNewClient(t *testing.T) {
	config := consulapi.DefaultConfig()
	consul, _ := consulapi.NewClient(config)
	_ = consul
}
