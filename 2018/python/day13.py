class cart(object):
  def __init__(self, dir, x, y):
    self.x = x
    self.y = y
    self.dir = dir
    self.nextTurn = "l"
    self.crashed = False

  def changeDir(self, track):
    dr = dirRules[self.dir]
    if track == "+":
      self.dir = dr.intersection[self.nextTurn]
      self.nextTurn = turningOrder[self.nextTurn]
    else:
      self.dir = dr.dirChange[track]

class directionRule(object):
  def __init__(self, xOffset, yOffset, dirChange, intersection):
    self.xOffset = xOffset
    self.yOffset = yOffset
    self.dirChange = dirChange
    self.intersection = intersection

cartParseRules = {"^":"|", ">":"-", "v":"|", "<":"-"}

dirRules = {"^":directionRule(0, -1, {"/":">", "|":"^", "\\":"<", "-":"X"}, {"l":"<","s":"^","r":">"}),
            "v":directionRule(0,  1, {"/":"<", "|":"v", "\\":">", "-":"X"}, {"l":">","s":"v","r":"<"}),
            "<":directionRule(-1, 0, {"/":"v", "|":"X", "\\":"^", "-":"<"}, {"l":"v","s":"<","r":"^"}),
            ">":directionRule( 1, 0, {"/":"^", "|":"X", "\\":"v", "-":">"}, {"l":"^","s":">","r":"v"})}

turningOrder = {"l":"s", "s":"r", "r":"l"}

def parseInput(tracks, carts):
  f = open("day13/input", "r")
  lines = f.read().splitlines()
  for y, line in enumerate(lines):
    row = []
    for x, char in enumerate(line):
      if char in cartParseRules:
        c = cart(char, x, y)
        carts.append(c)
        row.append(cartParseRules[char])
      else:
        row.append(char)
    tracks.append(row)

def run(tracks, carts):
  while len(carts) > 1:
    carts.sort(key=lambda c: (c.y, c.x))
    for c in carts:
      if c.crashed:
        continue
      newX = c.x+dirRules[c.dir].xOffset
      newY = c.y+dirRules[c.dir].yOffset
      if [newX, newY] in ([c.x, c.y] for c in carts if not c.crashed):
        print("BOOM!! @ %d,%d" % (newX, newY))
        otherCart = [c for c in carts if not c.crashed and c.x == newX and c.y == newY][0]
        otherCart.crashed = True
        c.crashed = True
      else: 
        c.x = newX
        c.y = newY
        c.changeDir(tracks[newY][newX])

    carts = [c for c in carts if not c.crashed]

  if len(carts) == 1:
    print("Last cart @ %d,%d" % (carts[0].x, carts[0].y))

carts = []
tracks = []
parseInput(tracks, carts)
run(tracks, carts)
