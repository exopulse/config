// Package config contains helper methods to simplify working with command line arguments.
package config

import (
	"github.com/exopulse/files"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/kballard/go-shellquote"
)

// SplitArgs parses input lines and splits those in tokens. Delimiter characters are space and equals. This method
// is suitable for normalization of command line arguments that originate from config file.
func SplitArgs(args []string) ([]string, error) {
	var s []string

	for _, a := range args {
		tokens, err := shellquote.Split(a)

		if err != nil {
			return nil, err
		}

		for _, token := range tokens {
			idx := strings.Index(token, "=")

			if idx != -1 {
				s = append(s, token[:idx])
				s = append(s, token[idx+1:])
			} else {
				s = append(s, token)
			}
		}
	}

	return s, nil
}

// ReadArgumentsFromConfigFileArg discovers config file argument, reads discovered config file and returns normalized
// arguments.
func ReadArgumentsFromConfigFileArg(longName string, shortName string, defaultConfigFile string) ([]string, error) {
	cf, err := discoverArgumentValue(os.Args[1:], longName, shortName, defaultConfigFile)

	if err != nil {
		return nil, err
	}

	return readNormalizedArgs(cf)
}

// discoverArgumentValue discovers argument value. Argument can be specified using any or both of the long and short
// name variant.
func discoverArgumentValue(args []string, longName string, shortName string, defaultValue string) (string, error) {
	args, err := SplitArgs(args)

	if err != nil {
		return "", err
	}

	v := ""

	shortTarget := "-" + shortName
	longTarget := "--" + longName
	incShort := len(shortName) > 0
	incLong := len(longName) > 0

	for i, p := range args {
		if ((incShort && p == shortTarget) || (incLong && p == longTarget)) && i+1 < len(args) {
			if v != "" {
				return defaultValue, errors.Errorf("duplicate flag value for %s/%s", longName, shortName)
			}

			v = args[i+1]
		}
	}

	if v == "" {
		v = defaultValue
	}

	return v, nil
}

// readNormalizedArgs reads arguments from specified file. Slice returned contains normalized arguments.
func readNormalizedArgs(file string) ([]string, error) {
	if !files.FileExists(file) {
		return nil, nil
	}

	lines, err := files.ReadLines(file)

	if err != nil {
		return nil, err
	}

	args, err := SplitArgs(lines)

	if err != nil {
		return nil, err
	}

	return args, nil
}
