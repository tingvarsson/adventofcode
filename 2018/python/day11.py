def calcFuelCell(x, y, gridSerialNumber):
    rackID = x + 10
    powerLevelStart = (rackID * y)
    addedSerial = powerLevelStart + gridSerialNumber
    multiplyRackID = addedSerial * rackID
    reduced = (multiplyRackID // 100) % 10
    return reduced - 5

def populateGrid(gridSerialNumber):
    for y in range(1, gridSize-1):
        for x in range(1, gridSize-1):
            grid[y][x] = calcFuelCell(x, y, gridSerialNumber)

def sumMinorSquareAddition(xStart, yStart, size):
    sum = sumGrid[yStart][xStart]
    for i in range(size):
        sum += grid[yStart+i][xStart+size-1]
    for i in range(size):
        sum += grid[yStart+size-1][xStart+i]
    sumGrid[yStart][xStart] = sum
    return sum

def findLargest(minSize, maxSize):
    largest = (0, 0, 0 ,0) # sum, x, y, size
    for size in range(1, maxSize+1):
        for y in range(1, gridSize-(size-2)):
            for x in range(1, gridSize-(size-2)):
                newSum = sumMinorSquareAddition(x, y, size)
                if size >= minSize and newSum > largest[0]:
                    largest = (newSum, x, y, size)
    return largest

gridSize = 300
grid = [[0 for x in range(gridSize+1)] for y in range(gridSize+1)]
sumGrid = [[0 for x in range(gridSize+1)] for y in range(gridSize+1)]
populateGrid(6548)
print(findLargest(3, 3))
print(findLargest(3, 300))
