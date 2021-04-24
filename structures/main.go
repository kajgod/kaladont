package structures

import "fmt"

type Fundus [][][]string

// generate two way mapping between characters and integers ex. 0 - aa
func GetLetterMapping(from, to int) (map[int]string, map[string]int, int) {
	intstr, strint := map[int]string{}, map[string]int{}
	c := 0
	for i := from; i <= to; i++ {
		a := fmt.Sprintf("%c", i)
		if i == to {
			a = "-"
		}
		for j := from; j <= to; j++ {
			b := fmt.Sprintf("%c", j)
			if j == to {
				b = "-"
			}
			d := fmt.Sprintf("%s%s", a, b)
			intstr[c] = d
			strint[d] = c
			c++
		}
	}
	l := c + from - to - 2
	return intstr, strint, l
}

func buildEmptyFundus(l int) Fundus {
	r := make([][][]string, l)
	for i := 0; i < l; i++ {
		r[i] = make([][]string, l)
		for j := 0; j < l; j++ {
			r[i][j] = make([]string, 0)
		}
	}
	return r
}

// GetMappedFundus maps real characters to possible values
func GetMappedFundus(words []string, strint map[string]int, l int) Fundus {
	r := buildEmptyFundus(l)
	for _, w := range words {
		p := strint[w[:2]]
		s := strint[w[len(w)-2:]]
		r[p][s] = append(r[p][s], w)
	}
	return r
}
