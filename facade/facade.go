package facade

import "fmt"

/*
Facade -> API This is facade interface of the module. Most codes use this interface use this interface to simplify access
to the class.

Facade patterns simplify complex systems by providing a unified interface,
making it easier for clients to interact with subsystems. They enhance code
readability and reduce dependencies, promoting better maintainability and
scalability in software design.
*/

// Unified Maid interface
type API interface {
	Test() string
}

// Adding two Modules for the systems
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

// Adding interface for Amodule
type AModuleAPI interface {
	TestA() string
}

// Adding interface for Bmodule
type BModuleAPI interface {
	TestB() string
}

// Calling aModuleImple
func newAModuleAPI() AModuleAPI {
	return &aModuleImple{}
}

// Calling bModuleImple
func newBModuleAPI() BModuleAPI {
	return &bModuleImple{}
}

// Creating New instance for API handling both Module
func NewAPI() API {
	return &apiImpl{
		a: newAModuleAPI(),
		b: newBModuleAPI(),
	}
}

// a module with empty type
type aModuleImple struct{}

// b module with empty type
type bModuleImple struct{}

// implemented One Test Method from Two Module
func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// Usage of aModuleImple
func (*aModuleImple) TestA() string {
	return "A module running"
}

// Usage of bModuleImple
func (*bModuleImple) TestB() string {
	return "B module running"
}
