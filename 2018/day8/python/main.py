class Node(object):
    def __init__(self):
        self.length = 2 # header
        self.children = []
        self.metadata = []

    def addChild(self, child):
        self.children.append(child)
        self.length += child.length

    def addMetadata(self, metadata):
        self.metadata.append(metadata)
        self.length += 1

def parseInput():
    f = open("../input", "r")
    rawInput = f.read().split(' ')
    input = []
    for i in rawInput:
        input.append(int(i))
    return input

def parseNode(data):
    numChildren = data[0]
    numMetadata = data[1]
    node = Node()
    for _ in range(numChildren):
        node.addChild(parseNode(data[node.length:]))
    for i in range(node.length, node.length+numMetadata):
        node.addMetadata(data[i])
    return node

def sumMetadata(node):
    sum = 0
    for c in node.children:
        sum += sumMetadata(c)
    for m in node.metadata:
        sum += m
    return sum

def sumMetadata2(node):
    sum = 0
    if len(node.children) == 0:
        for m in node.metadata:
            sum += m
    else:
        for m in node.metadata:
            if m-1 < len(node.children):
                sum += sumMetadata2(node.children[m-1])
    return sum

data = parseInput()
rootNode = parseNode(data)
print(sumMetadata(rootNode))
print(sumMetadata2(rootNode))