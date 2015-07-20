<img src="http://oi59.tinypic.com/2v29m5g.jpg" align="right" alt="Godong logo">

Godong
======
Easiest dynamic http route handler for golang.

Introduction
======
Godong is dynamic route handler for golang. Godong will automatically create route and the handler based on the method of registered struct.

Instalation
======
Use `go get` to download the library.

```
go get github.com/novalagung/godong
```

Simple example
======
Godong is very easy to use. First you need to prepare a `struct`. Then create methods with parameters are same as `http.handleFunc` callback function. The name of the method must start with `Action_`.

Method's name will become route name, and method's body will applied as callback of `http.handleFunc`.

```go
package controller

import (
    "fmt"
    "net/http"
)

type Dashboard struct{}

func (d *Dashboard) Action_Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "/dashboard/index")
}

func (d *Dashboard) Action_AboutUs(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "/dashboard/about-us")
}
```

On `main` package, include the package of our `struct`. Also include the `godong` library. Then create empty object variabel using `struct Dashboard`. Call `godong.Route`, use the variabel we just created as parameter. Pass it by reference.

```go
package main

import (
    "github.com/novalagung/godong"
    "github.com/novalagung/test/controller"
    "net/http"
)

func main() {
    godong.Route(&controller.Dashboard{})

    http.ListenAndServe(":3000", nil)
}
```

Godong only cover routes declaration. Thats make `http.ListenAndServer` need to be called manually. This make godong special, because of this we can also register another route like usual.

Documentation
======

Godong will register all methods of applied struct as route on `http.HandleRoute`.

Basically how to use godong is just by creating new instance of the controller's struct, then pass it by reference through `godong.Route`.

The generated route name will be `/controller/action`. 

On the routing process, godong will do several things over action method before it applied as route.

* Add *slash* character before controller name
* Separate controller and action name using *slash*
* Remove *Action_* prefix
* Convert *underscore* of action name to *slash*
* Separate capitalized action name using *dash*
* Then convert the action name to lower case

Example:

```
Controller action : Dashboard.Action_AboutUs
generated route   : /dashboard/about-us

Controller action : DataAnalytic.Action_Data_GetData
generated route   : /data-analytic/data/get-data
```

There are some other things which is configurable.

### Show applied routes 

Set value `DebugMode` to `true`, so whenever `.godong.Route` called, all actions which are successfully registered will printed. Also detailed error message will shown (if there is an error) by enabling this configuration.

```go
godong.Debug = true
godong.Route(&controller.Dashboard{})
```

Below is the sample output.

![Debug mode enabled](http://oi61.tinypic.com/4ut107.jpg)

### Set default action

All routes will defined on schema `/controller/action`. To define the `/` route, fill the `godong.DefaultAction` using action name of picked controller. Please see example below.

```go
godong.Debug = true
godong.DefaultAction = "Dashboard.Action_Index"
godong.Route(&controller.Dashboard{})
godong.Route(&controller.Analytic{})
```

`Dashboard.Action_Index` mean that we will use `Action_Index` method of struct `Dashboard` as `/` route.

![Default action enabled](http://oi60.tinypic.com/a5dtgh.jpg)

### Enable hidden index route

Godong will route index action as `/controller/index`. It's possible to make route `/controller` only for index action by change the value of `godong.HiddenIndex` to `true`. 2 routes will be registered using same handler: `/controller/index` and `/controller`.

```go
godong.Debug = true
godong.HiddenIndex = true
godong.Route(&controller.Dashboard{})
godong.Route(&controller.Analytic{})
```

If you enabled debug mode, the `/controller` will be displayed.

![Hidden index enabled](http://oi60.tinypic.com/23mpag.jpg)

### Change url mode to Capitalized

By default godong will make some changes on the action name, doing some conversion on it (you may have read the explanation on the opening of Documentation section).

The default *dashed mode* url can be changed to *capitalized mode*. Below is the comparison of dashed and capitalized mode url.

```
controller action : Dashboard.Action_Index
dashed mode       : /dashboard/index
capitalized mode  : /Dashboard/Index

controller action : DataAnalytic.Action_Data_GetData
dashed mode       : /data-analytic/data/get-data
capitalized mode  : /DataAnalytic/Data/GetData
```

How to do it, just by change the `godong.UrlMode`'s value to `godong.UrlModeCapitalized`.

```go
godong.Debug = true
godong.UrlMode = godong.UrlModeCapitalized
godong.Route(&controller.Dashboard{})
```

The default value of `godong.UrlMode` is `godong.UrlModeDashed`.

![Capitalized mode url](http://oi59.tinypic.com/qozqd1.jpg)

API Reference
======
### godong.Route()

Apply action method registered struct as route `http.HandleRoute`. Require one parameter, instance of the struct, need to be passed by reference.

Example:

```go
godong.Route(&controller.Dashboard{})
```

### godong.Debug

Show all route map on the command line if set to true. 

This property is boolean type. Default value is `false`.

Example:

```go
godong.Debug = true
```

### godong.DefaultAction

Set `/` url using registered action.

This property is boolean type. Default value is `""`.

Example:

```go
godong.DefaultAction = "Dashboard.Action_Index"
```

### godong.HiddenIndex

Enable `/controller` route which is represent the `/controller/index` route.

This property is boolean type. Default value is `false`.

Example:

```go
godong.HiddenIndex = true
```

### godong.UrlMode

Pick the url mode, is it dashed or capitalized.

This property is enum type. Availables option are:

* godong.UrlModeDashed
* godong.UrlModeCapitalized

Default value is `godong.UrlModeDashed`.

Example:

```go
godong.UrlMode = UrlModeCapitalized
```

Release Notes
======
You may see complete release notes on [RELEASE.md](https://github.com/novalagung/godong/blob/master/RELEASE.md).

Contribution
======
Feel free to add contribution to this project by fork -> pull request.

License
======
The MIT License (MIT)
