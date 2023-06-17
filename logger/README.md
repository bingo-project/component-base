# Logger
A logger based on [Zap](https://github.com/uber-go/zap)

## Installation
```bash
go get github.com/goer-project/goer-utils
```
## Usage

```go
l := logger.NewChannel(&logger.Channel{
    Path:    "/your-path/log.log",
    Level:   "debug",
    Days:    14,
    Console: true,
    Format: "json",
})

l.Info("This is a log.")
```