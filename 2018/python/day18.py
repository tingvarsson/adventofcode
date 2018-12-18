import copy
import hashlib


def parseInput(filepath):
    f = open(filepath, "r")
    data = f.read().splitlines()
    output = []
    for line in data:
        row = []
        for pos in line:
            row.append(pos)
        output.append(row)
    return output


def checkAdjacent(area, x, y):
    trees = 0
    lumberyards = 0
    for x2, y2 in (
        (x - 1, y - 1),
        (x - 1, y),
        (x - 1, y + 1),
        (x, y - 1),
        (x, y + 1),
        (x + 1, y - 1),
        (x + 1, y),
        (x + 1, y + 1),
    ):
        if 0 <= x2 < len(area) and 0 <= y2 < len(area):
            if area[y2][x2] == "|":
                trees += 1
            elif area[y2][x2] == "#":
                lumberyards += 1
    return (trees, lumberyards)


def newResource(area, x, y):
    trees, lumberyards = checkAdjacent(area, x, y)
    if area[y][x] == ".":
        return "|" if trees >= 3 else area[y][x]
    elif area[y][x] == "|":
        return "#" if lumberyards >= 3 else area[y][x]
    elif area[y][x] == "#":
        return "." if (lumberyards == 0 or trees == 0) else area[y][x]


def run(time, area):
    seenAreas = []
    t = 0
    while time is None or t < time:
        newArea = [["" for x in range(len(area[y]))] for y in range(len(area))]
        for y in range(len(newArea)):
            for x in range(len(newArea[y])):
                newArea[y][x] = newResource(area, x, y)

        if newArea in seenAreas:
            firstTime = seenAreas.index(newArea)
            period = t - firstTime
            shortcut = (1000000000 - t - 1) % period
            return seenAreas[firstTime + shortcut]
        else:
            seenAreas.append(newArea)
            area = newArea
            t += 1
    return newArea


def main():
    area = parseInput("day18/input")

    newArea = run(10, area)
    sumWood = sum(sum(1 for pos in row if pos == "|") for row in newArea)
    sumLumberyards = sum(sum(1 for pos in row if pos == "#") for row in newArea)
    print("total resource value @ 10:", sumWood * sumLumberyards)

    newArea = run(None, area)
    sumWood = sum(sum(1 for pos in row if pos == "|") for row in newArea)
    sumLumberyards = sum(sum(1 for pos in row if pos == "#") for row in newArea)
    print("total resource value @ 1000000000:", sumWood * sumLumberyards)


if __name__ == "__main__":
    main()
