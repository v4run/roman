package roman

import (
	"fmt"
	"strings"
)

var (
	decimals = map[string]int{"": 0, "I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	values   = []int{1, 5, 10, 50, 100, 500, 1000}
	romans   = map[int]string{0: "", 1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}
)

// Roman represent a roman numeral
type Roman struct {
	roman  string
	arabic int
}

// New creates a Roman from r (roman number). Throws an error if an invalid roman numberal is given.
func New(r string) (*Roman, error) {
	upper := strings.ToUpper(r)
	arabic, err := toArabic(upper)
	if err != nil {
		return nil, err
	}
	return &Roman{roman: upper, arabic: arabic}, nil
}

// FromArabic creates a Roman from a (arabic number).
func FromArabic(a int) *Roman {
	r := toRoman(a)
	return &Roman{roman: r, arabic: a}
}

// Roman returns roman representation of the number.
func (r Roman) Roman() string {
	return r.roman
}

// Roman returns arabic representation of the number.
func (r Roman) Arabic() int {
	return r.arabic
}

func toRoman(a int) string {
	t := a
	r := ""
	if decimal, ok := romans[a]; ok {
		return decimal
	}
	tp := 1
	for t > 0 {
		q := t % 10
		number := q * tp
		if decimal, ok := romans[number]; ok {
			r = decimal + r
		} else {
			for i := 0; i < len(values); i++ {
				iq := number / values[i]
				if iq == 3 {
					for ; iq > 0; iq-- {
						r = romans[values[i]] + r
					}
					break
				} else if number < values[i] {
					if values[i]-number == values[i]/10 {
						r = romans[values[i]/10] + romans[values[i]] + r
					} else {
						r = romans[values[i]/5] + romans[values[i]] + r
					}
					break
				}
			}
		}
		t /= 10
		tp *= 10
	}
	return r
}

func toArabic(v string) (int, error) {
	if len(v) == 0 {
		return -1, fmt.Errorf("number can't be empty")
	}
	counts := map[string]int{
		"I": 0,
		"V": 0,
		"X": 0,
		"L": 0,
		"C": 0,
		"D": 0,
		"M": 0,
	}
	prev := ""
	nxt := ""
	compoundFirst := 0
	length := len(v)
	sum := 0
	for i := 0; i < length; i++ {

		cur := string(v[i])

		if i != length-1 {
			nxt = string(v[i+1])
		} else {
			nxt = ""
		}
		counts[cur]++

		if compoundFirst > 0 && decimals[nxt] >= compoundFirst {
			return -1, fmt.Errorf("invalid symbol %s at position %d", nxt, i+2)
		}

		switch cur {
		case `D`, `L`, `V`:
			if counts[cur] > 1 {
				return -1, fmt.Errorf("invalid symbol %s at position %d", cur, i+1)
			}
			if nxt != "" && decimals[cur] < decimals[nxt] {
				return -1, fmt.Errorf("invalid symbol %c at position %d", v[i+1], i+2)
			}
			compoundFirst = 0
			sum += decimals[cur]

		case `I`:
			counts["X"] = 0
			counts["C"] = 0
			counts["M"] = 0
			if counts[cur] > 3 {
				return -1, fmt.Errorf("invalid symbol %s at position %d", cur, i+1)
			}
			if decimals[cur] < decimals[nxt] {
				switch nxt {
				case "V", "X":
					compoundFirst = decimals[cur]
					sum -= compoundFirst
				default:
					return -1, fmt.Errorf("invalid symbol %s at position %d", cur, i+1)
				}
			} else {
				compoundFirst = 0
				sum += decimals[cur]
			}

		case `X`:
			counts["I"] = 0
			counts["C"] = 0
			counts["M"] = 0
			if counts[cur] > 3 {
				return -1, fmt.Errorf("invalid symbol %s at position %d", cur, i+1)
			}

			if decimals[cur] < decimals[nxt] {
				switch nxt {
				case "L", "C":
					compoundFirst = decimals[cur]
					sum -= compoundFirst
				default:
					return -1, fmt.Errorf("invalid symbol %s at position %d", cur, i+1)
				}
			} else {
				compoundFirst = 0
				sum += decimals[cur]
			}

		case `C`:
			counts["I"] = 0
			counts["X"] = 0
			counts["M"] = 0
			if counts[cur] > 3 {
				return -1, fmt.Errorf("invalid symbol %s at position %d", cur, i+1)
			}
			if decimals[cur] < decimals[nxt] {
				compoundFirst = decimals[cur]
				sum -= compoundFirst
			} else {
				compoundFirst = 0
				sum += decimals[cur]
			}

		case `M`:
			counts["I"] = 0
			counts["X"] = 0
			counts["C"] = 0
			if counts[cur] > 3 {
				return -1, fmt.Errorf("invalid symbol %s at position %d", cur, i+1)
			}
			compoundFirst = 0
			sum += decimals[cur]

		default:
			return -1, fmt.Errorf("invalid symbol %s at position %d", cur, i+1)
		}

		if i > 0 && counts[prev] >= 2 && decimals[prev] < decimals[cur] {
			return -1, fmt.Errorf("invalid symbol %s at position %d", cur, i+1)
		}

		prev = cur
	}
	return sum, nil
}
