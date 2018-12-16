import re

def addr(reg, a, b, c):
  reg[c] = reg[a] + reg[b]
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
def eqir(reg, a, b, c):
  reg[c] = 1 if a == reg[b] else 0
def eqri(reg, a, b, c):
  reg[c] = 1 if reg[a] == b else 0
def eqrr(reg, a, b, c):
  reg[c] = 1 if reg[a] == reg[b] else 0

availableOpcodes = [addr, addi, 
           mulr, muli, 
           banr, bani,
           borr, bori,
           setr, seti,
           gtir, gtri, gtrr,
           eqir, eqri, eqrr]

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

def parseTrainingData(filepath, output):
  f = open(filepath, "r")
  samples = re.split(r'\n\n', f.read())
  sampleRegexp = r'Before: \[(\d+), (\d+), (\d+), (\d+)\]\n'\
                 r'(\d+) (\d+) (\d+) (\d+)\n'\
                 r'After:  \[(\d+), (\d+), (\d+), (\d+)\]'
  samplePattern = re.compile(sampleRegexp)
  for sample in samples:
    m = samplePattern.search(sample)
    s = sampleData([int(i) for i in m.group(1,2,3,4)],
                   [int(i) for i in m.group(9,10,11,12)],
                   *[int(i) for i in m.group(5,6,7,8)])
    output.append(s)

def checkTrainingData(trainingData):
  numSamplesMinTripleMatches = 0
  for sample in trainingData:
    for op in availableOpcodes:
      testResult = sample.InputReg[:]
      op(testResult, sample.A, sample.B, sample.C)
      if testResult == sample.OutputReg:
        sample.MatchingOpcodes.append(op)
    if len(sample.MatchingOpcodes) >= 3:
      numSamplesMinTripleMatches += 1
  print(numSamplesMinTripleMatches)

def reduceMatchingOpcodes(traningData, opcodeTable):
  while len(opcodeTable) != len(availableOpcodes):
    opcode, op = [(s.Opcode, s.MatchingOpcodes[0]) for s in trainingData if len(s.MatchingOpcodes) == 1][0]
    opcodeTable[opcode] = op
    for sample in trainingData:
      if op in sample.MatchingOpcodes:
        sample.MatchingOpcodes.remove(op)

def parseData(filepath, output):
  f = open(filepath, "r")
  samples = f.read().splitlines()
  sampleRegexp = r'(\d+) (\d+) (\d+) (\d+)'
  samplePattern = re.compile(sampleRegexp)
  for sample in samples:
    m = samplePattern.search(sample)
    e = executionData(*[int(i) for i in m.group(1,2,3,4)])
    output.append(e)

def runData(execData, opcodeTable):
  register = [0] * 4
  for exec in execData:
    opcodeTable[exec.Opcode](register, exec.A, exec.B, exec.C)
  print(register)

trainingData = []
opcodeTable = {}
parseTrainingData("day16/input", trainingData)
checkTrainingData(trainingData)
reduceMatchingOpcodes(trainingData, opcodeTable)
execData = []
parseData("day16/input2", execData)
runData(execData, opcodeTable)