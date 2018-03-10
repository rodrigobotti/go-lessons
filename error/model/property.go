package model

import (
	"errors"
)

// Property struct holding property data
type Property struct {
	X     int    `json:"coord_x"`
	Y     int    `json:"coord_y"`
	Name  string `json:"name"`
	value float64
}

var (
	// ErrLowPropertyValue when property value is too low
	ErrLowPropertyValue = errors.New("Property value is too low")
	// ErrHighPropertyValue when property value is too high
	ErrHighPropertyValue = errors.New("Property value is too damn high")
)

// SetValue sets property value
func (p *Property) SetValue(value float64) error {
	if value < 50000 {
		return ErrLowPropertyValue
	} else if value >= 10000000 {
		return ErrHighPropertyValue
	}
	p.value = value
	return nil
}

// GetValue returns property's value
func (p *Property) GetValue() float64 {
	return p.value
}
