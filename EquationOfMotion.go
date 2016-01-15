package intersect

// Basic x = mt + c
// To get the position x at any time t, resolve equation above
type EquationOfMotion struct {
	c float64 // constant for the line - not equivalent to position along this axis, equivalent to position at time 0
	m float64 // gradient of the line - equivalent to velocity along the axis
}

func (equation *EquationOfMotion) AlterVelocity(Δ float64) { // UNICODE! (It's a capital delta character)
	equation.m += Δ
}

func (equation *EquationOfMotion) GetPositionAtTime(t float64) float64 {
	m, c := equation.m, equation.c
	return m*t + c
}
