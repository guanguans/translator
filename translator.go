package translator

import (
	"github.com/go-resty/resty/v2"
)

type Translator interface {
	GetName() string
	Translate(q string, from string, to string) (*resty.Response, error)
}

/*************************************************************
 * config default Translator manager
 *************************************************************/

// default Translator driver manager instance
var std = NewManager()

// Register driver to manager instance
func Register(name string, driver Translator) *Manager {
	std.Register(name, driver)
	return std
}

// Unregister an Translator driver
func Unregister(name string) int {
	return std.Unregister(name)
}

// UnregisterAll Translator drivers
func UnregisterAll(fn ...func(Translator Translator)) int {
	return std.UnregisterAll(fn...)
}

// SetDefName set default driver name.
// Deprecated
//  please use DefaultUse() instead it
func SetDefName(driverName string) {
	std.DefaultUse(driverName)
}

// DefaultUse set default driver name
func DefaultUse(driverName string) {
	std.DefaultUse(driverName)
}

// Use driver object by name and set it as default driver.
func Use(driverName string) Translator {
	return std.Use(driverName)
}

// GetTranslator returns a driver instance by name. alias of Driver()
func GetTranslator(driverName string) Translator {
	return std.Translator(driverName)
}

// Driver get a driver instance by name
func Driver(driverName string) Translator {
	return std.Driver(driverName)
}

// Std get default Translator manager instance
func Std() *Manager {
	return std
}

// DefManager get default Translator manager instance
func DefManager() *Manager {
	return std
}

// Default get default Translator driver instance
func Default() Translator {
	return std.Default()
}

/*************************************************************
 * quick use by default Translator driver
 *************************************************************/

func Translate(q string, from string, to string) (*resty.Response, error) {
	return std.Default().Translate(q, from, to)
}
