# golog
A light weight logger for GoLang

- Wrap the buildin log module
- Rotate log file by time (day, hour, close)
- Support all levels and string format
- Thread safety

### Install

```sh
go get github.com/panjiang/golog
```

#### Example
See [source code](https://github.com/panjiang/golog/tree/master/example)

```golang
import (
	log "github.com/panjiang/golog"
)

func main() {
	// Config the logger
	// If no config, output to stdout
	//
	// Also accept json config
	// {
	// 	 "file": "./logs/test.log",
	// 	 "level": "debug",
	// 	 "rotate": "day"
	// }
	log.ParseConfig(&log.Config{
		File:   "./logs/test.log", // log filename template, empty will only to stdout
		Level:  "debug",           // output level: debug, info, warn, error
		Rotate: "day",             // rotate by: day, hour, close
	})

	log.Println("info message", 0)
	log.Printf("info message %d", 0)

	log.Debug("debug message", 1)
	log.Debugf("debug message %d", 1)

	log.Info("info message", 2)
	log.Infof("info message %d", 2)

	log.Warn("warn message", 3)
	log.Warnf("warn message %d", 3)

	log.Error("error message", 4)
	log.Errorf("error message %d", 4)

	log.Panic("panic message", 5)
	log.Panicf("panic message %d", 5)

	log.Fatal("fatal message", 6)
	log.Fatalf("fatal message %d", 6)
}
```

#### Output

- Auto generate a log file `./logs/test_20180703.log`

``` shell
2018/07/03 16:11:29 [info] info message 0
2018/07/03 16:11:29 [info] info message 0
2018/07/03 16:11:29 [debug] debug message 1
2018/07/03 16:11:29 [debug] debug message 1
2018/07/03 16:11:29 [info] info message 2
2018/07/03 16:11:29 [info] info message 2
2018/07/03 16:11:29 [warn] warn message 3
2018/07/03 16:11:29 [warn] warn message 3
2018/07/03 16:11:29 [error] error message 4
2018/07/03 16:11:29 [error] error message 4
2018/07/03 16:11:29 [error] panic message 5
```
