import re
import utils


class executionData(object):
    def __init__(self, opcode, a, b, c):
        self.Opcode = opcode
        self.A = a
        self.B = b
        self.C = c


class sampleData(executionData):
    def __init__(self, inputReg, outputReg, opcode, a, b, c):
        executionData.__init__(self, opcode, a, b, c)
        self.InputReg = inputReg
        self.OutputReg = outputReg
        self.MatchingOpcodes = []


def parseTrainingData(filepath):
    f = open(filepath, "r")
    samples = re.split(r"\n\n", f.read())
    sampleRegexp = (
        r"Before: \[(\d+), (\d+), (\d+), (\d+)\]\n"
        r"(\d+) (\d+) (\d+) (\d+)\n"
        r"After:  \[(\d+), (\d+), (\d+), (\d+)\]"
    )
    output = []
    for sample in samples:
        m = re.search(sampleRegexp, sample)
        s = sampleData(
            [int(i) for i in m.group(1, 2, 3, 4)],
            [int(i) for i in m.group(9, 10, 11, 12)],
            *[int(i) for i in m.group(5, 6, 7, 8)]
        )
        output.append(s)
    return output


def checkTrainingData(trainingData):
    numSamplesMinTripleMatches = 0
    for sample in trainingData:
        for op in utils.availableOpcodes:
            testResult = sample.InputReg[:]
            op(testResult, sample.A, sample.B, sample.C)
            if testResult == sample.OutputReg:
                sample.MatchingOpcodes.append(op)
        if len(sample.MatchingOpcodes) >= 3:
            numSamplesMinTripleMatches += 1
    print(numSamplesMinTripleMatches)


def reduceMatchingOpcodes(trainingData):
    opcodeTable = {}
    while len(opcodeTable) != len(utils.availableOpcodes):
        opcode, op = [
            (s.Opcode, s.MatchingOpcodes[0])
            for s in trainingData
            if len(s.MatchingOpcodes) == 1
        ][0]
        opcodeTable[opcode] = op
        for sample in trainingData:
            if op in sample.MatchingOpcodes:
                sample.MatchingOpcodes.remove(op)
    return opcodeTable


def parseData(filepath):
    samples = utils.readlines(filepath)
    output = []
    for sample in samples:
        m = re.search(r"(\d+) (\d+) (\d+) (\d+)", sample)
        e = executionData(*[int(i) for i in m.group(1, 2, 3, 4)])
        output.append(e)
    return output


def runData(execData, opcodeTable):
    register = [0] * 4
    for exec in execData:
        opcodeTable[exec.Opcode](register, exec.A, exec.B, exec.C)
    print(register)


def main():
    trainingData = parseTrainingData("day16/input")
    checkTrainingData(trainingData)
    opcodeTable = reduceMatchingOpcodes(trainingData)
    execData = parseData("day16/input2")

    runData(execData, opcodeTable)


if __name__ == "__main__":
    main()
