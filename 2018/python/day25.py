import re


def parseCoords(filepath):
    f = open(filepath, "r")
    lines = f.read().splitlines()

    coords = []
    for l in lines:
        m = re.search(r"(-?\d+),(-?\d+),(-?\d+),(-?\d+)", l)
        coords.append([int(i) for i in m.groups()])
    return coords


def dist(a, b):
    return abs(a[0] - b[0]) + abs(a[1] - b[1]) + abs(a[2] - b[2]) + abs(a[3] - b[3])


def main():
    coords = parseCoords("day25/input")
    constellations = []
    for c in coords:
        found = []
        for co in constellations:
            if any(dist(c, coord) <= 3 for coord in co):
                found.append(co)
        if found:
            newCo = [c]
            for f in found:
                newCo += f
                if f in constellations:
                    constellations.remove(f)
            constellations.append(newCo)
        else:
            constellations.append([c])
    print(len(constellations))


if __name__ == "__main__":
    main()

