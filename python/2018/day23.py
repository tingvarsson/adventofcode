import re


class nanobot(object):
    def __init__(self, x, y, z, r):
        self.x = x
        self.y = y
        self.z = z
        self.r = r

    def __str__(self):
        return "x: %d, y: %d, z: %d, r: %d" % (self.x, self.y, self.z, self.r)

    def distance(self, x, y, z):
        return abs(self.x - x) + abs(self.y - y) + abs(self.z - z)

    def inRange(self, other):
        return self.distance(other.x, other.y, other.z) <= self.r

    def inRangeCoord(self, x, y, z):
        return self.distance(x, y, z) <= self.r

    def intersecting(self, other):
        return self.distance(other.x, other.y, other.z) <= self.r + other.r


def parseNanobots(filepath):
    f = open(filepath, "r")
    lines = f.read().splitlines()
    bots = []
    for l in lines:
        m = re.search(r"pos=<(.*),(.*),(.*)>, r=(.*)", l)
        bots.append(nanobot(*[int(i) for i in m.groups()]))
    return bots


def main():
    bots = parseNanobots("day23/input")
    topBot = max(bots, key=lambda b: b.r)
    inRangeBots = sum(1 for b in bots if topBot.inRange(b))
    print(inRangeBots)

    xbound = min(bots, key=lambda p: p.x).x, max(bots, key=lambda p: p.x).x
    ybound = min(bots, key=lambda p: p.y).y, max(bots, key=lambda p: p.y).y
    zbound = min(bots, key=lambda p: p.z).z, max(bots, key=lambda p: p.z).z

    step = 1
    while step < (xbound[1] - xbound[0]):
        step *= 2

    bestCount = 0
    bestDistance = 0
    bestCoord = (0, 0, 0)
    while True:
        for x in range(xbound[0], xbound[1] + 1, step):
            for y in range(ybound[0], ybound[1] + 1, step):
                for z in range(zbound[0], zbound[1] + 1, step):
                    countInRange = 0
                    for b in bots:
                        countInRange += 1 if b.inRangeCoord(x, y, z) else 0
                    dist = abs(x) + abs(y) + abs(z)
                    if countInRange > bestCount:
                        bestCount = countInRange
                        bestDistance = dist
                        bestCoord = (x, y, z)
                    elif countInRange == bestCount and dist < bestDistance:
                        bestDistance = dist
                        bestCoord = (x, y, z)

        if step == 1:
            print(bestCoord, bestDistance)
            break
        else:
            xbound = (bestCoord[0] - step, bestCoord[0] + step)
            ybound = (bestCoord[1] - step, bestCoord[1] + step)
            zbound = (bestCoord[2] - step, bestCoord[2] + step)
            step //= 2


if __name__ == "__main__":
    main()
