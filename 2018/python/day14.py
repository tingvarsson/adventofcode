input = [3, 7]
elves = [0, 1]
pat = [7,6,5,0,7,1]

while True:
  newValue = input[elves[0]] + input[elves[1]]
  if newValue // 10 != 0: # two digit value
    input.append(newValue // 10) # first digit
    if pat == input[len(input)-len(pat):]:
      break
  input.append(newValue % 10) # second (or only) digit
  if pat == input[len(input)-len(pat):]:
    break

  elves[0] = (elves[0]+1+input[elves[0]]) % len(input)
  elves[1] = (elves[1]+1+input[elves[1]]) % len(input)

print("Ten recipes after sequence: {}".format(input[765071:765071+10]))
print("#recipes before sequence: {}".format(len(input)-len(pat)))
