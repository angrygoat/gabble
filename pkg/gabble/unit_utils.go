// utilities to support testing
package gabble

type UnitUtils struct {
	gabbleConfiguration     GabbleConfiguration // the general properties and settings that configure the client operations
	gabbleTestConfiguration GabbleTestConfiguration
}

// get an irods account for the first test user
func (unitUtils *UnitUtils) Build_test1_account() IrodsAccount {
	irodsAccount := IrodsAccount{}
	irodsAccount.AuthType = STANDARD_AUTH
	irodsAccount.DefaultStorageResource = ""
	irodsAccount.Host = unitUtils.gabbleTestConfiguration.Host
	irodsAccount.Password = unitUtils.gabbleTestConfiguration.Password1
	irodsAccount.Port = unitUtils.gabbleTestConfiguration.Port
	irodsAccount.UserName = unitUtils.gabbleTestConfiguration.User1
	irodsAccount.Zone = unitUtils.gabbleTestConfiguration.Zone
	return irodsAccount
}

// construct a new UnitUtils that can be used to manage unit tests
func NewUnitUtils() (*UnitUtils, error) {
	unitUtils := new(UnitUtils)
	gabbleConfiguration, err := EstablishGabbleConfig()

	if err != nil {
		return nil, err
	}

	gabbleTestConfiguration, err := EstablishGabbleTestConfig()

	if err != nil {
		return nil, err
	}

	unitUtils.gabbleConfiguration = gabbleConfiguration
	unitUtils.gabbleTestConfiguration = gabbleTestConfiguration
	return unitUtils, nil

}
