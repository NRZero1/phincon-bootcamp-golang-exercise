package model

type Calculator struct {
	val1, val2 float64
}

func (c *Calculator) Addition() float64 {
	return c.GetVal1() + c.GetVal2()
}

func (c *Calculator) Substraction() float64 {
	return c.GetVal1() - c.GetVal2()
}

func (c *Calculator) Division() float64 {
	if c.GetVal2() == 0 {
		panic("Division by 0")
	}
	return c.GetVal1() / c.GetVal2()
}

func (c *Calculator) Multiplication() float64 {
	return c.GetVal1() * c.GetVal2()
}

func (c *Calculator) SetVal1(val1 float64) {
	c.val1 = val1
}

func (c *Calculator) SetVal2(val2 float64) {
	c.val2 = val2
}

func (c *Calculator) GetVal1() float64 {
	return c.val1
}

func (c *Calculator) GetVal2() float64 {
	return c.val2
}

func (c *Calculator) GetResultIntFloat(calculateResult float64) (int64, float64) {
	return int64(calculateResult), calculateResult
}

func (c *Calculator) Calculate(menu int) (int64, float64) {
	var result float64

	switch menu {
	case 1:
		result = c.Addition()
	case 2:
		result = c.Substraction()
	case 3:
		result = c.Division()
	case 4:
		result = c.Multiplication()
	}
	return c.GetResultIntFloat(result)
}