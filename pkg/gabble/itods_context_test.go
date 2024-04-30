package gabble

import (
	"testing"
)

func TestConnectAgent(t *testing.T) {
	gabbleTestConfiguration, err := EstablishGabbleTestConfig()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	unitUtils, err := NewUnitUtils()
	irodsAccount := unitUtils.Build_test1_account()

	icf := IrodsContextFactory{}

	agentConnection := icf.ConnectAgent(irodsAccount)
	if a

}
