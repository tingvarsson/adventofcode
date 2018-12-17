from collections import deque
import re

def parseInput(filepath, output):
  f = open(filepath, "r")
  data = f.read().splitlines()
  xLineRegexp = r'x=(\d+), y=(\d+)\.\.(\d+)'
  yLineRegexp = r'y=(\d+), x=(\d+)\.\.(\d+)'
  xLinePattern = re.compile(xLineRegexp)
  yLinePattern = re.compile(yLineRegexp)
  for line in data:
    m = xLinePattern.search(line)
    if m is not None:
      for y in range(int(m.group(2)), int(m.group(3))+1):
        output.append((int(m.group(1)), y))
    m = yLinePattern.search(line)
    if m is not None:
      for x in range(int(m.group(2)), int(m.group(3))+1):
        output.append((x, int(m.group(1))))

def prepArea():
  maxX = max(clayCoords, key=lambda c: c[0])[0] + 5
  maxY = max(clayCoords, key=lambda c: c[1])[1] + 1
  area = [['.' for x in range(maxX)] for y in range(maxY)]
  area[0][500] = '+'
  for clay in clayCoords:
    area[clay[1]][clay[0]] = '#'

  return area

def printArea(area):
  for row in area:
    print(''.join(row[400:900]))

def endOfAreaBelow(area, pos):
  return pos[1] + 1 == len(area)
def runningWaterBelow(area, pos):
  return area[pos[1]+1][pos[0]] == '|'
def runningWaterLeft(area, pos):
  return area[pos[1]][pos[0]-1] == '|'
def runningWaterRight(area, pos):
  return area[pos[1]][pos[0]+1] == '|'
def sandBelow(area, pos):
  return area[pos[1]+1][pos[0]] == '.'
def sandLeft(area, pos):
  return area[pos[1]][pos[0]-1] == '.'
def sandRight(area, pos):
  return area[pos[1]][pos[0]+1] == '.'
def clayLeft(area, pos):
  return area[pos[1]][pos[0]-1] == '#'
def clayRight(area, pos):
  return area[pos[1]][pos[0]+1] == '#'


def runWater(area):
  currentPos = (500, 0)
  stack = deque()
  currentRow = deque(currentPos)
  while True:
    if currentPos == 500: # FIXME: ugly hack
      break
    elif endOfAreaBelow(area, currentPos):
      currentRow = stack.pop()
      currentPos = currentRow[0]
    elif runningWaterBelow(area, currentPos):
      if runningWaterLeft(area, currentPos):
        currentRow = stack.pop()
        currentPos = currentRow[0]
      elif runningWaterRight(area, currentPos):
        currentPos = currentRow[-1]
      else:
        currentRow = stack.pop()
        currentPos = currentRow[0]
    elif sandBelow(area, currentPos):
      area[currentPos[1]+1][currentPos[0]] = '|'
      currentPos = (currentPos[0], currentPos[1]+1)
      stack.append(currentRow)
      currentRow = deque([currentPos])
    elif sandLeft(area, currentPos):
      area[currentPos[1]][currentPos[0]-1] = '|'
      currentPos = (currentPos[0]-1, currentPos[1])
      currentRow.appendleft(currentPos)
    elif sandRight(area, currentPos):
      area[currentPos[1]][currentPos[0]+1] = '|'
      currentPos = (currentPos[0]+1, currentPos[1])
      currentRow.append(currentPos)
    elif clayRight(area, currentPos):
      leftMost = currentRow[0]
      if sandLeft(area, leftMost):
        currentRow = stack.pop()
        currentPos = currentRow[0]
        continue
      while currentRow:
        pos = currentRow.pop()
        area[pos[1]][pos[0]] = '~'
      currentRow = stack.pop()
      currentPos = currentRow[0]
    elif clayLeft(area, currentPos):
      currentPos = currentRow[-1]

clayCoords = []
parseInput("day17/input", clayCoords)
area = prepArea()
runWater(area)
minY = min(clayCoords, key=lambda c: c[1])[1]
print("Total water:", sum(sum(1 for pos in row if pos == "~" or pos == "|") for row in area[minY:]))
print("Still water:", sum(sum(1 for pos in row if pos == "~") for row in area))
