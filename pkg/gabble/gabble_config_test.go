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
