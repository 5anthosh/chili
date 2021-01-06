package function

import "math/big"

var limit = pow(new(2), 256)

func pow(a *big.Float, e uint64) *big.Float {
	result := zero().Copy(a)
	for i := uint64(0); i < e-1; i++ {
		result = mul(result, a)
	}
	return result
}

func root(a *big.Float, n uint64) *big.Float {
	mulvalue := new(1)
	if a.Sign() == -1 {
		mulvalue = new(-1)
		a = a.Mul(a, mulvalue)
	}

	n1 := n - 1
	n1f, rn := new(float64(n1)), div(new(1.0), new(float64(n)))
	x, x0 := new(1.0), zero()
	_ = x0
	for {
		potx, t2 := div(new(1.0), x), a
		for b := n1; b > 0; b >>= 1 {
			if b&1 == 1 {
				t2 = mul(t2, potx)
			}
			potx = mul(potx, potx)
		}
		x0, x = x, mul(rn, add(mul(n1f, x), t2))
		if lesser(mul(abs(sub(x, x0)), limit), x) {
			break
		}
	}
	x = x.Mul(x, mulvalue)
	return x
}

func abs(a *big.Float) *big.Float {
	return zero().Abs(a)
}

func new(f float64) *big.Float {
	r := big.NewFloat(f)
	r.SetPrec(256)
	return r
}

func div(a, b *big.Float) *big.Float {
	return zero().Quo(a, b)
}

func zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

func mul(a, b *big.Float) *big.Float {
	return zero().Mul(a, b)
}

func add(a, b *big.Float) *big.Float {
	return zero().Add(a, b)
}

func sub(a, b *big.Float) *big.Float {
	return zero().Sub(a, b)
}

func lesser(x, y *big.Float) bool {
	return x.Cmp(y) == -1
}
