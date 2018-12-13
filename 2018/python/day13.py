class cart(object):
  def __init__(self, id, dir, x, y, currTrack):
    self.id = id
    self.dir = dir
    self.nextTurn = "left"
    self.currTrack = currTrack
    self.x = x
    self.y = y
    self.crashed = False

  def changeDir(self):
    dr = dirRules[self.dir]
    if self.currTrack == "+":
      self.dir = dr.intersection[self.nextTurn]
      self.nextTurn = turningOrder[self.nextTurn]
    else:
      self.dir = dr.dirChange[self.currTrack]

class directionRule(object):
  def __init__(self, xOffset, yOffset, dirChange, intersection):
    self.xOffset = xOffset
    self.yOffset = yOffset
    self.dirChange = dirChange
    self.intersection = intersection

cartParseRules = {"^":"|", ">":"-", "v":"|", "<":"-"}

dirRules = {"^":directionRule(0, -1, {"/":">", "|":"^", "\\":"<", "-":"X"}, {"left":"<","straight":"^","right":">"}),
            "v":directionRule(0,  1, {"/":"<", "|":"v", "\\":">", "-":"X"}, {"left":">","straight":"v","right":"<"}),
            "<":directionRule(-1, 0, {"/":"v", "|":"X", "\\":"^", "-":"<"}, {"left":"v","straight":"<","right":"^"}),
            ">":directionRule( 1, 0, {"/":"^", "|":"X", "\\":"v", "-":">"}, {"left":"^","straight":">","right":"v"})}

turningOrder = {"left":"straight", "straight":"right", "right":"left"}

def parseInput(tracks, carts):
  f = open("day13/input", "r")
  lines = f.read().splitlines()
  for y, line in enumerate(lines):
    row = []
    for x, char in enumerate(line):
      if char in cartParseRules:
        c = cart(len(carts), char, x, y, cartParseRules[char])
        carts.append(c)
        row.append(str(c.id))
      else:
        row.append(char)
    tracks.append(row)
  return tracks

def run(tracks, carts):
  while len(carts) > 1:
    carts.sort(key=lambda c: (c.y, c.x))
    for c in carts:
      if c.crashed:
        continue
      newX = c.x+dirRules[c.dir].xOffset
      newY = c.y+dirRules[c.dir].yOffset
      if [newX, newY] in ([c.x, c.y] for c in carts if not c.crashed):
        otherCart = [c for c in carts if not c.crashed and c.x == newX and c.y == newY][0]
        print("BOOM!! @ %d,%d" % (newX, newY))
        tracks[otherCart.y][otherCart.x] = otherCart.currTrack
        otherCart.crashed = True
        tracks[c.y][c.x] = c.currTrack
        c.crashed = True
      else: 
        tracks[c.y][c.x] = c.currTrack
        c.currTrack = tracks[newY][newX]
        tracks[newY][newX] = str(c.id)
        c.x = newX
        c.y = newY
        c.changeDir()

    carts = [c for c in carts if not c.crashed]

  if len(carts) == 1:
    print("Last cart @ %d,%d" % (carts[0].x, carts[0].y))

carts = []
tracks = []
parseInput(tracks, carts)
run(tracks, carts)
