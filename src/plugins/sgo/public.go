package sgo

import "github.com/hashicorp/go-plugin"

const (
	AppName string = "sgo"
)

type Plugin struct {
	Builder   Builder
	Generator Generator
	Handshake plugin.HandshakeConfig
	Logger    Logger
}

type Builder interface {
	Init(items map[string]map[string]string)
	Build(appName string) error
	Clean(appName string) error
	SetLogger(logger Logger)
}

type Generator interface {
	Init(items map[string]map[string]string)
	Clean(appName string) error
	Generate(appName string) error
	SetLogger(logger Logger)
}

type Logger interface {
	Trace(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	IsTrace() bool
	IsDebug() bool
	IsInfo() bool
	IsWarn() bool
	IsError() bool
}
