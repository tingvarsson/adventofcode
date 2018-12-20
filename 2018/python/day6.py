import re
import utils


def main():
    lines = utils.readlines("day6/input")
    coords = [[int(m) for m in re.search(r"(\d+), (\d+)", l).groups()] for l in lines]

    maxX = max(coords, key=lambda c: c[0])[0] + 1
    maxY = max(coords, key=lambda c: c[1])[1] + 1
    coordAreas = [0] * len(coords)
    scenarioTwoArea = 0
    for x in range(maxX):
        for y in range(maxY):
            closestCoord = 0
            shortestDistance = None
            totalDistance = 0
            for i, coord in enumerate(coords):
                distance = abs(x - coord[0]) + abs(y - coord[1])
                if shortestDistance is None or distance < shortestDistance:
                    closestCoord = i
                    shortestDistance = distance
                elif distance == shortestDistance:
                    closestCoord = -1  # tie between several coords

                totalDistance += distance

            if totalDistance < 10000:
                scenarioTwoArea += 1

            if closestCoord == -1 or coordAreas[closestCoord] == -1:
                continue

            if 0 < x < maxX - 1 and 0 < y < maxY - 1:
                coordAreas[closestCoord] += 1
            else:
                coordAreas[closestCoord] = -1  # infinite area

    print("Largest area around one coord:", max(coordAreas))
    print("Area with a distance of less than 10000 to all coords:", scenarioTwoArea)


if __name__ == "__main__":
    main()
