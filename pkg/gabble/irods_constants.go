// constants used across library
package gabble

// Auth and connection related constants

const AuthOpenidScheme = "openid"
const PamInteractiveScheme = "pam_interactive"
const PamPasswordScheme = "pam_password"

const DefaultNegotiation = "CS_NEG_REQUIRE"
const CsNegFailure = "CS_NEG_FAILURE"
const CsUseSsl = "CS_NEG_USE_SSL"
const CsUseTcpIp = "CS_NEG_USE_TCP"
const CsNegRequire = "CS_NEG_REQUIRE"
const CsNegRefuse = "CS_NEG_REFUSE"
const CsNegDontCare = "CS_NEG_DONT_CARE"

const ReconnectFlag = 200

const RequestNegotiation = "request_server_negotiation"

type IrodsProtocol int

const (
	NativeProtocol = iota
	XmlProtocol
)

func (p IrodsProtocol) String() string {
	switch p {
	case NativeProtocol:
		return "1"
	case XmlProtocol:
		return "0"
	}
	return "1"
}
