package gabble

import (
	"bytes"
	"text/template"
)

// startup pack, see rodsDef.h

const STARTUP_PACK_TEMPLATE = `
<StartupPack_PI><irodsProt>{{.IrodsProtocol}}</irodsProt>
<reconnFlag>{{.ReconnFlag}}</reconnFlag>
<connectCnt>{{.ConnectCnt}}</connectCnt>
<proxyUser>{{.ProxyUserName}}</proxyUser>
<proxyRcatZone>{{.ProxyZone}}</proxyRcatZone>
<clientUser>{{.UserName}}</clientUser>
<clientRcatZone>{{.Zone}}</clientRcatZone>
<relVersion>{{.RelVersion}}</relVersion>
<apiVersion>{{.ApiVersion}}</apiVersion>
<option> {{range .StartupOptions}}{{.}} {{end}}</option>
</StartupPack_PI>
`

type StartupPack struct {
	IrodsProtocol  int
	ReconnFlag     int
	ConnectCnt     int
	ProxyUserName  string
	ProxyZone      string
	UserName       string
	Zone           string
	RelVersion     string
	ApiVersion     string
	StartupOptions []string
	RenderStyle    string
	IsStateful     bool
}

func (prot *StartupPack) NewStartupPack() *StartupPack {
	prot.IsStateful = false
	prot.RenderStyle = RENDER_PROTO_AS_XML
	return prot
}

func (prot *StartupPack) Validate() error {
	return nil
}

func (prot *StartupPack) PackMessage(pack StartupPack) ([]byte, error) {

	tmpl, err := template.New("startup").Parse(STARTUP_PACK_TEMPLATE)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, prot)
	return b.Bytes(), nil
}

/*

<StartupPack_PI><irodsProt>1</irodsProt>
<reconnFlag>0</reconnFlag>
<connectCnt>0</connectCnt>
<proxyUser>test1</proxyUser>
<proxyRcatZone>tempZone</proxyRcatZone>
<clientUser>test1</clientUser>
<clientRcatZone>tempZone</clientRcatZone>
<relVersion>rods3.2</relVersion>
<apiVersion>d</apiVersion>
<option>jargonrequest_server_negotiation</option>
</StartupPack_PI>

*/
