package program

// Program is a thingy
type Program struct {
	memory       []int
	instr        int
	relativeBase int
	input        []int
	Output       []int
	Halted       bool
}

// New constructs a Program
func New(memory []int) Program {
	m := make([]int, 10000)
	for i := range memory {
		m[i] = memory[i]
	}
	return Program{m, 0, 0, []int{}, []int{}, false}
}

func (p Program) get(addr, mode int) int {
	param := p.memory[addr]
	if mode == 0 {
		return p.memory[param]
	} else if mode == 1 {
		return param
	} else if mode == 2 {
		return p.memory[p.relativeBase+param]
	}
	return 0
}

func (p *Program) set(addr, mode, value int) {
	param := p.memory[addr]
	if mode == 2 {
		p.memory[p.relativeBase+param] = value
	} else {
		p.memory[param] = value
	}
}

// Run makes the program go around
func (p *Program) Run(newInput []int) bool {
	p.input = append(p.input, newInput...)

	for true {
		op := p.memory[p.instr] % 100
		m1 := (p.memory[p.instr] / 100) % 10
		m2 := (p.memory[p.instr] / 1000) % 10
		m3 := (p.memory[p.instr] / 10000) % 10
		if op == 1 {
			p.set(p.instr+3, m3, p.get(p.instr+1, m1)+p.get(p.instr+2, m2))
			p.instr += 4
		} else if op == 2 {
			p.set(p.instr+3, m3, p.get(p.instr+1, m1)*p.get(p.instr+2, m2))
			p.instr += 4
		} else if op == 3 {
			if len(p.input) == 0 {
				break
			}
			p.set(p.instr+1, m1, p.input[0])
			p.input = p.input[1:] // dequeue
			p.instr += 2
		} else if op == 4 {
			p.Output = append(p.Output, p.get(p.instr+1, m1))
			p.instr += 2
		} else if op == 5 {
			if p.get(p.instr+1, m1) != 0 {
				p.instr = p.get(p.instr+2, m2)
			} else {
				p.instr += 3
			}
		} else if op == 6 {
			if p.get(p.instr+1, m1) == 0 {
				p.instr = p.get(p.instr+2, m2)
			} else {
				p.instr += 3
			}
		} else if op == 7 {
			if p.get(p.instr+1, m1) < p.get(p.instr+2, m2) {
				p.set(p.instr+3, m3, 1)
			} else {
				p.set(p.instr+3, m3, 0)
			}
			p.instr += 4
		} else if op == 8 {
			if p.get(p.instr+1, m1) == p.get(p.instr+2, m2) {
				p.set(p.instr+3, m3, 1)
			} else {
				p.set(p.instr+3, m3, 0)
			}
			p.instr += 4
		} else if op == 9 {
			p.relativeBase += p.get(p.instr+1, m1)
			p.instr += 2
		} else if op == 99 {
			p.Halted = true
			break
		}
	}
	return p.Halted
}
