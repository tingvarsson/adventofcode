import re

inches = 1000
fabric = [[0 for i in range(inches)] for j in range(inches)]

f = open("../input", "r")
fl = f.read().splitlines()

def createClaim(line):
    #           id      x     y      xsize ysize
    regex = r'#(\d+) @ (\d+),(\d+): (\d+)x(\d+)'
    pattern = re.compile(regex)
    m = pattern.match(line)
    claim = { 'id' : int(m.group(1)),
              'x' : int(m.group(2)),
              'y' : int(m.group(3)),
              'xsize' : int(m.group(4)),
              'ysize' : int(m.group(5)) }
    return claim

def addClaim(claim):
    for x in range(claim["xsize"]):
        for y in range(claim["ysize"]):
            fabric[claim["x"]+x][claim["y"]+y] += 1

def testClaim(claim):
    for x in range(claim["xsize"]):
        for y in range(claim["ysize"]):
            if fabric[claim["x"]+x][claim["y"]+y] >= 2:
                return False
    return True

claims = []
for line in fl:
    claims.append(createClaim(line))

for claim in claims:
    addClaim(claim)

sumofmultiple = sum(sum(1 for y in x if y >= 2) for x in fabric)
print(sumofmultiple)

for claim in claims:
    if testClaim(claim):
        print(claim)

