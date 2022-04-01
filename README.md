# namegen

`namegen` is Ionburst Cloud's adaption of the name generation [functionality](https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go) used by Docker to name containers. We've then added our own Scottish twist, using the names of famous Scottish inventors, scientists and engineers.

## Using namegen

Our initial release will cover a Go package and accompanying CLI tool, with plans to support the following languages:
- .NET
- Javascript/Node.js
- Python

### CLI usage

```
[example@namegen-example ~]$ namegen
powerful_watt
```

### Go usage

```go
package main

import (
	"fmt"

	"gitlab.com/ionburst/namegen/namegen"
)

func main() {
	fmt.Println(namegen.GetRandomName())
}
```

## Contributing to namegen

If you have any suggestions for additional adjectives, names of famous Scottish inventors, scientists, or engineers for `namegen`, feel free to open an issue on [GitHub](https://github.com/ionburstcloud/namegen/issues) or [GitLab](https://gitlab.com/ionburst/namegen/-/issues).