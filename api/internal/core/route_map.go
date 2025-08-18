package core

import "fmt"

type routesMap map[string]RoutesInfo

func initRoutesMap() routesMap {
	routesName := make(routesMap)
	for _, route := range routesInfo {
		if _, ok := routesName[route.Name]; !ok {
			routesName[route.Name] = route
		} else {
			panic(fmt.Sprintf("路由名称重复: %s\n   %+v\n   %+v\n", route.Name, route, routesName[route.Name]))
		}
	}
	return routesName
}

func (r routesMap) Path(name string) string {
	return r[name].Path
}

func (r routesMap) Method(name string) string {
	return r[name].Method
}

func (r routesMap) Desc(name string) string {
	return r[name].Description
}

func (r routesMap) Router(name string) RoutesInfo {
	return r[name]
}
