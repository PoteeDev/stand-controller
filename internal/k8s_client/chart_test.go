package k8s_client

import (
	"testing"
)

var goodChart, _ = InitChart("testName", "repo")

func TestAddRepo(t *testing.T) {
	err := goodChart.AddRepo()
	if err != nil {
		t.Fatalf(`AddRepo() %v`, err)
	}
}

func TestValidate(t *testing.T) {
	err := goodChart.Validate()
	if err != nil {
		t.Fatalf(`Validate() %v`, err)
	}
}
