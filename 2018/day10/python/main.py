import re

class Point(object):
    def __init__(self, x, y, vx, vy):
        self.posX = x
        self.posY = y
        self.velocityX = vx
        self.velocityY = vy

def parseInput():
    f = open("../input", "r")
    lines = f.read().splitlines()
    # "position=< 9,  1> velocity=< 0,  2>""
    regex = r'position=<(.*), (.*)> velocity=<(.*), (.*)>'
    pattern = re.compile(regex)
    points = []
    for line in lines:
        m = pattern.search(line)
        p = Point(int(m.group(1)), 
                  int(m.group(2)), 
                  int(m.group(3)), 
                  int(m.group(4)))
        points.append(p)
    return points

def normalizePoints(p):
    minX = min(points, key=lambda p: p.posX).posX
    minY = min(points, key=lambda p: p.posY).posY

    for p in points:
        p.posX -= minX
        p.posY -= minY
    return points

points = parseInput()

minHeight = None
time = 0
while True:
    for p in points:
        p.posX += p.velocityX
        p.posY += p.velocityY
    
    points = normalizePoints(points)

    maxY = max(points, key=lambda p: p.posY).posY
    if minHeight == None or maxY < minHeight:
        minHeight = maxY
    elif maxY >= minHeight:
        for p in points:
            p.posX -= p.velocityX
            p.posY -= p.velocityY
        points = normalizePoints(points)
        break
    time += 1

print("time: %d" % time)

maxX = max(points, key=lambda p: p.posX).posX
maxY = max(points, key=lambda p: p.posY).posY
sky = [["." for x in range(maxX+1)] for y in range(maxY+1)]
for p in points:
    sky[p.posY][p.posX] = "#"

for y in sky:
    print(''.join([x for x in y]))