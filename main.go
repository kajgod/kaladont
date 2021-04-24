package main

import (
	"fmt"
	"kaladont/inout"
	s "kaladont/structures"
	"os"
)

type Node [4]int

func main() {
	from, to := 97, 123
	words := inout.Read()
	intstr, strint, strLength := s.GetLetterMapping(from, to)
	fundus := s.GetMappedFundus(words, strint, strLength)
	//algo1(words, intstr, strint, strLength, fundus)
	algo2(words, intstr, strint, strLength, fundus)
}

func algo1(
	words []string,
	intstr map[int]string,
	strint map[string]int,
	strLength int,
	fundus s.Fundus,
) {
	max := [2]int{0, 0}
	// defaults
	nodes := make([]Node, 0)
	nodes = append(nodes, Node{}) // first node means we're at the root
	wordsInNodes := make([]string, 0)
	wordsInNodes = append(wordsInNodes, "") // first node doesn't have a word
	var addNodes func(int, int, int, int)
	// build tree
	addNodes = func(pre, suf, length, spawnedFromId int) {
		used := usedWords(nodes, pre, suf, spawnedFromId)
		a := len(fundus[pre][suf]) // available words
		if used >= a {
			return
		}
		l := nodes[spawnedFromId][2] + 1 // number in order
		nodes = append(nodes, Node{pre, suf, l, spawnedFromId})
		wordsInNodes = append(wordsInNodes, fundus[pre][suf][used]) // add a word after last used
		spId := len(nodes) - 1
		if len(wordsInNodes) != len(nodes) {
			os.Exit(3)
		}
		if l > max[1] {
			max[0] = spId
			max[1] = l
			if o := l % 1000; o == 0 {
				writeSolution(nodes, max, fundus)
			}
		}
		p := suf
		for s, v := range fundus[pre] {
			if len(v) == 0 {
				continue
			}
			addNodes(p, s, l, spId)
		}
	}
	// start with top starts
	//iArray, jArray := getTopStarts(fundus, 200, 2)
	for i := 0; i<strLength;i++ {
		for j := 0; j<strLength;j++ {
			fmt.Println()
			fmt.Println("============================================================================")
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("i, i = ", i, j)
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("===========================================================================")
			fmt.Println()
			addNodes(i, j, 0, 0)
		}
	}
	fmt.Println()
	fmt.Println("== final solution ==")
	fmt.Println()
	writeSolution(nodes, max, fundus)
}

func algo2(
	words []string,
	intstr map[int]string,
	strint map[string]int,
	strLength int,
	fundus s.Fundus,
) {
	max := [2]int{0, 0}
	// defaults
	nodes := make([]Node, 0)
	nodes = append(nodes, Node{}) // first node means we're at the root
	wordsInNodes := make([]string, 0)
	wordsInNodes = append(wordsInNodes, "") // first node doesn't have a word
	var addNodes func(int, int, int, int)
	printStep := 1000
	// build tree
	addNodes = func(pre, suf, length, spawnedFromId int) {
		used := usedWords(nodes, pre, suf, spawnedFromId)
		a := len(fundus[pre][suf]) // available words
		if used >= a {
			return
		}
		l := nodes[spawnedFromId][2] + 1 // number in order
		nodes = append(nodes, Node{pre, suf, l, spawnedFromId})
		wordsInNodes = append(wordsInNodes, fundus[pre][suf][used]) // add a word after last used
		spId := len(nodes) - 1
		if len(wordsInNodes) != len(nodes) {
			os.Exit(3)
		}
		if l > max[1] {
			max[0] = spId
			max[1] = l
			if l % printStep == 0 {
				writeSolution(nodes, max, fundus)
				if printStep > 1 {
					printStep = printStep - printStep/30
					fmt.Println(printStep)
				}
			}
		}
		p := suf
		for s, v := range fundus[pre] {
			if len(v) == 0 {
				continue
			}
			addNodes(p, s, l, spId)
		}
	}
	// start with top starts
	i, j:= getTopStarts(fundus)
	addNodes(i, j, 0, 0)
	fmt.Println()
	fmt.Println("== final solution ==")
	fmt.Println()
	writeSolution(nodes, max, fundus)
}

func usedWords(nodes []Node, pre, suf, spawnedFromId int) int {
	count := 0
	var recurse func(int)
	recurse = func(s int) {
		if s == 0 {
			return
		}
		node := nodes[s]
		if pre == node[0] && suf == node[1] {
			count++
		}
		recurse(node[3])
	}
	recurse(spawnedFromId)
	return count
}

func writeSolution(nodes []Node, max [2]int, f s.Fundus) {
	solution := make([]string, 0)
	// copy values to tmp matrix
	fundus := make(s.Fundus, len(f))
	for i := range f {
		fundus[i] = make([][]string, len(f[i]))
		copy(fundus[i], f[i])
		for j := range f[i] {
			fundus[i][j] = make([]string, len(f[i][j]))
			copy(fundus[i][j], f[i][j])
		}
	}
	c := max[0]
	for {
		if nodes[c][3] == 0 {
			break
		}
		c = nodes[c][3]
		p := nodes[c][0]
		s := nodes[c][1]
		solution = append(solution, fundus[p][s][0])
		fundus[p][s] = fundus[p][s][1:]
	}
	for i, j := 0, len(solution)-1; i < j; i, j = i+1, j-1 {
		solution[i], solution[j] = solution[j], solution[i]
	}
	fmt.Println()
	fmt.Println(solution, max)
	fmt.Println()
}

func getTopStarts(fundus s.Fundus)(int, int){
	p, s := 0,0
	mP, mS := 0,0
	for i:=range fundus {
		l:=0
		for j:= range fundus[i]{
			if len(fundus[i][j])>mS {
				mS = len(fundus[i][j])
				s=j
			}
			if len(fundus[i][j])>0 {
				l++
			}
		}
		if l>mP{
			mP = l
			p=i
		}
	}
	fmt.Println("Starting prefix: ", p)
	fmt.Println("Starting sufix: ", s)
	return p,s
}
