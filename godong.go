package godong

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

const (
	UrlModeDashed      int    = 0
	UrlModeCapitalized int    = 1
	actionSlash        string = "/"
	actionMapSeparator string = "."
)

var (
	Debug         bool   = false
	Prefix        string = "Action"
	Separator     string = "_"
	DefaultAction string = ""
	UrlMode       int    = UrlModeDashed
)

func Route(o interface{}) {
	var reflectType = reflect.TypeOf(o)
	var reflectValue = reflect.ValueOf(o)

	var controllerName = reflect.Indirect(reflectValue).Type().Name()

	for i := 0; i < reflectType.NumMethod(); i++ {
		method := reflectType.Method(i)
		methodName := getMethodName(method)
		methodBody := reflectValue.MethodByName(method.Name)
		routePath := getRoutePath(actionSlash, controllerName, methodName)
		actionMap := strings.Join([]string{controllerName, method.Name}, actionMapSeparator)

		if actionMap == DefaultAction {
			handleRoute(actionSlash, methodBody, actionMap)
		}

		handleRoute(routePath, methodBody, actionMap)
	}
}

func getMethodName(method reflect.Method) string {
	methodName := method.Name

	methodName = strings.Replace(methodName, Prefix, "", -1)
	methodName = strings.Replace(methodName, Separator, actionSlash, -1)

	return methodName
}

func getRoutePath(actionSlash string, controllerName string, methodName string) string {
	routePath := actionSlash + controllerName + methodName

	if UrlMode == UrlModeDashed {
		reg, err := regexp.Compile("([a-z])([A-Z])")
		if err != nil {
			return routePath
		}

		routePath = strings.ToLower(reg.ReplaceAllString(routePath, "$1-$2"))
	}

	return routePath
}

func handleRoute(routePath string, methodBody reflect.Value, actionMap string) {
	methodHandler := methodBody.Interface().(func(w http.ResponseWriter, r *http.Request))
	http.HandleFunc(routePath, methodHandler)

	if Debug {
		fmt.Println("route", routePath)
		fmt.Println("   ->", actionMap)
	}
}
