package backends

import (
	"github.com/Graylog2/nxlog-sidecar/api/graylog"
	"github.com/Sirupsen/logrus"
)

type Backend interface {
	RenderOnChange(graylog.ResponseCollectorConfiguration) bool
}

type Creator func(string) Backend

type backendFactory struct {
	registry map[string]Creator
}

func (bf *backendFactory) register(name string, c Creator) error {
	if _, ok := bf.registry[name]; ok {
		logrus.Error("Collector backend named " + name + " is already registered")
		return nil
	}
	bf.registry[name] = c
	return nil
}

func (bf *backendFactory) get(name string) (Creator, error) {
	c, ok := bf.registry[name]
	if !ok {
		logrus.Fatal("No collector backend named " + name + " is registered")
		return nil, nil
	}
	return c, nil
}

// global registry
var factory = &backendFactory{registry: make(map[string]Creator)}

func RegisterBackend(name string, c Creator) error {
	return factory.register(name, c)
}

func GetBackend(name string) (Creator, error) {
	return factory.get(name)
}