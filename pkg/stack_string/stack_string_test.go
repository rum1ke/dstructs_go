package stack_string

import "testing"

func TestPush(t *testing.T) {
	testTable := []struct {
		pushing []string
		must    []string
	}{
		{
			pushing: []string{"1", "2", "3"},
			must:    []string{"1", "2", "3"},
		},
	}

	for _, test := range testTable {
		newStack := NewStackString()

		for idx := range test.pushing {
			newStack.Push(&test.pushing[idx])

			if *newStack.top.value != test.must[idx] {
				t.Error("Wrong push. Pushed, Must:", test.pushing[idx], test.must[idx])
			}
		}
	}
}

func TestPop(t *testing.T) {
	testTable := []struct {
		pushing []string
		must    []string
	}{
		{
			pushing: []string{"1", "2", "3"},
			must:    []string{"3", "2", "1"},
		},
	}

	for _, test := range testTable {
		newStack := NewStackString()

		for idx := range test.pushing {
			newStack.Push(&test.pushing[idx])
		}

		for idx := range test.must {
			popped := newStack.Pop()

			if test.must[idx] != *popped {
				t.Error("Wrong pop. Popped, Must:", *popped, test.must[idx])
			}
		}
	}
}
