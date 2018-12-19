import copy
import math
import operator
import re


def addr(reg, a, b, c):
    reg[c] = reg[a] + reg[b]


def addrspec(reg, a, b, c):
    reg[c] = reg[a] + reg[b]
    reg[c] = reg[4] + reg[b]


def addi(reg, a, b, c):
    reg[c] = reg[a] + b


def mulr(reg, a, b, c):
    reg[c] = reg[a] * reg[b]


def muli(reg, a, b, c):
    reg[c] = reg[a] * b


def banr(reg, a, b, c):
    reg[c] = reg[a] & reg[b]


def bani(reg, a, b, c):
    reg[c] = reg[a] & b


def borr(reg, a, b, c):
    reg[c] = reg[a] | reg[b]


def bori(reg, a, b, c):
    reg[c] = reg[a] | b


def setr(reg, a, b, c):
    reg[c] = reg[a]


def seti(reg, a, b, c):
    reg[c] = a


def gtir(reg, a, b, c):
    reg[c] = 1 if a > reg[b] else 0


def gtri(reg, a, b, c):
    reg[c] = 1 if reg[a] > b else 0


def gtrr(reg, a, b, c):
    reg[c] = 1 if reg[a] > reg[b] else 0


HOLYNUMBER = math.sqrt(10551343)


def gtrrsqrt(reg, a, b, c):
    reg[c] = 1 if reg[a] > HOLYNUMBER else 0


def eqir(reg, a, b, c):
    reg[c] = 1 if a == reg[b] else 0


def eqri(reg, a, b, c):
    reg[c] = 1 if reg[a] == b else 0


def eqrr(reg, a, b, c):
    reg[c] = 1 if reg[a] == reg[b] else 0


availableOpcodes = [
    addr,
    addi,
    mulr,
    muli,
    banr,
    bani,
    borr,
    bori,
    setr,
    seti,
    gtir,
    gtri,
    gtrr,
    eqir,
    eqri,
    eqrr,
]


class executionData(object):
    def __init__(self, opcode, a, b, c):
        for op in availableOpcodes:
            if opcode == op.__name__:
                if op == gtrr and a == 5:
                    self.Op = gtrrsqrt
                    break
                elif op == addr and c == 0:
                    self.Op = addrspec
                    break
                self.Op = op
                break
        self.A = a
        self.B = b
        self.C = c

    def __str__(self):
        return "op: {} A: {} B: {} C: {}".format(
            self.Op.__name__, self.A, self.B, self.C
        )


def parseData(filepath, instructions):
    f = open(filepath, "r")
    samples = f.read().splitlines()
    ipRegexp = r"#ip (\d+)"
    ipPattern = re.compile(ipRegexp)
    ipMatch = ipPattern.search(samples[0])
    ipReg = int(ipMatch.group(1))

    instrRegexp = r"(.*) (\d+) (\d+) (\d+)"
    instrPattern = re.compile(instrRegexp)
    for sample in samples[1:]:
        instrMatch = instrPattern.search(sample)
        e = executionData(
            instrMatch.group(1), *[int(i) for i in instrMatch.group(2, 3, 4)]
        )
        instructions.append(e)
    return ipReg


from functools import reduce


def factors(n):
    return set(
        reduce(
            list.__add__,
            ([i, n // i] for i in range(1, int(n ** 0.5) + 1) if n % i == 0),
        )
    )


def runProgram(ipReg, instructions):
    registers = [0] * 6
    lenInstruction = len(instructions)
    while registers[ipReg] < lenInstruction:
        i = instructions[registers[ipReg]]
        i.Op(registers, i.A, i.B, i.C)
        registers[ipReg] += 1
    print("Register[0] at end of program: ", registers[0])


def runProgramInit(ipReg, instructions):
    registers = [0] * 6
    registers[0] = 1
    while registers[ipReg] < len(instructions):
        i = instructions[registers[ipReg]]
        if registers[ipReg] == 1:
            return registers[1]
        i.Op(registers, i.A, i.B, i.C)
        registers[ipReg] += 1


def main():
    instructions = []
    ipReg = parseData("day19/input", instructions)
    runProgram(ipReg, instructions)
    print("Modified run:", sum(factors(runProgramInit(ipReg, instructions))))


if __name__ == "__main__":
    main()
