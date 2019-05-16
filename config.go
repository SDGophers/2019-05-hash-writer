package hw

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"fmt"
)

const ErrNoConfOpt = "config option %q not found"
const ErrSyntax = "invalid syntax %d"

type ConfigOption struct {
	capture *regexp.Regexp
	tmpls   []string
}

func (opt *ConfigOption) Write(path string, info os.FileInfo, err error) error {
	return nil
}

type Config struct {
	m Map
}

func (conf *Config) Write(root string, opt string) error {
	v, ok := conf.m.Get(opt)
	if !ok {
		return fmt.Errorf(ErrNoConfOpt, opt)
	}

	vv := v.(*ConfigOption)

	return filepath.Walk(root, vv.Write)
}

func ParseConfig(r io.Reader) (*Config, error) {
	conf := &Config{}

	conf.m = NewMap()

	scanner := bufio.NewScanner(r)
	for line := 1; scanner.Scan(); line++ {
		if scanner.Text() == "" {
			continue
		}

		optionName := scanner.Text()
		if !scanner.Scan() {
			return nil, fmt.Errorf(ErrSyntax, line)
		} else {
			line++
		}

		option := &ConfigOption{}
		var err error
		option.capture, err = regexp.Compile(scanner.Text())
		if err != nil {
			return nil, err
		}

		if !scanner.Scan() {
			return nil, fmt.Errorf(ErrSyntax, line)
		} else {
			line++
		}

		option.tmpls = []string{scanner.Text()}

		for line++; scanner.Scan(); line++ {
			if scanner.Text() == "" {
				break
			}
			option.tmpls = append(option.tmpls, scanner.Text())
		}

		conf.m.Set(optionName, option)
	}

	return conf, nil
}
