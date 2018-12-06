import math
import re

f = open("../input", "r")
lines = f.read().splitlines()

coordRegex = r'(\d+), (\d+)'
coordPattern = re.compile(coordRegex)
coords = [] # [x, y, area]
for line in lines:
    coordMatch = coordPattern.search(line)
    coord = [int(coordMatch.group(1)), int(coordMatch.group(2)), 0]
    coords.append(coord)

maxX = max(coords, key=lambda c: c[0])[0] + 1
maxY = max(coords, key=lambda c: c[1])[1] + 1

scenarioTwoArea = 0
for x in range(maxX):
    for y in range(maxY):
        closestCoord = 0
        shortestDistance = 1000000
        totalDistance = 0
        for i, coord in enumerate(coords):
            distance = abs(x-coord[0]) + abs(y-coord[1])
            if distance < shortestDistance:
                closestCoord = i
                shortestDistance = distance
            elif distance == shortestDistance:
                closestCoord = -1 # tie between several coords
            
            totalDistance += distance
        
        if totalDistance < 10000:
            scenarioTwoArea += 1

        if closestCoord == -1 or coords[closestCoord][2] == -1:
            continue
        if x == maxX-1 or x == 0 or y == maxY-1 or y == 0:
            coords[closestCoord][2] = -1 # infinite area
        else:
            coords[closestCoord][2] += 1

largestAreaCoord = max(coords, key=lambda c: c[2])
print(largestAreaCoord)

print(scenarioTwoArea)