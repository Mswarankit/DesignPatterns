package adapter

/*
Adapter -> The adapter pattern is used to convert one interface to another interface.

In actual use, Adapter patterns allow incompatible interfaces to work together by converting one interface into another,
enabling integration of diverse systems. They promote code reusability and
flexibility, making it easier to modify or extend functionality without
altering existing code.
*/

// Target is the target interface to be adapted
type Target interface {
	Request() string
}

// Adaptee is the target interface too be adapted
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee is the factory function of the adapted interface
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// AdapteeImpl is the target class to be adopted
type adapteeImpl struct{}

// Specific Request is a method of the target class
func (*adapteeImpl) SpecificRequest() string {
	return "adapte method"
}

// NewAdapter is the factory function of Adapter
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// Adapter is an adapter that converts Adaptee to target interface
type adapter struct {
	Adaptee
}

// Request implements the Target interface
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
