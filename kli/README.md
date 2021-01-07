
# A Simple CLI library. Dependency free.

### Features

  * Nested Subcommands
  * Uses the standard library `flag` package
  * Auto-generated help
  * Custom banners
  * Hidden Subcommands
  * Default Subcommand
  * Dependency free

### Example

```go
package main

import (
   "fmt"
   "log"

	"github.com/openware/pkg/kli"
)

func main() {

	// Create new cli
	cli := kli.NewCli("Flags", "A simple example", "v0.0.1")

	// Name
	name := "Anonymous"
	cli.StringFlag("name", "Your name", &name)

	// Define action for the command
	cli.Action(func() error {
		fmt.Printf("Hello %s!\n", name)
		return nil
	})

	if err := cli.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}

}
```

#### Generated Help

```shell
$ flags --help
Flags v0.0.1 - A simple example

Flags:

  -help
        Get help on the 'flags' command.
  -name string
        Your name
```
