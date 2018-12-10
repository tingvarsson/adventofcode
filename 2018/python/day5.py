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

f = open("day5/input", "r")
lines = f.read().splitlines()
input = lines[0]

removeAllPolymer = reactPolymer(input)
print(len(removeAllPolymer))

removedCharLength = {}
for uc in string.ascii_lowercase:
    newInput = re.sub(uc, '', input)
    newInput = re.sub(uc.upper(), '', newInput)
    shortenedPolymer = reactPolymer(newInput)
    removedCharLength[uc] = len(shortenedPolymer)

bestChar = min(removedCharLength.items(), key=lambda c: c[1])
print(bestChar)
