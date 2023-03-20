package deque

type Deque struct {
	front stack
	back  stack
}

type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() (int, bool) {
	if s.len() == 0 {
		return 0, false
	}

	i := len(*s) - 1
	x := (*s)[i]
	*s = (*s)[:i]

	return x, true
}

func (s *stack) len() int {
	return len(*s)
}

func (d *Deque) EnqueueF(x int) {
	d.front.push(x)
	d.balance()
}

func (d *Deque) balance() {
	small, big := order(&d.front, &d.back)
	if small.len() == 0 && big.len() == 1 {
		return
	}
}
