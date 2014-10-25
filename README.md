![Golang Gopher](./images/small-gopher.png) ![GoVend](./images/govend.png)

### What It Is

The command `govend` takes yet another stab at easily integrating dependency management with golang projects. While many different solutions already exist to manage third party golang packages, `govend` tries a more simple approach.

### How It Works

In my (limited) experience, `go get` has been very effective for downloading and adding golang packages into a local development `$GOPATH`. Yet, when using `go get` as a step in production deployments it has not been as nearly effective. In my (again limited) experience, depending on the OS build, network environment, and hosting provider `go get` might fail.

`govend` solves this problem by pulling golang dependencies into your project repo. By creating a `deps.json` file indicate your `vendor` directory and list your `go get` dependencies, running `govend` will copy your third party packages into your repository.

`govend` achieve this by being a very simple wrapper around `go get`.

### Example

> First create a `deps.json` file at the root of your repository.

```javascript
{
    "deps": [
        "github.com/gorilla/mux",
        "gopkg.in/mgo.v2"
    ]
}
```

> Next run the command `govend`.

```bash
$ govend

 ↓ github.com/gorilla/mux
 ↓ gopkg.in/mgo.v2

Vending complete
```

> Lastly change your imports to use the `_vendor` directory.

```go
package example

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"_vendor/github.com/gorilla/mux"
	"_vendor/gopkg.in/mgo.v2"
)
```
