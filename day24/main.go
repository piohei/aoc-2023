package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
    "math/big"
)

//go:embed input
var input string

const (
	debug = true

	MAX_INT = 2147483647
)

var (
	// MIN_X = &Frac{big.NewInt(int64(7)), big.NewInt(int64(1))}
	// MAX_X = &Frac{big.NewInt(int64(27)), big.NewInt(int64(1))}
	// MIN_Y = &Frac{big.NewInt(int64(7)), big.NewInt(int64(1))}
	// MAX_Y = &Frac{big.NewInt(int64(27)), big.NewInt(int64(1))}
	// MIN_Z = &Frac{big.NewInt(int64(7)), big.NewInt(int64(1))}
	// MAX_Z = &Frac{big.NewInt(int64(27)), big.NewInt(int64(1))}

	MIN_X = &Frac{big.NewInt(int64(200000000000000)), big.NewInt(int64(1))}
	MAX_X = &Frac{big.NewInt(int64(400000000000000)), big.NewInt(int64(1))}
	MIN_Y = &Frac{big.NewInt(int64(200000000000000)), big.NewInt(int64(1))}
	MAX_Y = &Frac{big.NewInt(int64(400000000000000)), big.NewInt(int64(1))}
	MIN_Z = &Frac{big.NewInt(int64(200000000000000)), big.NewInt(int64(1))}
	MAX_Z = &Frac{big.NewInt(int64(400000000000000)), big.NewInt(int64(1))}
)

type Frac struct {
	nom, den *big.Int
}

func (f *Frac) String() string {
	n, _ := f.nom.Float64()
	d, _ := f.den.Float64()
	return fmt.Sprintf("%.3f", n / d)
	// return fmt.Sprintf("%v / %v", f.nom, f.den)
}

func (f1 *Frac) Equals(f2 *Frac) bool {
	a := big.NewInt(int64(0))
	a = a.Mul(f1.nom, f2.den)
	b := big.NewInt(int64(0))
	b = b.Mul(f2.nom, f1.den)
	return a.Cmp(b) == 0
}

func (f1 *Frac) Add(f2 *Frac) *Frac {
	// (a1b2 - a2b1) / b1b1
	nomL := big.NewInt(int64(0)).Mul(f1.nom, f2.den)
	nomR := big.NewInt(int64(0)).Mul(f2.nom, f1.den)
	nom := big.NewInt(int64(0)).Add(nomL, nomR)
	den := big.NewInt(int64(0)).Mul(f1.den, f2.den)

	// if debug {
	// 	fmt.Printf("%v + %v = %v / %v\n", f1, f2, nom, den)
	// }

	return &Frac{
		nom: nom, den: den,
	}
}

func (f1 *Frac) Sub(f2 *Frac) *Frac {
	// (a1b2 - a2b1) / b1b2
	nomL := big.NewInt(int64(0)).Mul(f1.nom, f2.den)
	nomR := big.NewInt(int64(0)).Mul(f2.nom, f1.den)
	nom := big.NewInt(int64(0)).Sub(nomL, nomR)
	den := big.NewInt(int64(0)).Mul(f1.den, f2.den)

	// if debug {
	// 	fmt.Printf("%v - %v = %v / %v\n", f1, f2, nom, den)
	// }

	return &Frac{
		nom: nom, den: den,
	}
}

func (f1 *Frac) Mul(f2 *Frac) *Frac {
	// a1a2 / b1b2
	nom := big.NewInt(int64(0)).Mul(f1.nom, f2.nom)
	den := big.NewInt(int64(0)).Mul(f1.den, f2.den)

	// if debug {
	// 	fmt.Printf("%v * %v = %v / %v\n", f1, f2, nom, den)
	// }

	return &Frac{
		nom: nom, den: den,
	}
}

func (f1 *Frac) Div(f2 *Frac) *Frac {
	// a1b2 / b1a2
	nom := big.NewInt(int64(0)).Mul(f1.nom, f2.den)
	den := big.NewInt(int64(0)).Mul(f1.den, f2.nom)

	// if debug {
	// 	fmt.Printf("%v / %v = %v / %v\n", f1, f2, nom, den)
	// }

	return &Frac{
		nom: nom, den: den,
	}
}

func (f1 *Frac) Cmp(f2 *Frac) int {
	nomL := big.NewInt(int64(0)).Mul(f1.nom, f2.den)
	nomR := big.NewInt(int64(0)).Mul(f2.nom, f1.den)
	signum := f2.den.Sign() * f1.den.Sign()

	// if debug {
	// 	fmt.Printf("Cmp(%v, %v) -> %v\n", f1, f2, signum * nomL.Cmp(nomR))
	// }

	return signum * nomL.Cmp(nomR)
}

func (f1 *Frac) Sign() int {
	// if debug {
	// 	fmt.Printf("Sig(%v) -> %v\n", f1, f1.nom.Sign() * f1.den.Sign())
	// }

	return f1.nom.Sign() * f1.den.Sign()
}

