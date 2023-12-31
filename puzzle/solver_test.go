package puzzle

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		initialState State
		seed         int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"stage0", args{Stages[0], 0}, 81 + 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.initialState, tt.args.seed); len(got) != tt.want {
				t.Errorf("Solve() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestState_isGoal(t *testing.T) {
	tests := []struct {
		name string
		s    State
		want bool
	}{
		{"goal", State{
			4, 2, 5, 3,
			4, 2, 5, 3,
			6, 6, 7, 8,
			10, 1, 1, 0,
			9, 1, 1, 0,
		}, true},
		{"start", Stages[0], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isGoal(); got != tt.want {
				t.Errorf("State.isGoal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestState_nextStates(t *testing.T) {
	tests := []struct {
		name string
		s    State
		want int
	}{
		{"stage0:start", Stages[0], 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.nextStates(); len(got) != tt.want {
				t.Errorf("State.nextStates() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestState_move(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		s    State
		args args
		want []State
	}{
		{
			"stage0", Stages[0], args{7}, []State{
				{
					2, 1, 1, 3,
					2, 1, 1, 3,
					4, 6, 6, 5,
					4, 0, 8, 5,
					9, 7, 0, 10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.move(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("State.move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestState_key(t *testing.T) {
	tests := []struct {
		name string
		s    State
		want string
	}{
		{"stage0", Stages[0], "26622662233221121001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.key(); got != tt.want {
				t.Errorf("State.key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestState_Output(t *testing.T) {
	type args struct {
		color bool
	}
	tests := []struct {
		name string
		s    State
		args args
		want string
	}{
		{"stage0", Stages[0], args{color: false},
			`父娘娘母
親娘娘親
祖兄弟祖
父茶華母
書　　舞
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Output(tt.args.color); got != tt.want {
				t.Errorf("State.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}
