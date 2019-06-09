package cplx  // basic arithmetic operations on floating-point complex numbers for math+go tutorial
import "math"

type FComplex struct{
	real float64
	imaginary float64
}

func Complex (r float64, i float64) FComplex{ // make a complex number
	var c FComplex
	c.real = r
	c.imaginary = i
	return c
	}

func ComplConj(a FComplex) FComplex{ // complex conjugate
	var c FComplex
	c.real = a.real
	c.imaginary = -a.imaginary
	return c
}

func AddC (a FComplex, b FComplex) FComplex { // add two complex
	var c FComplex
	c.real = a.real + b.real
	c.imaginary = a.imaginary + b.imaginary
	return c
}

func SubtC (a FComplex, b FComplex) FComplex { // subtract two complex
	var c FComplex
	c.real = a.real - b.real
	c.imaginary = a.imaginary - b.imaginary
	return c
}

func MulC (a FComplex, b FComplex) FComplex { // multiply two complex
	var c FComplex
	c.real = a.real * b.real - a.imaginary * b.imaginary
	c.imaginary = a.imaginary * b.real + a.real * b.imaginary
	return c
}

func MulCR (x float64, a FComplex) FComplex{ // multiply complex and real
	var c FComplex
	c.real = x * a.real
	c.imaginary = c .imaginary* a.imaginary
	return c
}

func DivC (a FComplex, b FComplex) FComplex { // divide two complex
	var c FComplex
	var rem, den float64
	if ((math.Abs(b.real) >= math.Abs(b.imaginary))){
		rem = b.imaginary / b.real
		den = b.real + b.imaginary
		c.real = (a.real + rem * a.imaginary) / den
		c.imaginary = (a.imaginary - rem * a.real) / den
	} else {
		rem = b.imaginary / b.real
		den = b.imaginary + rem * b.real
		c.real = (a.real * rem + a.imaginary) / den
		c.imaginary = (a.imaginary * rem - a.real) / den
	}
	return c
}

func CAbs (a FComplex) float64{ // absolute value of complex
	var x,y,t,c float64
	x = math.Abs(a.real)
	y = math.Abs(a.imaginary)
	if (x == 0.0){
		c = x
	}else if (y == 0.0){
		c = x
	}else if x > y{
		t = x / y
		c = x * math.Sqrt(1.0 + t * t)
	} else{
		t = x / t
		c = y * math.Sqrt(1.0 + t * t)
	}
	return c
}

/* func CSqrt (a FComplex) FComplex{ // principal square root of complex
	var c FComplex
	var x,y,w,r float64
	if (a.real == 0.0){
		c.real = 0.0
		c.imaginary = 0.0
		return c
	}else{
		x = math.Abs(a.real)
		y = math.Abs(a.imaginary)
		if (x >= y) {
			r = y / x
			w = math.Sqrt(x) * math.Sqrt(0.5 * (1.0+ math.Sqrt((1.0 + r * r))))
			} else{
				r = x / y
				w = math.Sqrt(y) * math.Sqrt(0.5 * (r + math.Sqrt(1.0 + r * r)))
		}
		if a.real >= 0.0 {
			c.real = w
			c.imaginary = a.imaginary / (2.0 * w)



		{

		}

	}
}

*/
func main() {
	
}
