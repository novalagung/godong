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
	UrlModeDashed             int    = 0
	UrlModeCapitalized        int    = 1
	actionSlash               string = "/"
	actionSeparator           string = "_"
	actionControllerSeparator string = "."
	actionIndexName           string = "Index"
	actionPrefix              string = "Action"
)

var (
	Debug         bool   = false
	DefaultAction string = ""
	HiddenIndex   bool   = false
	UrlMode       int    = UrlModeDashed
)

func Route(o interface{}) {
	var reflectType = reflect.TypeOf(o)
	var reflectValue = reflect.ValueOf(o)

	for i := 0; i < reflectType.NumMethod(); i++ {
		handleRoute(reflectType, reflectValue, i)
	}
}

func getRoutePath(method reflect.Method, controllerName string) string {
	validMethodName := (func() string {
		result := method.Name

		result = strings.Replace(result, actionPrefix, "", -1)
		result = strings.Replace(result, actionSeparator, actionSlash, -1)

		return result
	})

	routePath := actionSlash + controllerName + validMethodName()

	if UrlMode == UrlModeDashed {
		reg, err := regexp.Compile("([a-z])([A-Z])")
		if err != nil {
			return routePath
		}

		routePath = strings.ToLower(reg.ReplaceAllString(routePath, "$1-$2"))
	}

	return routePath
}

func handleRoute(reflectType reflect.Type, reflectValue reflect.Value, i int) {
	controllerName := reflect.Indirect(reflectValue).Type().Name()
	method := reflectType.Method(i)
	methodBody := reflectValue.MethodByName(method.Name)
	routePath := getRoutePath(method, controllerName)
	actionMap := fmt.Sprintf("%s%s%s", controllerName, actionControllerSeparator, method.Name)

	if Debug {
		fmt.Println("route", routePath)
	}

	checkError := (func() error {
		var errorInvalidActionSchema error = errors.New(fmt.Sprintf("Invalid action parameter.\n      Action %s should be written in:\n      func (d *%s) %s (w http.ResponseWriter, r *http.Request)", actionMap, controllerName, method.Name))

		if methodBody.Type().NumIn() != 2 {
			return errorInvalidActionSchema
		} else if methodBody.Type().In(0).String() != "http.ResponseWriter" {
			return errorInvalidActionSchema
		} else if methodBody.Type().In(1).String() != "*http.Request" {
			return errorInvalidActionSchema
		}

		return nil
	})

	if err := checkError(); err != nil {
		fmt.Println("error", err.Error())
		os.Exit(0)
	}

	methodHandler := methodBody.Interface().(func(w http.ResponseWriter, r *http.Request))
	http.HandleFunc(routePath, methodHandler)

	if Debug {
		fmt.Println("   ->", actionMap)
	}

	if HiddenIndex && method.Name == fmt.Sprintf("%s%s%s", actionPrefix, actionSeparator, actionIndexName) {
		routePathWithoutIndex := strings.Join(strings.Split(routePath, actionSlash)[:2], actionSlash)
		http.HandleFunc(routePathWithoutIndex, methodHandler)
		if Debug {
			fmt.Println("route", routePathWithoutIndex)
			fmt.Println("   ->", actionMap)
		}
	}

	if actionMap == DefaultAction {
		http.HandleFunc(actionSlash, methodHandler)

		if Debug {
			fmt.Println("route", actionSlash)
			fmt.Println("   ->", actionMap)
		}
	}
}
