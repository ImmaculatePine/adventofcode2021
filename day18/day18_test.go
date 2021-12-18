package day18

import "testing"

func TestExplode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"[[[[0,7],4],[15,[0,13]]],[1,1]]", "[[[[0,7],4],[15,[0,13]]],[1,1]]"},
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}

	for _, test := range tests {
		tree, err := parseTree(test.input)
		if err != nil {
			t.Fatalf("failed to parse tree, %v", err)
		}
		explode(tree)
		got := tree.ToString()
		if got != test.want {
			t.Fatalf("got %s, want %s", got, test.want)
		}
	}
}
