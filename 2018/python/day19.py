from functools import reduce
import re
import utils


class executionData(object):
    def __init__(self, opcode, a, b, c):
        for op in utils.availableOpcodes:
            if opcode == op.__name__:
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
        if i.Op == utils.eqrr:
            # from one hard coded value to one that at least works for two inputs
            # there is only one eqrr that checks if a factor is found
            # presumes the second argument is the number under test
            return registers[i.B]
        i.Op(registers, i.A, i.B, i.C)
        registers[ipReg] += 1


def main():
    instructions = []
    ipReg = parseData("day19/input", instructions)
    runProgram(ipReg, instructions)
    print("Modified run:", sum(factors(runProgramInit(ipReg, instructions))))


if __name__ == "__main__":
    main()
