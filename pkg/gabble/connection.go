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
	ConSpec    ConSpec         // conspec used to create this connection
	Context    context.Context // context shared among functions that ties to the underlying connection
	Connection net.Conn        // connection associated with the agent, the underlying TCP connection to the iRRODS agent
}

func ConnectAgent(conSpec ConSpec, context context.Context) (AgentConnection, error) {
	log.Info().Msg("Hello from Zerolog global logger")
	secs, _ := translateTimeout(conSpec.TimeoutSeconds)
	dialer := net.Dialer{
		Timeout: secs,
	}
	addr := fmt.Sprintf("%s:%d", conSpec.Host, conSpec.Port)
	conn, err := dialer.DialContext(context, "tcp", addr)
	//conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println(err)
	}

	agentConnection := AgentConnection{
		ConSpec:    conSpec,
		Context:    context,
		Connection: conn,
	}

	return agentConnection, err

}

// turn the timeout expressed as seconds (or as 0 for no timeout) into a time value
func translateTimeout(timeout int) (time.Duration, error) {
	timestr := strconv.Itoa(timeout) + "s"
	dur, error := time.ParseDuration(timestr)
	return dur, error

}
