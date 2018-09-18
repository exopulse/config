# exopulse config package
Golang package config contains helper methods to simplify working with command line arguments.

[![CircleCI](https://circleci.com/gh/exopulse/config.svg?style=svg)](https://circleci.com/gh/exopulse/config)
[![Build Status](https://travis-ci.org/exopulse/config.svg?branch=master)](https://travis-ci.org/exopulse/config)
[![GitHub license](https://img.shields.io/github/license/exopulse/config.svg)](https://github.com/exopulse/config/blob/master/LICENSE)

# Overview

This package contains helper methods to simplify working with command line arguments.

## Features

Use method _config.ReadArgumentsFromConfigFileArg_ to find config file argument and read configuration switched from the
file this argument points to. Arguments in config file can be split over multiple lines.

# Using config package

## Installing package

Use go get to install the latest version of the library.

    $ go get github.com/exopulse/config
 
Include config in your application.
```go
import "github.com/exopulse/config"
```
## Methods

    // ReadArgumentsFromConfigFileArg discovers config file argument, reads discovered config file and returns normalized
    // arguments.
    func ReadArgumentsFromConfigFileArg(longName string, shortName string, defaultConfigFile string) ([]string, error) {}

## Example

```go
args, err := config.ReadArgumentsFromConfigFileArg("c", "--config-file", "")
```    
    $ mytool --work-dir /tmp --config-file ~/.mytool.cfg

This will read config arguments from ~/.mytool.cfg and append those to arguments already found on config line.

# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

Config package is released under the MIT license. See
[LICENSE](https://github.com/exopulse/config/blob/master/LICENSE)
