# uber-bo/zap logger for fasthttp and glog ....



## zap log to console for debug make easy

```
package main

import (
	"go.uber.org/zap"

	"github.com/tsingson/zaplogger"
)

func main() {
	core := zaplogger.NewConsoleDebug()

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("constructed a logger")
}

```

output

```
2019-08-01T20:59:41.879+0800    INFO    constructed a logger

```


add  stack trace

```
package main

import (
	"go.uber.org/zap"

	"github.com/tsingson/zaplogger"
)

func main() {
	core := zaplogger.NewConsoleDebug()

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core).WithOptions(zap.AddCaller())
	defer logger.Sync()
	logger.Info("constructed a logger")
}
```

output
```
2019-08-01T21:23:20.549+0800    INFO    example/main.go:15  constructed a logger

```