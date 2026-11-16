package algo

type State struct {
	arr           []int
	LocalProgress int
	GlobalProgres int
}

func (s *State) Arr() []int {
	dest := make([]int, len(s.arr))
	copy(dest, s.arr)
	return dest
}

type StateFn func(s *State) StateFn

func startBubble(s *State) StateFn {
	return walkRowBubble
}
func endBubble(s *State) StateFn {
	return nil
}

// walkRowBubble returns advanceRowBubble
func walkRowBubble(s *State) StateFn {

	// This narrowing is actually optimizing;
	// we can make this codition as
	// s.LocalProgress < len(s.arr)-1
	for s.LocalProgress < len(s.arr)-s.GlobalProgres-1 {
		left, right := s.arr[s.LocalProgress], s.arr[s.LocalProgress+1]
		if left > right {
			s.arr[s.LocalProgress] = right
			s.arr[s.LocalProgress+1] = left
		}
		s.LocalProgress++
		// fmt.Println(s.arr)
	}
	s.GlobalProgres++
	return advanceRowBubble
}

// Advance returns Walk or End
func advanceRowBubble(s *State) StateFn {
	if s.GlobalProgres == len(s.arr)-1 {
		return endBubble
	}
	s.LocalProgress = 0
	return walkRowBubble
}

func BubbleSort(arr []int) []int {
	s := &State{arr: arr}
	for state := startBubble; state != nil; {
		state = state(s)
	}
	return s.Arr()
}
