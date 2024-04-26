// low level networking for iRODS agent connections
// for networking tools in Go consult: https://awesome-go.com/networking/
package gabble

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

// low level props for a connection to an agent
type ConSpec struct {
	Host             string // host name
	Port             int    // port
	TimeoutSeconds   int    // timeout in seconds, leave as zero for no timeout specified
	UseMPTCP         bool   // use multipath tcp if available
	KeepaliveSeconds int    // keep alive in seconds, leave as 0 for default or set to negative for no keep-alive
}

// AgentConnection  represents the low level connection to an iRODS agent
type AgentConnection struct {
	conSpec    ConSpec         // conspec used to create this connection
	context    context.Context // context shared among functions that ties to the underlying connection
	connection net.Conn        // connection associated with the agent, the underlying TCP connection to the iRRODS agent
}

// get the ConSpec() elements that are derived from the underlying config
func ConSpecFromConfig() (ConSpec, error) {
	log.Info().Msg("ConSpecFromConfig()")
	gabbleConfiguration, error := EstablishGabbleConfig()
	if error != nil {
		return ConSpec{}, error
	}
	conSpec := ConSpec{}
	conSpec.KeepaliveSeconds = gabbleConfiguration.ConnectionKeepaliveSeconds
	conSpec.TimeoutSeconds = gabbleConfiguration.ConnectionTimeoutSeconds
	conSpec.UseMPTCP = gabbleConfiguration.ConnectionUseMultipathTcp
	return conSpec, nil

}

func (ac *AgentConnection) ConnectAgent(conSpec ConSpec, context context.Context) error {
	log.Info().Msg("ConnectAgent()")
	timeout, error := translateSecondsToDuration(conSpec.TimeoutSeconds)

	if error != nil {
		log.Error().Msg(error.Error())
		return error
	}

	keepalive, error := translateSecondsToDuration(conSpec.KeepaliveSeconds)

	if error != nil {
		log.Error().Msg(error.Error())
		return error
	}

	dialer := net.Dialer{
		Timeout:   timeout,
		KeepAlive: keepalive,
	}

	addr := fmt.Sprintf("%s:%d", conSpec.Host, conSpec.Port)
	conn, err := dialer.DialContext(context, "tcp", addr)
	//conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println(err)
	}

	ac.conSpec = conSpec
	ac.context = context
	ac.connection = conn

	return err

}

// turn the timeout expressed as seconds (or as 0 for no timeout) into a time value
func translateSecondsToDuration(timeout int) (time.Duration, error) {
	timestr := strconv.Itoa(timeout) + "s"
	dur, error := time.ParseDuration(timestr)
	return dur, error

}
