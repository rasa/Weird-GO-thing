package main
import "fmt"
import "math/big"
func main() {
	var d int = 40 // number of decimal places	
	p := uint(d * 4) // bits of precision
	w := uint(1 + d / 2) // width of integers a & b
	a := new(big.Float).SetPrec(p).SetInt64(0)
	b := new(big.Float).SetPrec(p).SetInt64(1)
	q := new(big.Float).SetPrec(p) // quotient (b / a)
	l := new(big.Float).SetPrec(p) // last quotient	
	s := new(big.Float).SetPrec(p) // difference between quotients (q - l)
	e := new(big.Int).SetInt64(int64(d))
	t := new(big.Int).SetInt64(10)
	t.Exp(t, e, nil)
	y := new(big.Float).SetPrec(p).SetInt64(1)
	z := new(big.Float).SetPrec(p).SetInt(t)
	z.Quo(y, z)
	for i := 1; i <= 2000; i++ {
		a.Add(a, b) // a = a + b
		a, b = b, a // a = b, b = a		
		q.Quo(b, a) // q = b / a		
                fmt.Printf("%3d %*s / %*s = %*s\n", i, w, b.Text('f', 0), w, a.Text('f', 0), d + 2, q.Text('f', d))
		s.Abs(s.Sub(q, l)) // s = abs(q - l)
		if s.Cmp(z) <= 0 {
			break
		}
		l.Set(q) // l = q
	}
}
