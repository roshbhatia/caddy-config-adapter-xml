package caddyconfigadapterxml

import (
	"encoding/json"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/clbanning/mxj"
)

var (
	_ caddyconfig.Adapter = (*XmlAdapter)(nil)
)

func init(){
	caddyconfig.RegisterAdapter("xml", XmlAdapter{})
}

type XmlAdapter struct{}

func (adapter XmlAdapter) Adapt (body []byte, options map[string]interface{}) ([]byte, []caddyconfig.Warning, error) {
	m, err := mxj.NewMapXml(body)
	if err != nil {
		return nil, nil, err
	}

	configAsJson, err := json.Marshal(m)
	if err != nil {
		return nil, nil, err
	}

	return configAsJson, nil, nil
}
