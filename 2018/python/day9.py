class Node(object):
    def __init__(self, v):
        self.value = v
        self.next = None
        self.prev = None

totalPlayers = 418
lastMarble = 70769
superLastMarble = lastMarble*100

def play(end):
    currPlayer = 0
    score = [0]*totalPlayers
    currMarble = Node(0)
    currMarble.next = currMarble
    currMarble.prev = currMarble
    for marble in range(1, end+1):
        if marble % 23 != 0:
            newMarble = Node(marble)
            currMarble = currMarble.next
            newMarble.prev = currMarble
            newMarble.next = currMarble.next
            currMarble.next.prev = newMarble
            currMarble.next = newMarble
            currMarble = newMarble
        else:
            score[currPlayer] += marble
            currMarble = currMarble.prev.prev.prev.prev.prev.prev.prev
            score[currPlayer] += currMarble.value
            currMarble = currMarble.next
            currMarble.prev.prev.next = currMarble
            currMarble.prev = currMarble.prev.prev
        currPlayer = (currPlayer + 1) % totalPlayers
    return max(score)

print("high score: %d" % play(lastMarble))
print("high score: %d" % play(superLastMarble))