package gabble

import (
	"testing"
)

func TestIrods(t *testing.T) {

	unitUtils, err := NewUnitUtils()

	actual := unitUtils.gabbleConfiguration

	if err != nil {
		t.Errorf("failed to create unit utils: %v", err)
	}

	irods, err := NewIrods()

	if err != nil {
		t.Errorf("failed to create irods: %v", err)
	}

	if irods.gabbleConfiguration != actual {
		t.Errorf("gabble configuration missing in irods")
	}
}

func TestConnectAgent(t *testing.T) {

	unitUtils, err := NewUnitUtils()
	irodsAccount := unitUtils.Build_test1_account()

	irods, err := NewIrods()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// do not override and configuration settings
	irodsContext, err := irods.ConnectAgent(irodsAccount, nil)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if irodsContext == nil {
		t.Errorf("nil irodsContext")
	}

	if irodsContext.agentConnection == nil {
		t.Error("no connection initialized")
	}
}
