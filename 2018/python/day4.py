import bisect
import re

# We preprocess the list with a simple cli call to sort
f = open("../sortedinput", "r")
lines = f.read().splitlines()

guardIdRegex = r'#(\d+)'
guardIdPattern = re.compile(guardIdRegex)
asleepRegex = r':(\d+)] falls asleep'
asleepPattern = re.compile(asleepRegex)
wakesRegex = r':(\d+)] wakes up'
wakesPattern = re.compile(wakesRegex)

sleeptime = 60

guards = {}
for line in lines:
    guardIdMatch = guardIdPattern.search(line)
    if guardIdMatch is not None:
        guardId = int(guardIdMatch.group(1))
        if guardId in guards:
            currentSleepRecord = guards[guardId]
        else:
            currentSleepRecord = [0] * sleeptime
            guards[guardId] = currentSleepRecord

    asleepMatch = asleepPattern.search(line)
    if asleepMatch is not None:
        asleepMinute = int(asleepMatch.group(1))

    wakesMatch = wakesPattern.search(line)
    if wakesMatch is not None:
        wakesMinute = int(wakesMatch.group(1))
        for t in range(asleepMinute, wakesMinute):
            currentSleepRecord[t] += 1

scenarioOneGuard = max(guards.items(), key=lambda g: sum(g[1]))
print(scenarioOneGuard[0] * scenarioOneGuard[1].index(max(scenarioOneGuard[1])))

scenarioTwoGuard = max(guards.items(), key=lambda g: max(g[1]))
print(scenarioTwoGuard[0] * scenarioTwoGuard[1].index(max(scenarioTwoGuard[1])))
