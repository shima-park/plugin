package plugin

import (
	"flag"
	"fmt"
	"strings"
)

var _ flag.Value = &PluginList{}

type PluginList struct {
	Paths []string
}

func (p *PluginList) String() string {
	return strings.Join(p.Paths, ",")
}

func (p *PluginList) Set(v string) error {
	for _, path := range p.Paths {
		if path == v {
			return fmt.Errorf("%s is already a registered plugin", path)
		}
	}
	p.Paths = append(p.Paths, v)
	return nil
}
