package gabble

import (
	"context"
	"testing"
)

func TestConnect(t *testing.T) {
	conSpec := ConSpec{
		Host:             "icat.example.org",
		Port:             1247,
		TimeoutSeconds:   30,
		UseMPTCP:         true,
		KeepaliveSeconds: 30,
	}

	context := context.Background()

	agentConnection, err := ConnectAgent(conSpec, context)

	if err != nil {
		t.Error(err)
	}

	if agentConnection.ConSpec.Host != conSpec.Host {
		t.Fail()
	}

}
