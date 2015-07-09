package godong

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

var (
	Debug         bool   = false
	Prefix        string = "Action"
	Separator     string = "_"
	DefaultAction string = ""
	UrlMode       string = "dashed"
)

const (
	slash string = "/"
)

func Route(o interface{}) {
	var reflectType = reflect.TypeOf(o)
	var reflectValue = reflect.ValueOf(o)

	var controllerName = reflect.Indirect(reflectValue).Type().Name()

	for i := 0; i < reflectType.NumMethod(); i++ {
		method := reflectType.Method(i)
		methodName := getMethodName(method)
		methodBody := reflectValue.MethodByName(method.Name)
		routePath := getRoutePath(slash, controllerName, methodName)

		if i == 0 && DefaultAction == strings.Replace(methodName, slash, "", -1) {
			handleRoute(slash, methodBody)
		}

		handleRoute(routePath, methodBody)
	}
}

func getMethodName(method reflect.Method) string {
	methodName := method.Name

	methodName = strings.Replace(methodName, Prefix, "", -1)
	methodName = strings.Replace(methodName, Separator, slash, -1)

	return methodName
}

func getRoutePath(slash string, controllerName string, methodName string) string {
	routePath := slash + controllerName + methodName

	if UrlMode == "dashed" {
		reg, err := regexp.Compile("([a-z])([A-Z])")
		if err != nil {
			return routePath
		}

		routePath = strings.ToLower(reg.ReplaceAllString(routePath, "$1-$2"))
	}

	return routePath
}

func handleRoute(routePath string, methodBody reflect.Value) {
	http.HandleFunc(routePath, func(w http.ResponseWriter, r *http.Request) {
		methodHandler := methodBody.Interface().(func(w http.ResponseWriter, r *http.Request))
		methodHandler(w, r)
	})

	if Debug {
		fmt.Println("route", routePath)
	}
}
