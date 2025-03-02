// low level networking for iRODS agent connections
// for networking tools in Go consult: https://awesome-go.com/networking/

// TODO: tcp keep alive and other params
// TODO: add local ip bind into gabble config

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
	LocalIp          string
}

// AgentConnection  represents the low level connection to an iRODS agent
type AgentConnection struct {
	conSpec    ConSpec         // conspec used to create this connection
	context    context.Context // context shared among functions that ties to the underlying connection
	connection net.Conn        // connection associated with the agent, the underlying TCP connection to the iRODS agent
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
	conSpec.LocalIp = gabbleConfiguration.LocalIp
	return conSpec, nil

}

func (ac *AgentConnection) ConnectAgent(conSpec ConSpec, context context.Context) error {
	log.Info().Msg("ConnectAgent()")
	timeout, err := translateSecondsToDuration(conSpec.TimeoutSeconds)

	if err != nil {
		log.Error().Err(err).Msgf("cannot tanslate timout: %v", conSpec)
		return err
	}

	keepalive, err := translateSecondsToDuration(conSpec.KeepaliveSeconds)

	if err != nil {
		log.Error().Err(err).Msgf("cannot tanslate keepalive: %v", conSpec)
		return err
	}

	var localAddr net.Addr
	if conSpec.LocalIp != "" {
		myAddr, err := net.ResolveTCPAddr("tcp", conSpec.LocalIp)
		if err != nil {
			log.Error().Err(err).Msgf("unable to resolve local TCP address: %v", conSpec)
		}
		localAddr = myAddr
	}

	dialer := net.Dialer{
		Timeout:   timeout,
		KeepAlive: keepalive,
		LocalAddr: localAddr,
	}

	addr := fmt.Sprintf("%s:%d", conSpec.Host, conSpec.Port)
	conn, err := dialer.DialContext(context, "tcp", addr)
	//conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Err(err).Msgf("cannot connect to agent: %v", conSpec)
	}

	ac.conSpec = conSpec
	ac.context = context
	ac.connection = conn

	return err

}

// turn the timeout expressed as seconds (or as 0 for no timeout) into a time value
func translateSecondsToDuration(timeout int) (time.Duration, error) {
	timestr := strconv.Itoa(timeout) + "s"
	dur, err := time.ParseDuration(timestr)
	return dur, err

}
