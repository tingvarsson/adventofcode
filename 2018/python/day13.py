class cart(object):
  def __init__(self, dir, currTrack):
    self.dir = dir
    self.nextTurn = "left"
    self.currTrack = currTrack
    self.time = 0
    self.x = 0
    self.y = 0

  def changeDir(self, dirRule):
    if self.currTrack == "+":
      self.dir = dirRule.intersection[self.nextTurn]
      self.nextTurn = turningOrder[self.nextTurn]
    else:
      self.dir = dirRule.dirChange[self.currTrack]

class directionRule(object):
  def __init__(self, xOffset, yOffset, dirChange, intersection):
    self.xOffset = xOffset
    self.yOffset = yOffset
    self.dirChange = dirChange
    self.intersection = intersection

parseRules = {"^":"|", ">":"-", "v":"|", "<":"-"}

dirRules = {"^":directionRule(0, -1, {"/":">", "|":"^", "\\":"<", "-":"X"}, {"left":"<","straight":"^","right":">"}),
            "v":directionRule(0,  1, {"/":"<", "|":"v", "\\":">", "-":"X"}, {"left":">","straight":"v","right":"<"}),
            "<":directionRule(-1, 0, {"/":"v", "|":"X", "\\":"^", "-":"<"}, {"left":"v","straight":"<","right":"^"}),
            ">":directionRule( 1, 0, {"/":"^", "|":"X", "\\":"v", "-":">"}, {"left":"^","straight":">","right":"v"})}

turningOrder = {"left":"straight", "straight":"right", "right":"left"}

carts = {}
tracks = []

def parseInput():
  f = open("day13/input", "r")
  lines = f.read().splitlines()
  for line in lines:
    row = []
    for char in line:
      if char in parseRules:
        c = cart(char, parseRules[char])
        carts[str(len(carts))] = c
        row.append(str(len(carts)-1))
      else:
        row.append(char)
    tracks.append(row)
  return tracks

def boomCheck(x, y):
  if tracks[y][x].isdigit():
    cKey = tracks[y][x]
    tracks[y][x] = carts[cKey].currTrack
    del(carts[cKey])
    print("BOOM!! @ %d,%d" % (x, y))
    return True
  return False

def run():
  time = 0
  while True:
    for y, row in enumerate(tracks):
      for x, track in enumerate(row):
        if track.isdigit():
          c = carts[track]
          if c.time >= time:
            continue
          c.time = time
          
          dr = dirRules[c.dir]
          xNext = x+dr.xOffset
          yNext = y+dr.yOffset
          if boomCheck(xNext, yNext):
              tracks[y][x] = carts[track].currTrack
              del(carts[track])
          else: # move to new track and update pos/direction
            tracks[y][x] = c.currTrack
            c.currTrack = tracks[yNext][xNext]
            tracks[yNext][xNext] = track
            c.x = xNext
            c.y = yNext
            c.changeDir(dr)
    
    if len(carts) == 0:
      return
    elif len(carts) == 1:
      print("Last cart @ %d,%d" % (next(iter(carts.values())).x, next(iter(carts.values())).y))
      return
    time += 1

parseInput()
run()
