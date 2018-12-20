import re
import utils


def createClaim(line):
    #                id      x     y      xsize ysize
    m = re.match(r"#(\d+) @ (\d+),(\d+): (\d+)x(\d+)", line)
    claim = {
        "id": int(m.group(1)),
        "x": int(m.group(2)),
        "y": int(m.group(3)),
        "xsize": int(m.group(4)),
        "ysize": int(m.group(5)),
    }
    return claim


def addClaim(fabric, claim):
    for y in range(claim["y"], claim["y"] + claim["ysize"]):
        for x in range(claim["x"], claim["x"] + claim["xsize"]):
            fabric[y][x] += 1


def testClaim(fabric, claim):
    for y in range(claim["y"], claim["y"] + claim["ysize"]):
        for x in range(claim["x"], claim["x"] + claim["xsize"]):
            if fabric[y][x] >= 2:
                return False
    return True


def main():
    lines = utils.readlines("day3/input")
    claims = [createClaim(l) for l in lines]

    inches = 1000
    fabric = [[0 for x in range(inches)] for y in range(inches)]
    for claim in claims:
        addClaim(fabric, claim)

    sumMultiple = sum(sum(1 for x in y if x >= 2) for y in fabric)
    print("Square inches with multiple claims:", sumMultiple)

    for claim in claims:
        if testClaim(fabric, claim):
            print("Non-overlapping claim:", claim)


if __name__ == "__main__":
    main()
