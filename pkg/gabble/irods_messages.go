// irods messages that are sent and received by packer
package gabble

const XMLOPEN = "<"
const XMLCLOSE = ">"
const OPENENDTAG = "</"

/*
An IrodsMessage interface defines the methods that each iRODS protocol operation provides. Each protocol is
defined as a struct with all of the possible values, and the various 'proto_' implementations are responsible for
turning this struct into the correct representation to be incorporated into the iRODS message. This is typically a binary
record but in some cases may be a network order value or an XML-like value
*/
type IrodsProtocolDefinition interface {
	Validate() error
	PackMessage(pack StartupPack) ([]byte, error)
}

type IrodsMessage struct {
}
