package smodule // import "github.com/sapplications/sbuilder/src/smodule"

Package smod manages smart module

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com

CONSTANTS

const (
	MainItemName string = "main"
	// notifications
	ModuleIsCreatedF string = "%s file has been created\n"
	// errors
	ItemExistsF             string = "the %s item already exists in %s module"
	ItemIsMissingF          string = "the %s item does not exist"
	ModuleFilesMissingF     string = "no sb files in %s"
	ModuleLanguageMismatchF string = "the %s language of %s module is mismatch the %s selected language"
)

FUNCTIONS

func GetModuleFileName(name string) string
func IsItemExists(lang, item string) (bool, string)
func IsModuleExists(name string) bool

TYPES

type Formatter struct {
}

func (f *Formatter) Item(name string, deps map[string]string) string

func (f *Formatter) String(module Reader) string

type Item = map[string]string

type Items = map[string]Item

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

type Manager struct {
	Lang   func() string
	Logger Logger
}

func (m *Manager) AddDependency(item, dependency, resolver string, update bool) error

func (m *Manager) AddItem(module, item string) error

func (m *Manager) DeleteDependency(item, dependency string) error

func (m *Manager) DeleteItem(item string) error

func (m *Manager) Init(module, lang string) error

func (m *Manager) ReadAll(lang string) (Reader, error)

func (m *Manager) SetLogger(logger Logger)

type Reader interface {
	Lang() string
	Items() map[string]map[string]string
	Dependency(string, string) string
	Main() (map[string]string, error)
}

