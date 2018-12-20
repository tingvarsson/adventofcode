from collections import deque


def addDoorAndRoom(area, x, y, c):
    if c == "N":
        y -= 1
        area[y][x - 1] = "#"
        area[y][x] = "-"
        area[y][x + 1] = "#"
        y -= 1
        area[y][x] = "."
    elif c == "S":
        y += 1
        area[y][x - 1] = "#"
        area[y][x] = "-"
        area[y][x + 1] = "#"
        y += 1
        area[y][x] = "."
    elif c == "W":
        x -= 1
        area[y - 1][x] = "#"
        area[y][x] = "|"
        area[y + 1][x] = "#"
        x -= 1
        area[y][x] = "."
    elif c == "E":
        x += 1
        area[y - 1][x] = "#"
        area[y][x] = "|"
        area[y + 1][x] = "#"
        x += 1
        area[y][x] = "."
    return (x, y)


def readToMap(filepath):
    f = open(filepath, "r")
    input = f.read()
    maxY = min(input.count("N") + input.count("S") + 5, 300)
    maxX = min(input.count("W") + input.count("E") + 5, 300)
    offsetY = maxY // 2
    offsetX = maxX // 2
    area = [[" " for x in range(maxX)] for y in range(maxY)]
    x = offsetX
    y = offsetY
    area[y][x] = "X"
    stack = deque()  # remember pos of all opening parentesis
    for c in input:
        if c == "(":
            stack.append((x, y))
        elif c == ")":
            stack.pop()
        elif c == "|":
            x, y = stack[-1]  # jump back to start pos of the last parentesis
        else:
            x, y = addDoorAndRoom(area, x, y, c)
    return (area, offsetX, offsetY)


def search(area, start):
    queue = deque([[start]])
    seen = set([start])
    paths = []
    while queue:
        path = queue.popleft()
        x, y = path[-1]
        for x2, y2 in ((x - 2, y), (x + 2, y), (x, y - 2), (x, y + 2)):
            if (
                0 <= x2 < len(area)
                and 0 <= y2 < len(area)
                and (
                    area[(y + y2) // 2][(x + x2) // 2] == "|"
                    or area[(y + y2) // 2][(x + x2) // 2] == "-"
                )
                and area[y2][x2] == "."
                and (x2, y2) not in seen
            ):
                queue.append(path + [(x2, y2)])
                paths.append(path + [(x2, y2)])
                seen.add((x2, y2))
    return paths


def main():
    area, offsetX, offsetY = readToMap("day20/input")
    paths = search(area, (offsetX, offsetY))
    print("Furthest:", max([len(p) - 1 for p in paths]))
    print("Rooms at least 1000 doors away:", sum(1 for p in paths if len(p) > 1000))


if __name__ == "__main__":
    main()
