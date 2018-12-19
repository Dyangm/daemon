# daemon
后台守护进程

## 运用

``` 
--- main.go
package main

import (
	"fmt"
	"github.com/Dyangm/daemon"
	"time"
)

func main() {
	daemon.Daemon()
	for {
		time.Sleep(1 * time.Minute)
		fmt.Println("hello world")
	}
}

```

执行
$go build main.go
```$xslt

main
```

部署
```
./main -d

```

