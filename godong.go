package godong

import (
	"errors"
	"fmt"
	"net/http"
	"os"
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
		routePath := getRoutePath(method, controllerName)
		actionMap := fmt.Sprintf("%s%s%s", controllerName, actionMapSeparator, method.Name)

		if actionMap == DefaultAction {
			handleRoute(reflectValue, method, controllerName, actionSlash, actionMap)
		}

		handleRoute(reflectValue, method, controllerName, routePath, actionMap)
	}
}

func getRoutePath(method reflect.Method, controllerName string) string {
	validMethodName := (func() string {
		result := method.Name

		result = strings.Replace(result, Prefix, "", -1)
		result = strings.Replace(result, Separator, actionSlash, -1)

		return result
	}())

	routePath := actionSlash + controllerName + validMethodName

	if UrlMode == UrlModeDashed {
		reg, err := regexp.Compile("([a-z])([A-Z])")
		if err != nil {
			return routePath
		}

		routePath = strings.ToLower(reg.ReplaceAllString(routePath, "$1-$2"))
	}

	return routePath
}

func handleRoute(reflectValue reflect.Value, method reflect.Method, controllerName, routePath string, actionMap string) {
	if Debug {
		fmt.Println("route", routePath)
	}

	var err error

	errorInvalidActionSchema := errors.New(fmt.Sprintf("Invalid action parameter.\n      Action %s should be written in:\n      func (d *%s) %s (w http.ResponseWriter, r *http.Request)", actionMap, controllerName, method.Name))
	methodBody := reflectValue.MethodByName(method.Name)

	if methodBody.Type().NumIn() != 2 {
		err = errorInvalidActionSchema
	} else if methodBody.Type().In(0).String() != "http.ResponseWriter" {
		err = errorInvalidActionSchema
	} else if methodBody.Type().In(1).String() != "*http.Request" {
		err = errorInvalidActionSchema
	}

	if err != nil {
		fmt.Println("error", err.Error())
		os.Exit(0)
	}

	methodHandler := methodBody.Interface().(func(w http.ResponseWriter, r *http.Request))
	http.HandleFunc(routePath, methodHandler)

	if Debug {
		fmt.Println("   ->", actionMap)
	}
}
