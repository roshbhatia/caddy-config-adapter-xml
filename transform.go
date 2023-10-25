package caddyconfigadapterxml


func applyTransforms(configAsXml map[string]interface{}) {
	transformListenToArray(configAsXml)
	transformRoutesToArray(configAsXml)
	transformListenToArray(configAsXml)
}

// apps.http.servers.*.listen
func transformListenToArray(configAsXml map[string]interface{}) {
	if apps, ok := configAsXml["apps"].(map[string]interface{}); ok {
		if http, ok := apps["http"].(map[string]interface{}); ok {
			if servers, ok := http["servers"].(map[string]interface{}); ok {
				for _, server := range servers {
					if serverMap, ok := server.(map[string]interface{}); ok {
						if listen, ok := serverMap["listen"].(string); ok {
							serverMap["listen"] = []string{listen}
						}
					}
				}
			}
		}
	}
}

// Path: apps.http.servers.*.routes
func transformRoutesToArray(configAsXml map[string]interface{}) {
	if httpApp, ok := configAsXml["apps"].(map[string]interface{})["http"].(map[string]interface{}); ok {
		if servers, ok := httpApp["servers"].(map[string]interface{}); ok {
			for serverKey, serverVal := range servers {
				serverMap := serverVal.(map[string]interface{})
				if routes, ok := serverMap["routes"].(map[string]interface{}); ok {
					routesArray := make([]interface{}, 0)
					for _, routeVal := range routes {
						routesArray = append(routesArray, routeVal)
					}
					serverMap["routes"] = routesArray
				}
				servers[serverKey] = serverMap
			}
		}
	}
}

// apps.tls.automation.policies
func transformPoliciesToArray(configAsXml map[string]interface{}) {
	if apps, ok := configAsXml["apps"].(map[string]interface{}); ok {
		if tls, ok := apps["tls"].(map[string]interface{}); ok {
			if automation, ok := tls["automation"].(map[string]interface{}); ok {
				if policies, ok := automation["policies"].(map[string]interface{}); ok {
					automation["policies"] = []map[string]interface{}{policies}
				}
			}
		}
	}
}
