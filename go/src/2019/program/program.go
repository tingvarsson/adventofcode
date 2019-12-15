package program

// Program is a thingy
type Program struct {
	memory       []int
	instr        int
	relativeBase int
	Input        []int
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

// GetOutput pops and returns all output
func (p *Program) GetOutput() []int {
	o := p.Output
	p.Output = []int{}
	return o
}

// PopOutput pops and returns the oldest output
func (p *Program) PopOutput() int {
	o := p.Output[0]
	p.Output = p.Output[1:]
	return o
}

// RunWoArgs makes the program go around without any input
func (p *Program) RunWoArgs() bool {
	return p.Run([]int{})
}

// Run makes the program go around
func (p *Program) Run(newInput []int) bool {
	p.Input = append(p.Input, newInput...)

	for true {
		op := p.memory[p.instr] % 100
		m1 := (p.memory[p.instr] / 100) % 10
		m2 := (p.memory[p.instr] / 1000) % 10
		m3 := (p.memory[p.instr] / 10000) % 10
		switch op {
		case 1:
			p.set(p.instr+3, m3, p.get(p.instr+1, m1)+p.get(p.instr+2, m2))
			p.instr += 4
		case 2:
			p.set(p.instr+3, m3, p.get(p.instr+1, m1)*p.get(p.instr+2, m2))
			p.instr += 4
		case 3:
			if len(p.Input) == 0 {
				return p.Halted
			}
			p.set(p.instr+1, m1, p.Input[0])
			p.Input = p.Input[1:] // dequeue
			p.instr += 2
		case 4:
			p.Output = append(p.Output, p.get(p.instr+1, m1))
			p.instr += 2
		case 5:
			if p.get(p.instr+1, m1) != 0 {
				p.instr = p.get(p.instr+2, m2)
			} else {
				p.instr += 3
			}
		case 6:
			if p.get(p.instr+1, m1) == 0 {
				p.instr = p.get(p.instr+2, m2)
			} else {
				p.instr += 3
			}
		case 7:
			if p.get(p.instr+1, m1) < p.get(p.instr+2, m2) {
				p.set(p.instr+3, m3, 1)
			} else {
				p.set(p.instr+3, m3, 0)
			}
			p.instr += 4
		case 8:
			if p.get(p.instr+1, m1) == p.get(p.instr+2, m2) {
				p.set(p.instr+3, m3, 1)
			} else {
				p.set(p.instr+3, m3, 0)
			}
			p.instr += 4
		case 9:
			p.relativeBase += p.get(p.instr+1, m1)
			p.instr += 2
		case 99:
			p.Halted = true
			return p.Halted
		}
	}
	return p.Halted
}

func (p Program) get(addr, mode int) int {
	param := p.memory[addr]
	switch mode {
	case 0:
		return p.memory[param]
	case 1:
		return param
	case 2:
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
