import re


def parsePotsAndPatters(filepath):
    f = open(filepath, "r")
    lines = f.read().splitlines()
    initialStateRegexp = r"initial state: (.*)"
    initialStatePattern = re.compile(initialStateRegexp)
    initialStateMatch = initialStatePattern.search(lines[0])
    pots = initialStateMatch.group(1)
    ruleRegexp = r"(.*) => #"
    rulePattern = re.compile(ruleRegexp)
    rules = []
    for line in lines[1:]:
        ruleMatch = rulePattern.search(line)
        if ruleMatch is None:
            continue
        rules.append(ruleMatch.group(1))
    return (pots, rules)


def incrementGeneration(pots, rules):
    newPots = "." * len(pots)
    for r in rules:
        pattern = re.compile("(?=" + re.escape(r) + ")")
        for match in pattern.finditer(pots):
            newPots = newPots[: match.start() + 2] + "#" + newPots[match.start() + 3 :]
    return newPots


def sumPotIndices(pots, offset):
    sum = 0
    for i in range(len(pots)):
        if pots[i] == "#":
            sum += i - offset
    return sum


def main():
    pots, rules = parsePotsAndPatters("day12/input")
    startOffset = 10
    pots = "." * 10 + pots + "." * 150
    g = 0
    while True:
        newPots = incrementGeneration(pots, rules)
        if g == 19:
            print("sum #20:", sumPotIndices(newPots, startOffset - g))

        if newPots[1:] + "." == pots:
            print(
                "sum #50000000000:",
                sumPotIndices(newPots, startOffset + 1 - 5000000000),
            )
            break

        pots = newPots
        g += 1
        startOffset += 1


if __name__ == "__main__":
    main()
