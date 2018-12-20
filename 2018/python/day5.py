import re
import string


def reactPolymer(input):
    output = ""
    for c in input:
        if output and output[-1] != c and output[-1].lower() == c.lower():
            output = output[:-1]
        else:
            output += c
    return output


def main():
    f = open("day5/input", "r")
    input = f.read()

    removeAllPolymer = reactPolymer(input)
    print("Length after reaction:", len(removeAllPolymer))

    removedCharLength = []
    for uc in string.ascii_lowercase:
        newInput = re.sub(uc, "", input, flags=re.IGNORECASE)
        shortenedPolymer = reactPolymer(newInput)
        removedCharLength.append((uc, len(shortenedPolymer)))

    bestChar = min(removedCharLength, key=lambda c: c[1])
    print("Shortest polymer when removing:", bestChar)


if __name__ == "__main__":
    main()
