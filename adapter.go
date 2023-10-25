package caddyconfigadapterxml

import (
	"encoding/json"
	"fmt"

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

func (adapter XmlAdapter) Adapt(body []byte, options map[string]interface{}) ([]byte, []caddyconfig.Warning, error) {
	m, err := mxj.NewMapXml(body)
	if err != nil {
		return nil, nil, err
	}

	configAsXml, ok := m["caddyfile"].(map[string]interface{})
	if !ok {
		return nil, nil, fmt.Errorf("missing <caddyfile> root element in XML")
	}

	// Due to issues converting arrays from XML to JSON, we need to explicitly reshape certain elements embedded wtihin the XML map.
	// Transformations should live within `./transform.go`.
	applyTransforms(configAsXml)

	configAsJson, err := json.Marshal(configAsXml)
	if err != nil {
		return nil, nil, err
	}

	return configAsJson, nil, nil
}
