import utils


def strIntersection(s1, s2):
    out = ""
    for i in range(len(s1)):
        if s1[i] == s2[i]:
            out += s1[i]
    return out


def main():
    lines = utils.readlines("day2/input")

    sum2 = 0
    sum3 = 0
    for line in lines:
        letters = {}
        for c in line:
            if c in letters:
                letters[c] += 1
            else:
                letters[c] = 1

        found2 = False
        found3 = False
        for c in letters:
            if not found2 and letters[c] == 2:
                sum2 += 1
                found2 = True
            elif not found3 and letters[c] == 3:
                sum3 += 1
                found3 = True

    print("Checksum:", sum2 * sum3)

    for line in lines:
        for secondline in lines:
            if len(line) - 1 == len(strIntersection(line, secondline)):
                print("Common letters:", strIntersection(line, secondline))
                return


if __name__ == "__main__":
    main()
