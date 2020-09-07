package database

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/vipin23/vmart-api/api_server"
	"github.com/vipin23/vmart-api/models"
)

func InitData() {
	AddAdminResourceOperations()
	//AddAdminRole()
	// AddAdminUser()
	// AddAdminProfile()
}

func AddAdminResourceOperations() {
	routes := api_server.DefaultServer.App.Routes()
	mapRoutes := map[string]*models.ResourceOperation{}
	for _, route := range routes {
		res := models.ResourceOperation{
			Description: GetFunctionName(route.Handlers[0]),
			HTTPMethod:  route.Method,
			Path:        route.Path,
		}
		mapRoutes[res.HTTPMethod+":"+res.Path] = &res
		fmt.Printf("%s %s %s\n", route.Method, route.Path, GetFunctionName(route.Handlers[0]))
	}
	for _, v := range mapRoutes {
		_, err := v.Upsert()
		if err != nil {
			fmt.Printf("FAILED to Save or Update : %s %s %s\n", v.HTTPMethod, v.Path, v.Description)
		}
	}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
