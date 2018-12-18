import copy
import hashlib


def parseInput(filepath):
    f = open(filepath, "r")
    data = f.read().splitlines()
    output = [[" "] * len(data[0])]
    for line in data:
        row = [" "]
        for pos in line:
            row.append(pos)
        output.append(row)
        row.append(" ")
    output.append([" "] * len(data[0]))
    return output


def checkAdjacentTrees(area, toCheck):
    trees = 0
    for x, y in toCheck:
        if area[y][x] == "|":
            trees += 1
        if trees >= 3:
            return True
    return False


def checkAdjacentLumberyards(area, toCheck):
    lumberyards = 0
    for x, y in toCheck:
        if area[y][x] == "#":
            lumberyards += 1
        if lumberyards >= 3:
            return True
    return False


def checkAnyAdjacentTreesAndLumberyards(area, toCheck):
    trees = 0
    lumberyards = 0
    for x, y in toCheck:
        if area[y][x] == "|":
            trees += 1
        elif area[y][x] == "#":
            lumberyards += 1
        if trees >= 1 and lumberyards >= 1:
            return True
    return False


def newResource(area, x, y):
    adjacent = (
        (x - 1, y - 1),
        (x - 1, y),
        (x - 1, y + 1),
        (x, y - 1),
        (x, y + 1),
        (x + 1, y - 1),
        (x + 1, y),
        (x + 1, y + 1),
    )
    if area[y][x] == ".":
        return "|" if checkAdjacentTrees(area, adjacent) else area[y][x]
    elif area[y][x] == "|":
        return "#" if checkAdjacentLumberyards(area, adjacent) else area[y][x]
    elif area[y][x] == "#":
        return (
            area[y][x] if checkAnyAdjacentTreesAndLumberyards(area, adjacent) else "."
        )


def run(time, area):
    seenAreas = []
    t = 0
    while time is None or t < time:
        newArea = [[" " for x in range(AREASIZE)] for y in range(AREASIZE)]
        for y in range(1, AREASIZE - 1):
            for x in range(1, AREASIZE - 1):
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


AREASIZE = 0


def main():
    area = parseInput("day18/input")
    global AREASIZE
    AREASIZE = len(area) - 2

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
