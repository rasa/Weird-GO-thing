package main
import "fmt"
import "math/big"
func main() {
	const d = 40  // number of decimal places	
	const p = d * 4 // bits of precision
	const w = 1 + d / 2  // width of integers a & b
	a := new(big.Float).SetPrec(p).SetInt64(0)
	b := new(big.Float).SetPrec(p).SetInt64(1)
	q := new(big.Float).SetPrec(p) // quotient (b / a)
	l := new(big.Float).SetPrec(p) // last quotient	
	s := new(big.Float).SetPrec(p) // difference between quotients (q - l)
	z := new(big.Float).SetPrec(p).SetInt64(0) // 0.0
	for i := 1; i <= 1000; i++ {
		a.Add(a, b) // a = a + b
		a, b = b, a // a = b, b = a		
		q.Quo(b, a) // q = b / a		
                fmt.Printf("%3d %*s / %*s = %*s\n", i, w, b.Text('f', 0), w, a.Text('f', 0), d + 2, q.Text('f', d))
		s.Sub(q, l) // s = q - l
		if s.Text('f', d) == z.Text('f', d) { // break if q == r	
			break;			
		}
		l.Set(q) // l = q
	}
}
