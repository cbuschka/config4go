# config4go [![Build Status](https://travis-ci.org/cbuschka/config4go.svg)](https://travis-ci.org/cbuschka/config4go) [![Go Report Card](https://goreportcard.com/badge/github.com/cbuschka/config4go)](https://goreportcard.com/report/github.com/cbuschka/config4go) [![GoDoc](https://godoc.org/github.com/cbuschka/config4go?status.svg)](https://godoc.org/github.com/cbuschka/config4go) [![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)

A simple config loader written in go.

## Status

Still a prototype.

## Example

### Load config as map:
```
import "github.com/cbuschka/config4go"

func doIt() (map[string]string, error) {
    configReader := NewConfigReader()
	config, err := configReader.ReadConfigFromFile("example.conf")
	if err != nil {
	    return err
	}
	return config.ToMap()
}
```


## License

Copyright (c) 2016 by [Cornelius Buschka](https://github.com/cbuschka).

[MIT License](LICENSE)
