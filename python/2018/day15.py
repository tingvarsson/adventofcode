import collections
import copy

class unit(object):
  damage = 3
  def __init__(self, x, y, symbol):
    self.x = x
    self.y = y
    self.symbol = symbol
    self.hp = 200

class elf(unit):
  damage = 3
  def __init__(self, x, y):
    unit.__init__(self, x, y, "E")

class goblin(unit):
  def __init__(self, x, y):
    unit.__init__(self, x, y, "G")

def parseInput(area, units):
  f = open("day15/input", "r")
  lines = f.read().splitlines()
  for y, line in enumerate(lines):
    row = []
    for x, pos in enumerate(line):
      if pos == "E":
        units.append(elf(x,y))
      elif pos == "G":
        units.append(goblin(x,y))
      row.append(pos)
    area.append(row)

def attackAdjecentEnemy(self, enemies, area):
  adjecentEnemies = [e for e in enemies if abs(e.x - self.x) + abs(e.y - self.y) == 1]
  if len(adjecentEnemies) == 0:
    return False

  adjecentEnemies.sort(key=lambda e: (e.hp, e.y, e.x))
  adjecentEnemies[0].hp -= self.damage
  if adjecentEnemies[0].hp <= 0:
    area[adjecentEnemies[0].y][adjecentEnemies[0].x] = "."
  return True

def adjecentCoords(x, y):
  return ((x, y-1), (x-1, y), (x+1, y), (x, y+1))

def search(area, start, goal):
  queue = collections.deque([[start]])
  seen = set([start])
  while queue:
    path = queue.popleft()
    x, y = path[-1]
    for x2, y2 in adjecentCoords(x, y):
      if (x2, y2) == goal:
        return path + [(x2, y2)]
      if 0 <= x2 < len(area) and 0 <= y2 < len(area) and area[y2][x2] == "." and (x2, y2) not in seen:
        queue.append(path + [(x2, y2)])
        seen.add((x2, y2))

def run(area, units):
  time = 0
  while True:
    units = [u for u in units if u.hp > 0]
    units.sort(key=lambda u: (u.y, u.x))
    for u in units:
      if u.hp <= 0:
        continue
      enemies = [e for e in units if e.hp > 0 and type(e) is not type(u)]
      if len(enemies) == 0:
        return (time, type(u))

      if not attackAdjecentEnemy(u, enemies, area):
        attackVectors = []
        for e in enemies:
          for x, y in adjecentCoords(e.x, e.y):
            if area[y][x] == ".":
              attackVectors.append((x, y))
        if len(attackVectors) == 0:
          continue
        paths = []
        for av in attackVectors:
          shortestPath = search(area, (u.x, u.y), av)
          if shortestPath != None:
            paths.append(shortestPath)
        if len(paths) == 0:
          continue
        shortest = min(paths, key=lambda p: len(p))
        nextStepShortestPath = [p[1] for p in paths if len(p) == len(shortest)]
        nextStepShortestPath.sort(key=lambda p: (p[1], p[0]))
        step = nextStepShortestPath[0]
        area[u.y][u.x] = "."
        u.x, u.y = step
        area[u.y][u.x] = u.symbol
        attackAdjecentEnemy(u, enemies, area)
    time += 1

origArea = []
origUnits = []
parseInput(origArea, origUnits)

while True:
  area = copy.deepcopy(origArea)
  units = copy.deepcopy(origUnits)
  print("Experiment Elf damage: %d" % elf.damage)
  time, winner = run(area, units)
  print("Time: %d" % time)
  print("Winner: %s" % winner)
  sumAlive = sum(u.hp for u in units if u.hp > 0)
  print("Winner sum HP: %d" % sumAlive)
  print("Magic number: %d" % (time*sumAlive))
  numElves = sum(1 for u in units if type(u) == elf)
  numElvesAlive = sum(1 for u in units if type(u) == elf and u.hp > 0)
  print("Elves alive: %d/%d" % (numElvesAlive, numElves))
  if numElves == numElvesAlive:
    break
  elf.damage += 1