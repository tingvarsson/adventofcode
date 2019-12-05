import bisect
import utils


def main():
    lines = utils.readlines("day1/input")
    numbers = [int(l) for l in lines]

    print("Sum of numbers:", sum(numbers))

    currentSum = 0
    i = 0
    knownSums = []
    while True:
        currentSum += numbers[i % len(numbers)]
        pos = bisect.bisect_left(knownSums, currentSum)
        if pos != len(knownSums) and knownSums[pos] == currentSum:
            break
        else:
            bisect.insort(knownSums, currentSum)
            i += 1
    print("First reoccuring sum:", currentSum)


if __name__ == "__main__":
    main()
