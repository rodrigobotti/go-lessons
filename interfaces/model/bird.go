package model

// Chicken is a bird of type Chicken
type Chicken interface {
	Cluck() string
}

// Duck is a bird of type Duck
type Duck interface {
	Quack() string
}

// Bird is a generic bird
type Bird struct {
	Name string
}

// Cluck sound made by a Chicken
func (bird Bird) Cluck() string {
	return bird.Name + " says Cluck..."
}

// Quack sound made by a Duck
func (bird Bird) Quack() string {
	return bird.Name + " says Quack..."
}
