[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/github/license/hetiansu5/adapter)](LICENSE)

## Introduction
A cache adapter based on GoFrame master-slave redis client

## Keywords
cache adapter redis GoFrame


## Quick Start

```golang
package main

import (
	"github.com/hetiansu5/msredis"
	"github.com/hetiansu5/adapter"
)

masterName := "cache" // master redis config name
slaveName := "cache_slave" // slave redis config name, can be more
group := msredis.Group(masterName, slaveName)
cache := adapter.NewRedis(msredis.group)
cache.Get(ctx, key)
```


### License
MIT
