package model

// Property struct holding property data
type Property struct {
	X     int    `json:"coord_x"`
	Y     int    `json:"coord_y"`
	Name  string `json:"name"`
	value float64
}

// SetValue sets property value
func (p *Property) SetValue(value float64) {
	p.value = value
}

// GetValue returns property's value
func (p *Property) GetValue() float64 {
	return p.value
}
