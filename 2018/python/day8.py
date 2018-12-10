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
    f = open("day8/input", "r")
    input = f.read().split(' ')
    return [int(i) for i in input]

def parseNode(data):
    node = Node()
    for _ in range(data[0]):
        node.addChild(parseNode(data[node.length:]))
    for _ in range(data[1]):
        node.addMetadata(data[node.length])
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
            if m <= len(node.children):
                sum += sumMetadata2(node.children[m-1])
    return sum

data = parseInput()
rootNode = parseNode(data)
print(sumMetadata(rootNode))
print(sumMetadata2(rootNode))