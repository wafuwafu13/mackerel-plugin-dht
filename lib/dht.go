package mpdht

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"

	"go.bug.st/serial"
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
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	port, err := serial.Open("/dev/cu.usbmodem1101", mode)
	if err != nil {
		return nil, fmt.Errorf("Failed to open port: %s", err)
	}
	scanner := bufio.NewScanner(port)
	scanner.Split(bufio.ScanLines)
	var degrees []string
	for scanner.Scan() {
		degrees = append(degrees, scanner.Text())
		if len(degrees) == 2 {
			break
		}
	}
	degree, _ := strconv.ParseFloat(degrees[1], 64)
	return map[string]float64{"degrees": degree}, nil
}

// Do the plugin
func Do() {
	u := DhtPlugin{}
	helper := mp.NewMackerelPlugin(u)
	helper.Run()
}
