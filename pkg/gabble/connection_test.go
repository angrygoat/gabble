package gabble

import (
	"context"
	"testing"
)

// You can use testing.T, if you want to test the code without benchmarking
func setupSuite(t testing.T) func(t testing.T) {
	t.Log("setupSuite()")

	// Return a function to teardown the test
	return func(t testing.T) {
		t.Log("cleanup()")
	}
}

func TestConnect(t *testing.T) {
	gabbleTestConfiguration, err := EstablishGabbleTestConfig()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	conSpec, err := ConSpecFromConfig()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	conSpec.Host = gabbleTestConfiguration.Host
	conSpec.Port = gabbleTestConfiguration.Port

	context := context.Background()

	agentConnection, err := ConnectAgent(conSpec, context)

	if err != nil {
		t.Error(err)
	}

	if agentConnection.ConSpec.Host != conSpec.Host {
		t.Fail()
	}

}

func TestConSpecFromConfig(t *testing.T) {

	gabbleConfiguration, err := EstablishGabbleConfig()

	if err != nil {
		t.Errorf("error getting gabble config: %v", err)
	}

	actual, err := ConSpecFromConfig()

	if err != nil {
		t.Errorf("error getting conspec from config: %v", err)
	}

	if actual.TimeoutSeconds != gabbleConfiguration.ConnectionTimeoutSeconds {
		t.Errorf("timoutSeconds mismatch")
	}

	if actual.KeepaliveSeconds != gabbleConfiguration.ConnectionKeepaliveSeconds {
		t.Errorf("timoutSeconds mismatch")
	}

}
