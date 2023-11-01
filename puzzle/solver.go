package puzzle

import (
	"fmt"
	"math/rand"
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
		k := state.Key()
		if done[k] {
			continue
		}
		done[k] = true

		if state.IsGoal() {
			break
		}

		states := state.NextStates()
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

	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}
	return results
}

const SpaceID = 0
const GirlID = 1

type State [20]int

func (s State) IsGoal() bool {
	return s[13] == GirlID && s[14] == GirlID && s[17] == GirlID && s[18] == GirlID
}

func (s State) NextStates() []State {
	used := map[State]bool{s: true}
	states := []State{}
	for id := 1; id <= 10; id++ {
		states1 := s.Move(id)
		for _, s1 := range states1 {
			used[s1] = true
			states = append(states, s1)

			states2 := s1.Move(id)
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

func (s State) Move(id int) []State {
	states := []State{}
	if id == SpaceID {
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
			if ns[ni] != SpaceID {
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

func (s State) Key() string {
	keys := map[int]int{}
	for i, id := range s {
		keys[id]++
		if i%4 < 3 && s[i] == s[i+1] {
			keys[id]++
		}
	}
	keys[SpaceID] = 0

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
