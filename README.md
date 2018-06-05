# noip

`noip` implements the dynamic update API of no-ip in go. To learn about the API please refere to https://www.noip.com/integrate/request.

## Usage

Fetch the library via the go tool:

```
# go get -u github.com/unprofession-al/noip
```

Use the lib in your program: 

```
package main

import (
    ...

    "github.com/unprofession-al/noip"
)

func main() {
    ...
    c := noip.New(user, pass, hostname, myip, useragent)
    c.Run(60, false)
    ...
}
```
