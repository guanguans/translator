package translator

import (
	"github.com/go-resty/resty/v2"
)

// default supported Translator driver name
const (
	DriverBaidu  = "baidu"
	DriverYoudao = "youdao"
)

/*************************************************************
 * Translator Manager
 *************************************************************/

// Manager definition
type Manager struct {
	// default driver name
	defName string
	// drivers map
	drivers map[string]Translator
}

// NewManager create a Translator manager instance
func NewManager() *Manager {
	return &Manager{
		// defName: driverName,
		drivers: make(map[string]Translator, 8),
	}
}

// Register new Translator driver
func (m *Manager) Register(name string, driver Translator) *Manager {
	// always use latest as default driver.
	m.defName = name
	// save driver instance
	m.drivers[name] = driver
	return m
}

// Unregister an Translator driver
func (m *Manager) Unregister(name string) int {
	if _, ok := m.drivers[name]; !ok {
		return 0
	}

	delete(m.drivers, name)

	// reset default driver name.
	if m.defName == name {
		m.defName = ""
	}
	return 1
}

// SetDefName set default driver name. alias of DefaultUse()
// Deprecated
//  please use DefaultUse() instead it
func (m *Manager) SetDefName(driverName string) {
	m.DefaultUse(driverName)
}

// DefaultUse set default driver name
func (m *Manager) DefaultUse(driverName string) {
	if _, ok := m.drivers[driverName]; !ok {
		panic("Translator driver: " + driverName + " is not registered")
	}

	m.defName = driverName
}

// Default returns the default driver instance
func (m *Manager) Default() Translator {
	if c, ok := m.drivers[m.defName]; ok {
		return c
	}

	panic("Translator driver: " + m.defName + " is not registered")
}

// Use driver object by name and set it as default driver.
func (m *Manager) Use(driverName string) Translator {
	m.DefaultUse(driverName)
	return m.Driver(driverName)
}

// Translator get driver by name. alias of Driver()
func (m *Manager) Translator(driverName string) Translator {
	return m.drivers[driverName]
}

// Driver get a driver instance by name
func (m *Manager) Driver(driverName string) Translator {
	return m.drivers[driverName]
}

// DefName get default driver name
func (m *Manager) DefName() string {
	return m.defName
}

// UnregisterAll Translator drivers
func (m *Manager) UnregisterAll(fn ...func(Translator Translator)) int {
	num := len(m.drivers)

	// unregister
	m.defName = ""
	for name, driver := range m.drivers {
		if len(fn) > 0 {
			fn[0](driver)
		}

		delete(m.drivers, name)
	}
	return num
}

/*************************************************************
 * Quick use by default Translator driver
 *************************************************************/

func (m *Manager) Translate(q string, from string, to string) (*resty.Response, error) {
	return m.Default().Translate(q, from, to)
}
