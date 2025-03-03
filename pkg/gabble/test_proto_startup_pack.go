package gabble

import (
	"testing"
)

func TestPackMessageStartupPack(t *testing.T) {

	gabbleTestConfiguration, err := EstablishGabbleTestConfig()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	startupPack := StartupPack{
		IrodsProtocol: RENDER_PROTO_AS_XML,
		ReconnFlag: 0,
		ConnectCnt: 0,
		ProxyUserName: gabbleTestConfiguration.User1,
		ProxyZone: gabbleTestConfiguration.Zone,
		UserName: gabbleTestConfiguration.User1,
		Zone: gabbleTestConfiguration.Zone,
		RelVersion:



	}




}
