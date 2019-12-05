import bisect
import re
import utils


def main():
    # We preprocess the list with a simple cli call to sort
    lines = utils.readlines("day4/sortedinput")

    guards = {}
    for line in lines:
        guardIdMatch = re.search(r"#(\d+)", line)
        if guardIdMatch is not None:
            guardId = int(guardIdMatch.group(1))
            if guardId in guards:
                currentSleepRecord = guards[guardId]
            else:
                currentSleepRecord = [0] * 60
                guards[guardId] = currentSleepRecord

        asleepMatch = re.search(r":(\d+)] falls asleep", line)
        if asleepMatch is not None:
            asleepMinute = int(asleepMatch.group(1))

        awakeMatch = re.search(r":(\d+)] wakes up", line)
        if awakeMatch is not None:
            awakeMinute = int(awakeMatch.group(1))
            for t in range(asleepMinute, awakeMinute):
                currentSleepRecord[t] += 1

    scenarioOneGuard = max(guards.items(), key=lambda g: sum(g[1]))
    print(scenarioOneGuard[0] * scenarioOneGuard[1].index(max(scenarioOneGuard[1])))

    scenarioTwoGuard = max(guards.items(), key=lambda g: max(g[1]))
    print(scenarioTwoGuard[0] * scenarioTwoGuard[1].index(max(scenarioTwoGuard[1])))


if __name__ == "__main__":
    main()
