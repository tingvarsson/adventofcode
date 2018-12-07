import math
import re

def parseInput():
    f = open("../input", "r")
    lines = f.read().splitlines()
    regex = r'Step (.) .* step (.)'
    pattern = re.compile(regex)
    for line in lines:
        match = pattern.search(line)
        parent = match.group(1)
        child = match.group(2)
        if parent in parentChildrenGraph:
            parentChildrenGraph[parent].append(child)
        else:
            parentChildrenGraph[parent] = [child]
        
        if child in childrenParentGraph:
            childrenParentGraph[child].append(parent)
        else:
            childrenParentGraph[child] = [parent]

def allParentsInOrder(order, parents):
    foundAll = True
    for parent in parents:
        if parent not in order:
            foundAll = False
    return foundAll

def addAvailable(parent, order, available):
    if parent not in parentChildrenGraph:
        return available
    for child in parentChildrenGraph[parent]:
        if allParentsInOrder(order, childrenParentGraph[child]):
            available.append(child)
    available = sorted(available)

# Bi-directional relationship of all parents and children
parentChildrenGraph = {} # [parent, [children]]
childrenParentGraph = {} # [child, [parents]]
parseInput()

rootParent = None
for parent in parentChildrenGraph:
    if parent not in childrenParentGraph:
        rootParent = parent

order = rootParent
available = []
while True:
    addAvailable(order[-1], order, available)
    if len(available) == 0:
        break
    order += available[0]
    available = available[1:]

print(order)

workers = [("", 0)] * 5 # 5 workers á (currentTask, whenAvailableAgain)
available = [rootParent]
time = 0
order = "" # keep track of processed instructions
while True:
    for i, w in enumerate(workers): # Free workers and add new availables
        if w[0] != "" and w[1] < time:
            order += w[0]
            addAvailable(w[0], order, available)
            workers[i] = ("", 0) # Worker freed up

    removeAvailable = []
    for a in available: # Assign new work if available
        for j, w in enumerate(workers):
            if w[0] == "":
                workers[j] = (a, time+ord(a)-ord("A")+60)
                removeAvailable.append(a)
                break # Work assigned, continue with next available
    
    for remove in removeAvailable:
        i = available.index(remove)
        available = available[:i]+available[i+1:]

    workOngoing = False
    for w in workers: # Determine if we are finished yet or not
        if w[0] != "":
            workOngoing = True
            break
    
    if workOngoing:
        time += 1
    else:
        break

print(time)