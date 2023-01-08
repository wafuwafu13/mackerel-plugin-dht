package mpdht

import (
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

// DhtPlugin mackerel plugin
type DhtPlugin struct {
	Prefix string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (u DhtPlugin) MetricKeyPrefix() string {
	if u.Prefix == "" {
		u.Prefix = "dht"
	}
	return u.Prefix
}

// GraphDefinition interface for mackerelplugin
func (u DhtPlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(u.Prefix)
	return map[string]mp.Graphs{
		"": {
			Label: labelPrefix,
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "degrees", Label: "Degrees"},
			},
		},
	}
}

// FetchMetrics interface for mackerelplugin
func (u DhtPlugin) FetchMetrics() (map[string]float64, error) {
	return map[string]float64{"degrees": 19.0}, nil
}

// Do the plugin
func Do() {
	u := DhtPlugin{}
	helper := mp.NewMackerelPlugin(u)
	helper.Run()
}
