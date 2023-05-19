package avltree

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 2, 3, 4}, []int{1, 2, 2, 3, 4}},
		{[]int{2, 1, 3, 2, 4}, []int{1, 2, 2, 3, 4}},
		{[]int{4, 3, 2, 2, 1}, []int{1, 2, 2, 3, 4}},
	}
	for i, test := range tests {
		tree := makeTree(false, test.args)
		if got := tree.Slice(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v\n", i, got, test.want)
		}
	}
}

func TestAddNoDups(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{2, 1, 3, 2, 4}, []int{1, 2, 3, 4}},
		{[]int{4, 3, 2, 2, 1}, []int{1, 2, 3, 4}},
	}
	for i, test := range tests {
		tree := makeTree(true, test.args)
		if got := tree.Slice(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v\n", i, got, test.want)
		}
	}
}

func TestAddPermutations(t *testing.T) {
	want := []int{1, 2, 2, 3, 4}
	for _, args := range getData() {
		tree := makeTree(false, args)
		if got := tree.Slice(); !reflect.DeepEqual(got, want) {
			t.Errorf("%v: got %v, want %v", args, got, want)
		}
	}
}

func TestContains(t *testing.T) {
	tree := New(CmpOrd[int], false)
	for _, i := range []int{1, 2, 3, 4, 5} {
		tree.Add(i)
	}
	tests := []struct {
		arg  int
		want bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, false},
	}
	for i, test := range tests {
		if got := tree.Contains(test.arg); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestDel(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2}, []int{}},
		{[]int{1, 2, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{2, 1, 3, 2, 4}, []int{1, 2, 3, 4}},
		{[]int{4, 3, 2, 2, 1}, []int{1, 2, 3, 4}},
	}
	for i, test := range tests {
		tree := makeTree(false, test.args)
		tree.Del(2)
		if got := tree.Slice(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestDelPermutations(t *testing.T) {
	empty := []int{}
	for _, args := range getData() {
		tree := makeTree(false, args)
		for _, arg := range args {
			tree.Del(arg)
		}
		if got := tree.Slice(); !reflect.DeepEqual(got, empty) || tree.Count() != 0 {
			t.Errorf("%v: got %v and %d, want %v and 0, ", args, got, empty, tree.Count())
		}
	}
}

func TestEach(t *testing.T) {
	tree := New(CmpOrd[int], false)
	n := 0
	tree.Each(func(v int) {
		n += v
	})
	if n != 0 {
		t.Errorf("got %d, want 0", n)
	}
	for i := 1; i < 10; i++ {
		tree.Add(i)
	}
	tree.Each(func(v int) {
		n += v
	})
	if n != 45 {
		t.Errorf("got %d, want 45", n)
	}
}

func TestIsEmpty(t *testing.T) {
	tree := New(CmpOrd[int], false)
	if got := tree.IsEmpty(); got != true {
		t.Errorf("got %v, want true", got)
	}
	tree.Add(0)
	if got := tree.IsEmpty(); got != false {
		t.Errorf("got %v, want false", got)
	}
}

func TestCount(t *testing.T) {
	tree := New(CmpOrd[int], false)
	for i := 0; i < 5; i++ {
		if got := tree.Count(); got != i {
			t.Errorf("got %d, want %d", got, i)
		}
		tree.Add(i)
	}
	cnt := 5
	for i := 0; i < 5; i++ {
		tree.Del(i)
		cnt--
		if got := tree.Count(); got != cnt {
			t.Errorf("got %d, want %d", got, cnt)
		}
	}
}

func makeTree(nodups bool, args []int) *Tree[int] {
	tree := New(CmpOrd[int], nodups)
	for _, arg := range args {
		tree.Add(arg)
	}
	return tree
}

func getData() [][]int {
	f, err := os.Open("test_args.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	result := [][]int{}
	for scanner.Scan() {
		sl := make([]int, 5)
		for i, arg := range strings.Split(scanner.Text(), " ") {
			v, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			sl[i] = v
		}
		result = append(result, sl)
	}
	return result
}

func TestClone(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 2, 3, 4}, []int{1, 2, 2, 3, 4}},
		{[]int{2, 1, 3, 2, 4}, []int{1, 2, 2, 3, 4}},
		{[]int{4, 3, 2, 2, 1}, []int{1, 2, 2, 3, 4}},
	}
	for i, test := range tests {
		tree := makeTree(false, test.args)
		clone := tree.Clone()
		if got := clone.Slice(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestNewPanic(t *testing.T) {
	defer func() { recover() }()
	New[int](nil, false)
	t.Errorf("did not panic")
}