type Hailstone struct {
	px, py, pz *Frac
	vx, vy, vz *Frac
	a, b *Frac
}

func NewHailstone(pxString, pyString, pzString, vxString, vyString, vzString string) *Hailstone {
	px := big.NewInt(ignoreErr(strconv.ParseInt(pxString, 10, 64)))
	py := big.NewInt(ignoreErr(strconv.ParseInt(pyString, 10, 64)))
	pz := big.NewInt(ignoreErr(strconv.ParseInt(pzString, 10, 64)))
	vx := big.NewInt(ignoreErr(strconv.ParseInt(vxString, 10, 64)))
	vy := big.NewInt(ignoreErr(strconv.ParseInt(vyString, 10, 64)))
	vz := big.NewInt(ignoreErr(strconv.ParseInt(vzString, 10, 64)))

	return &Hailstone{
		px: &Frac{px, big.NewInt(int64(1))},
		py: &Frac{py, big.NewInt(int64(1))},
		pz: &Frac{pz, big.NewInt(int64(1))},
		vx: &Frac{vx, big.NewInt(int64(1))},
		vy: &Frac{vy, big.NewInt(int64(1))},
		vz: &Frac{vz, big.NewInt(int64(1))},
		a: &Frac{vy, vx},
		b: (&Frac{py, big.NewInt(int64(1))}).Sub((&Frac{vy, vx}).Mul(&Frac{px, big.NewInt(int64(1))})),
	}
}

func (h *Hailstone) String() string {
	return fmt.Sprintf("{(%v, %v, %v) [%v,%v,%v] a=%v, b=%v}", h.px, h.py, h.pz, h.vz, h.vy, h.vz, h.a, h.b)
}

var hailstones []*Hailstone

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for _, line := range inputAsLines {
		s := strings.Split(line, "@")
		p := strings.Split(s[0], ",")
		v := strings.Split(s[1], ",")
		h := NewHailstone(
			strings.TrimSpace(p[0]), strings.TrimSpace(p[1]), strings.TrimSpace(p[2]),
			strings.TrimSpace(v[0]), strings.TrimSpace(v[1]), strings.TrimSpace(v[2]),
		)

		if debug {
			fmt.Printf("h = %v\n", h)
			if h.a.Mul(h.px).Add(h.b).Cmp(h.py) != 0 {
				fmt.Printf("Wrong calc.\n")
			}
			nextX := h.px.Add(h.vx)
			nextY := h.py.Add(h.vy)
			if h.a.Mul(nextX).Add(h.b).Cmp(nextY) != 0 {
				fmt.Printf("Wrong calc next step.\n")
			}
		}

		hailstones = append(hailstones, h)
	}
	fmt.Printf("hailstones = %v\n", hailstones)

	res := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			h1 := hailstones[i]
			h2 := hailstones[j]

			fmt.Printf("Hailstone A: %v\nHailstone B: %v\n", h1, h2)
			intersected, x, y := GetIntersection(h1, h2)
			if intersected == 0 {
				fmt.Printf("Hailstones' paths are parallel; they never intersect.\n")
			} else {
				// if debug {
				// 	fmt.Printf("GetIntersection() -> %v, %v\n", x, y)
				// }
				aInPast := (h1.vx.Sign() == 1 && x.Cmp(h1.px) == -1) || (h1.vx.Sign() == -1 && x.Cmp(h1.px) == 1)
				bInPast := (h2.vx.Sign() == 1 && x.Cmp(h2.px) == -1) || (h2.vx.Sign() == -1 && x.Cmp(h2.px) == 1)
				if aInPast && bInPast {
					fmt.Printf("Hailstones' paths crossed in the past for both hailstones.\n")
				} else if aInPast {
					fmt.Printf("Hailstones' paths crossed in the past for hailstone A.\n")
				} else if bInPast {
					fmt.Printf("Hailstones' paths crossed in the past for hailstone B.\n") 
				} else {
					if x.Cmp(MIN_X) >= 0 && x.Cmp(MAX_X) <= 0 && y.Cmp(MIN_Y) >= 0 && y.Cmp(MAX_Y) <= 0 {
						fmt.Printf("Hailstones' paths will cross inside the test area (at x=%v, y=%v).\n", x, y)
						res++
					} else {
						fmt.Printf("Hailstones' paths will cross outside the test area (at x=%v, y=%v).\n", x, y)
					}
				}
			}
		}
	}
	fmt.Printf("res = %v\n", res)
}

func ignoreErr(v int64, _ error) int64 {
	return v
}

func GetIntersection(h1 *Hailstone, h2 *Hailstone) (res int, x, y *Frac) {
	if h1.a.Cmp(h2.a) == 0 {
		return 0, nil, nil
	}

	x = h2.b.Sub(h1.b).Div(h1.a.Sub(h2.a))
	y = h1.a.Mul(x).Add(h1.b)

	return 1, x, y
}