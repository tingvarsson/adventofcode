from collections import deque

TARGET = (14, 778, 1)
REGION_BUFFER = 50


def calculateErosion(x, y, erosions):
    depth = 11541
    magic = 20183
    magicX = 48271
    magicY = 16807
    if x == TARGET[0] and y == TARGET[1]:
        erosion = depth % magic
    elif x == 0 and y == 0:
        erosion = depth % magic
    elif x == 0:
        erosion = ((y * magicX) + depth) % magic
    elif y == 0:
        erosion = ((x * magicY) + depth) % magic
    else:
        erosion = ((erosions[y - 1][x] * erosions[y][x - 1]) + depth) % magic
    return erosion


def createMap():
    erosions = []
    regionMap = []
    for y in range(TARGET[1] + REGION_BUFFER):
        erosionRow = []
        erosions.append(erosionRow)
        typeRow = []
        regionMap.append(typeRow)
        for x in range(TARGET[0] + REGION_BUFFER):
            erosion = calculateErosion(x, y, erosions)
            erosionRow.append(erosion)
            typeRow.append(erosion % 3)
    return regionMap


def search(area, start, goal):
    queue = deque([[start]])
    costs = {}
    paths = []
    while queue:
        path = queue.popleft()
        x, y, tool, cost = path[-1]
        for x2, y2 in ((x - 1, y), (x + 1, y), (x, y - 1), (x, y + 1)):
            if (x2, y2, tool) == goal:
                paths.append(path + [(x2, y2, tool, cost + 1)])
            if (
                0 <= x2 < TARGET[0] + REGION_BUFFER
                and 0 <= y2 < TARGET[1] + REGION_BUFFER
                and (
                    (area[y2][x2] == 0 and tool != 0)  # rocky, not neither
                    or (area[y2][x2] == 1 and tool != 1)  # wet, not torch
                    or (area[y2][x2] == 2 and tool != 2)  # narrow, not climbing gear
                )
                and ((x2, y2, tool) not in costs or cost + 1 < costs[(x2, y2, tool)])
            ):
                costs[(x2, y2, tool)] = cost + 1
                queue.append(path + [(x2, y2, tool, cost + 1)])

        for tool2 in ((tool + 1) % 3, (tool + 2) % 3):
            if (x, y, tool2) == goal:
                paths.append(path + [(x, y, tool2, cost + 7)])

            if (
                (area[y][x] == 0 and tool2 != 0)  # rocky, not neither
                or (area[y][x] == 1 and tool2 != 1)  # wet, not torch
                or (area[y][x] == 2 and tool2 != 2)  # narrow, not climbing gear
            ) and ((x, y, tool2) not in costs or cost + 7 < costs[(x, y, tool2)]):
                costs[(x, y, tool2)] = cost + 7
                queue.append(path + [(x, y, tool2, cost + 7)])

    return paths


def main():
    regionMap = createMap()
    print(
        "Total risk level:",
        sum(sum(row[: TARGET[0] + 1]) for row in regionMap[: TARGET[1] + 1]),
    )
    print(
        "Fewest minutes to target:",
        min(search(regionMap, (0, 0, 1, 0), TARGET), key=lambda p: p[-1][-1])[-1][-1],
    )


if __name__ == "__main__":
    main()
