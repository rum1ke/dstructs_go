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

	for _, ent := range testTable {
		newStack := NewStackString()

		for idx := range ent.pushing {
			newStack.Push(&ent.pushing[idx])

			if *newStack.top.value != ent.must[idx] {
				t.Error("wrong push. Pushed, Must:", ent.pushing[idx], ent.must[idx])
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

	for _, ent := range testTable {
		newStack := NewStackString()

		for idx := range ent.pushing {
			newStack.Push(&ent.pushing[idx])
		}

		for idx := range ent.must {
			popped := newStack.Pop()

			if ent.must[idx] != *popped {
				t.Error("wrong pop. Popped, Must:", *popped, ent.must[idx])
			}
		}
	}
}
