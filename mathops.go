package mathops

// Operation defines the interface for mathematical operations
type Operation interface {
	Apply(float64, float64) float64
}

// Add represents the addition operation
type Add struct{}

// Apply performs addition
func (Add) Apply(x, y float64) float64 {
	return x + y
}

// Subtract represents the subtraction operation
type Subtract struct{}

// Apply performs subtraction
func (Subtract) Apply(x, y float64) float64 {
	return x - y
}

// Multiply represents the multiplication operation
type Multiply struct{}

// Apply performs multiplication
func (Multiply) Apply(x, y float64) float64 {
	return x * y
}

// Divide represents the division operation
type Divide struct{}

// Apply performs division
func (Divide) Apply(x, y float64) float64 {
	if y == 0 {
		return 0 // handle division by zero
	}
	return x / y
}

// Calculate takes an operation and two operands, performs the operation, and sends the result through a channel
func Calculate(op Operation, x, y float64, result chan<- float64) {
	result <- op.Apply(x, y)
	close(result)
}
