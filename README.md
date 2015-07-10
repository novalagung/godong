Introduction
======
Godong is dynamic route handler for golang, so you don't need to register it manually. Godong will automatically create route handler based on the method inside of struct.

Instalation
======
Use `go get` to download the library.

```
go get github.com/novalagung/godong
```

Simple example
======
Godong is very easy to use. First you need to prepare a file which contains a struct. The struct will become the controller name.

And then create a method with parameters is same as `http.handleFunc` callback function. The name of the method must start with `Action_`.

```go
package controller

import (
	"fmt"
	"net/http"
)

type Dashboard struct{}

func (d *Dashboard) Action_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index")
}
```

After that on main package, include the file that we've ben created just now. Also include the `godong` library. Then create empty variabel which type is our controller. And apply `godong.Route` function to it.

Basically we have completelly succeed registering all method inside struct `Dashboard` as an route. The method name will become the route name.

```go
package main

import (
	"github.com/novalagung/godong"
	"github.com/novalagung/test/controller"
	"net/http"
)

func main() {
	var dashboard controller.Dashboard
	godong.Route(&dashboard)

	http.ListenAndServe(":3000", nil)
}
```

Godong only cover the declaration. So you need to call `http.ListenAndServer`. This make godong special, you also can register custom route (like usual) which is not included on the controller.

Documentation
======
not yet

Contribution
======
Feel free to add contribution to this project by fork -> pull request

License
======
The MIT License (MIT)

Copyright (c) 2015 Noval Agung Prayogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
