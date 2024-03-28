package gabble

import (
	"testing"
)

func TestEstablishGabbleConfig(t *testing.T) {
	gabbleConfiguration, err := EstablishGabbleConfig()
	if err != nil {
		t.Errorf("error establishing gabble config: %s", err)
	}

	t.Logf("Got gabble configuration:  %+v", gabbleConfiguration)

}

func TestEstablishGabbleTestConfig(t *testing.T) {
	gabbleTestConfiguration, err := EstablishGabbleTestConfig()
	if err != nil {
		t.Errorf("error establishing gabble test config: %s", err)
	}

	if gabbleTestConfiguration.Port != 1247 {
		t.Error("did not get expected port")

	}

	t.Logf("Got gabble testing configuration:  %+v", gabbleTestConfiguration)

}
