// an irods_context is the central entry point for a gabble program. An irods_context represents a shared
// object that is passed
package gabble

import "github.com/rs/zerolog/log"

// iRODS Authentication types
type AuthType int

const (
	STANDARD_AUTH AuthType = iota
)

// Defines user identity and target iRODS server. Contains information about a session with an iRODS server agent
type IrodsAccount struct {
	Host                   string
	Port                   int
	Zone                   string
	DefaultStorageResource string
	UserName               string
	Password               string
	AuthType               AuthType
}

type IrodsContext struct {
	gabbleConfiguration GabbleConfiguration // the general properties and settings that configure the client operations
	agentConnection     AgentConnection     // the current connection to the iRODS agent, if initialized (nil otherwise)
	irodsAccount        IrodsAccount        // The session inforation used to connect to an iRODS agent
}

type IrodsContextFactory struct {
	gabbleConfiguration GabbleConfiguration // the general properties and settings that configure the client operations
}

// construct a new IrodsContextFactory which is the main entry point to Gabble. Using this factory
// one can create agent connections, access configuration information and other tasks
func NewIrodsContextFactory() (*IrodsContextFactory, error) {
	log.Info().Msg("NewIrodsContextFactory()")
	icf := new(IrodsContextFactory)
	config, err := EstablishGabbleConfig()
	icf.gabbleConfiguration = config

	if err != nil {
		return nil, err
	}

	return icf, nil

}

func (ic *IrodsContextFactory) ConnectAgent(irodsAccount IrodsAccount) (*IrodsContext, error) {
	log.Info().Msg("ConnectAgent()")
	config, err := EstablishGabbleConfig()
	if err != nil {
		return err
	}
	ic.gabbleConfiguration = config
	return nil
}
