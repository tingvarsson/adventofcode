f = open("day2/input", "r")
fl = f.read().splitlines()

sum2 = 0
sum3 = 0
for line in fl:
    letters = {}
    for c in line:
        if c in letters:
            letters[c] += 1
        else:
            letters[c] = 1
    
    found2 = False
    found3 = False
    for c in letters:
        if letters[c] == 2:
            found2 = True
        elif letters[c] == 3:
            found3 = True
    
    if found2:
        sum2 += 1
    if found3:
        sum3 += 1

print("checksum:")
print(sum2*sum3)

def strIntersection(s1, s2):
    out = ""
    for i in range(0, len(s1)):
        if s1[i] == s2[i]:
            out += s1[i]
    return out

for line in fl:
    for secondline in fl:
        if len(line)-1 == len(strIntersection(line, secondline)):
            print("same line:")
            print(strIntersection(line, secondline))
            exit(0)
