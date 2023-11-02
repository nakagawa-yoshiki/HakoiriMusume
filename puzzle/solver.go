package puzzle

import (
	"fmt"
	"math/rand"
	"slices"
)

func Solve(initialState State, seed int64) []State {
	rnd := rand.New(rand.NewSource(seed))
	state := initialState
	queue := []State{state}
	done := map[string]bool{}
	pres := map[State]State{}
	for len(queue) > 0 {
		state = queue[0]
		queue = queue[1:]
		k := state.key()
		if done[k] {
			continue
		}
		done[k] = true

		if state.isGoal() {
			break
		}

		states := state.nextStates()
		rnd.Shuffle(len(states), func(i, j int) {
			states[i], states[j] = states[j], states[i]
		})

		for _, s := range states {
			queue = append(queue, s)
			if _, ok := pres[s]; !ok {
				pres[s] = state
			}
		}
	}

	results := []State{state}
	for state != initialState {
		pre, ok := pres[state]
		if !ok {
			break
		}
		results = append(results, pre)
		state = pre
	}
	slices.Reverse(results)
	return results
}

const (
	spaceID = 0
	girlID  = 1
)

type State [20]int

func (s State) isGoal() bool {
	return s[13] == girlID && s[14] == girlID && s[17] == girlID && s[18] == girlID
}

func (s State) nextStates() []State {
	used := map[State]bool{s: true}
	states := []State{}
	for id := 1; id <= 10; id++ {
		states1 := s.move(id)
		for _, s1 := range states1 {
			used[s1] = true
			states = append(states, s1)

			states2 := s1.move(id)
			for _, s2 := range states2 {
				if !used[s2] {
					states = append(states, s2)
				}
				used[s2] = true
			}
		}
	}
	return states
}

func (s State) move(id int) []State {
	states := []State{}
	if id == spaceID {
		return states
	}

	cells := []int{}
	for i, v := range s {
		if id == v {
			cells = append(cells, i)
		}
	}

	dxs := []int{0, 0, -1, 1}
	dys := []int{-1, 1, 0, 0}
	for di := 0; di < len(dxs); di++ {
		dx, dy := dxs[di], dys[di]
		ns := s
		ok := false
		cs := append([]int{}, cells...)
		for len(cs) > 0 {
			ci := cs[0]
			cs = cs[1:]
			cx, cy := ci%4, ci/4
			nx, ny := cx+dx, cy+dy
			if !(0 <= nx && nx < 4 && 0 <= ny && ny < 5) {
				ok = false
				break
			}

			ni := nx + ny*4
			if ns[ni] == id {
				cs = append(cs, ci)
				continue
			}
			if ns[ni] != spaceID {
				ok = false
				break
			}
			ns[ni], ns[ci] = ns[ci], ns[ni]
			ok = true
		}
		if ok {
			states = append(states, ns)
		}
	}
	return states
}

func (s State) key() string {
	keys := map[int]int{}
	for i, id := range s {
		keys[id]++
		if i%4 < 3 && s[i] == s[i+1] {
			keys[id]++
		}
	}
	keys[spaceID] = 0

	key := ""
	for _, id := range s {
		key += fmt.Sprint(keys[id])
	}
	return key
}

func (s State) Output(color bool) string {
	ls := [][]string{
		{"　", "　"},
		{"娘", "娘", "娘", "娘"},
		{"父", "親"},
		{"母", "親"},
		{"祖", "父"},
		{"祖", "母"},
		{"兄", "弟"},
		{"茶"},
		{"華"},
		{"書"},
		{"舞"},
	}
	cs := []int{0, 41, 46, 43, 44, 45, 42, 47, 100, 47, 100}

	res := ""
	for i, id := range s {
		if color {
			res += fmt.Sprintf("\x1b[%dm%s\x1b[0m", cs[id], ls[id][0])
		} else {
			res += fmt.Sprint(ls[id][0])
		}
		ls[id] = ls[id][1:]
		if (i+1)%4 == 0 {
			res += "\n"
		}
	}
	return res
}
