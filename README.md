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
	fmt.Fprintf(w, "index")
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
	var dashboard controller.Dashboard
	godong.Route(&dashboard)

	http.ListenAndServe(":3000", nil)
}
```

Godong only cover routes declaration. Thats make `http.ListenAndServer` need to be called manually. This make godong special, because of this we can also register another route like usual.

Documentation
======
Not yet.

Contribution
======
Feel free to add contribution to this project by fork -> pull request.

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
