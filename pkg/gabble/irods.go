// irods is the central entry point for a gabble program. Irods represents a shared
// object (thread safe) that is a factory for connections and a central configuration point
package gabble

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
)

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
	ProxyUserName          string
	ProxyZone              string
	AuthType               AuthType
}

type IrodsContext struct {
	gabbleConfiguration GabbleConfiguration // the general properties and settings that configure the client operations
	agentConnection     *AgentConnection    // the current connection to the iRODS agent, if initialized (nil otherwise)
	irodsAccount        IrodsAccount        // The session inforation used to connect to an iRODS agent
}

type Irods struct {
	gabbleConfiguration GabbleConfiguration // the general properties and settings that configure the client operations
}

// construct a new Irods which is the main entry point to Gabble. Using this factory
// one can create agent connections, access configuration information and other tasks
func NewIrods() (*Irods, error) {
	log.Info().Msg("NewIrods()")
	irods := new(Irods)
	config, err := EstablishGabbleConfig()
	irods.gabbleConfiguration = config

	if err != nil {
		return nil, err
	}

	return irods, nil

}

// main method from the IrodsContextFactory, creates a new, connected IrodsContext that represents the client
// side of an iRODS agent connection and its configuration
func (i *Irods) ConnectAgent(irodsAccount IrodsAccount, overrideGabbleConfiguration *GabbleConfiguration) (*IrodsContext, error) {
	log.Info().Msg("ConnectAgent()")

	if overrideGabbleConfiguration == nil {
		config, err := EstablishGabbleConfig()
		if err != nil {
			return nil, err
		}
		i.gabbleConfiguration = config
	} else {
		i.gabbleConfiguration = *overrideGabbleConfiguration
	}

	context := IrodsContext{}
	context.irodsAccount = irodsAccount
	context.gabbleConfiguration = i.gabbleConfiguration
	err := connectionInitialization(&context)
	if err != nil {
		log.Error().
			Str("host", irodsAccount.Host).
			Err(errors.New("connection error")).
			Msg("unable to connect agent")
		return nil, err
	}
	log.Debug().Msg("connection established")
	return &context, nil
}

// given an IrodsContext, take the configuration and iRODS account information and produce a live
// connection to an iRODS Agent (go through the authentication and other initialization steps, SSL tunnel, etc)
func connectionInitialization(irodsContext *IrodsContext) error {
	log.Info().Msg("connectionInitialization()")
	// do the tcp connection
	conSpec, err := ConSpecFromConfig()

	if err != nil {
		return err
	}

	conSpec.Host = irodsContext.irodsAccount.Host
	conSpec.Port = irodsContext.irodsAccount.Port

	context := context.Background()
	agentConnection := AgentConnection{}
	agentConnection.context = context
	log.Debug().Msg("connecting...")
	err = agentConnection.ConnectAgent(conSpec, context)
	if err != nil {
		return err
	}
	irodsContext.agentConnection = &agentConnection
	// ssl negotiation? set negotiation and encryption
	// check /lib/core/src/rcConnect.cpp...
	// startup pack
	log.Debug().Msg("connected tcp...now send startup pack")
	return nil

}

func sendStartupPack(irodsContext *IrodsContext) error {
	log.Debug().Msg("sendStartupPack()")
	return nil
}
