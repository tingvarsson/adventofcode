import bisect

f = open("day1/input", "r")
fl = f.readlines()
numbers = []
for l in fl:
    numbers.append(int(l))

sum = 0
for n in numbers:
    sum += n
print(sum)

sum = 0
i = 0
knownsums = [0]
while True:
    sum += numbers[i % len(numbers)]
    pos = bisect.bisect_left(knownsums, sum)
    if pos != len(knownsums) and knownsums[pos] == sum:
        break
    else:
        bisect.insort(knownsums, sum)
        i += 1
print(sum)