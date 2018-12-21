import bisect
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
    lines = utils.readlines(filepath)
    ipMatch = re.search(r"#ip (\d+)", lines[0])
    ip = int(ipMatch.group(1))

    for line in lines[1:]:
        instrMatch = re.search(r"(.*) (\d+) (\d+) (\d+)", line)
        e = executionData(
            instrMatch.group(1), *[int(i) for i in instrMatch.group(2, 3, 4)]
        )
        instructions.append(e)
    return ip


def runProgram(ip, instructions):
    registers = [0] * 6
    seenNumbers = []
    while True:
        i = instructions[registers[ip]]
        if i.Op == utils.addi and i.A == 2 and i.B == 1 and i.C == 2:
            registers[2] = (registers[3] // 256) - 1
        if i.Op == utils.eqrr:
            if registers[4] in seenNumbers:
                print("Integer to break as earliest:", seenNumbers[0])
                print("Integer to break at the latest:", seenNumbers[-1])
                break
            else:
                seenNumbers.append(registers[4])
        i.Op(registers, i.A, i.B, i.C)
        registers[ip] += 1


def main():
    instructions = []
    ip = parseData("day21/input", instructions)
    runProgram(ip, instructions)


if __name__ == "__main__":
    main()
