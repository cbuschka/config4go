# config4go [![Build Status](https://travis-ci.org/cbuschka/config4go.svg)](https://travis-ci.org/cbuschka/config4go) [![Code Coverage](https://img.shields.io/codecov/c/github/cbuschka/config4go.svg)](https://codecov.io/gh/cbuschka/config4go) [![Go Report Card](https://goreportcard.com/badge/github.com/cbuschka/config4go)](https://goreportcard.com/report/github.com/cbuschka/config4go) [![GoDoc](https://godoc.org/github.com/cbuschka/config4go?status.svg)](https://godoc.org/github.com/cbuschka/config4go) [![Release](https://img.shields.io/github/release/cbuschka/config4go.svg)](https://github.com/cbuschka/config4go/releases/latest) [![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)

A simple config loader written in go.

## Features

* Unicode support (uses go runes, not bytes)
* Very efficient (implementation based on a state machine)

## Status

Still a prototype.

## Examples

### Config format:
```
# a comment
key1 = value1

key2=value2 # another comment

#eof
```

### Load config as map:
```
import "github.com/cbuschka/config4go"

func doIt() (map[string]string, error) {
    configReader := NewConfigReader()
	config, err := configReader.ReadConfigFromFile("example.conf")
	if err != nil {
	    return nil, err
	}
	return config.ToMap(), nil
}
```

## Contributing

Config4Go is an open source project, and contributions are welcome! Feel free to raise an issue or submit a pull request.

## License

Copyright (c) 2016 by [Cornelius Buschka](https://github.com/cbuschka).

[MIT License](LICENSE)
