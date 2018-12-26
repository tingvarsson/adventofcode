import re

attackTypes = {}


class group(object):
    def __init__(self, units, hp, abilities, damage, attackType, initiative):
        self.units = int(units)
        self.hp = int(hp)
        self.immune, self.weak = parseAbilities(abilities)
        self.damage = int(damage)
        if attackType not in attackTypes:
            attackTypes[attackType] = len(attackTypes)
        self.attackType = attackTypes[attackType]
        self.initiative = int(initiative)
        self.enemy = None
    def __str__(self):
        return "{}: {}, {}, {}, {}, {}, {}, {}, {}".format(type(self).__name__, self.effectivePwr(), self.units, self.hp, self.immune, self.weak, self.damage, self.attackType, self.initiative)
    def effectivePwr(self):
        return self.units*self.damage
    def isEnemy(self, other):
        return self.enemy == type(other)
    def calcDamage(self, other):
        if self.attackType in other.immune:
            return 0
        dmg = self.effectivePwr()
        if self.attackType in other.weak:
            dmg *= 2
        return dmg
    def attack(self, other):
        dmg = self.calcDamage(other)
        other.units -= min(other.units, dmg // other.hp)

class immuneSystem(group):
    def __init__(self, units, hp, abilities, damage, attackType, initiative):
        group.__init__(self, units, hp, abilities, damage, attackType, initiative)
        self.enemy = infection
        #self.damage += 43
        
class infection(group):
    def __init__(self, units, hp, abilities, damage, attackType, initiative):
        group.__init__(self, units, hp, abilities, damage, attackType, initiative)
        self.enemy = immuneSystem

def parseArmies(filepath):
    f = open(filepath, "r")
    inputImmune, inputInfection = f.read().split("Infection:")
    armies = []
    for l in inputImmune.splitlines():
        m = re.search(r"(\d+) .* (\d+) hit points (.*)with.* (\d+) ([a-zA-Z]+\w).* (\d+)", l)
        if m is not None:
            armies.append(immuneSystem(*m.groups()))
    for l in inputInfection.splitlines():
        m = re.search(r"(\d+) .* (\d+) hit points (.*)with.* (\d+) ([a-zA-Z]+\w).* (\d+)", l)
        if m is not None:
            armies.append(infection(*m.groups()))
    return armies


def parseAbilities(input):
    if input is None:
        return ([], [])
    input = input.replace("(", "").replace(") ", "").replace("to", "").replace(" ", ",").replace(",,", ",").replace(";", "").split(",")
    immune = []
    weak = []
    current = weak
    for i in input:
        if i == "immune":
            current = immune
            continue
        elif i == "weak":
            current = weak
            continue
        elif i == "":
            continue
        
        if i not in attackTypes:
            attackTypes[i] = len(attackTypes)
        current.append(attackTypes[i])
    return (immune, weak)


def run(armies):
    while True:
        armies = [a for a in armies if a.units > 0]
        armies.sort(key=lambda a: (a.effectivePwr(), a.initiative), reverse=True)
        attacking = {}
        targets = {}
        for a in armies:
            enemies = [o for o in armies if a.isEnemy(o) and o not in targets and a.calcDamage(o)]
            enemies.sort(key=lambda e: (a.calcDamage(e), e.effectivePwr(), e.initiative), reverse=True)
            if enemies:
                attacking[a] = enemies[0]
                targets[enemies[0]] = a
        if not attacking:
            break
        armies.sort(key=lambda a: a.initiative, reverse=True)
        for a in armies:
            if a in attacking:
                a.attack(attacking[a])

    print("Sum of surviving team:", sum(a.units for a in armies))
    for a in armies:
        print(a)


def main():
    armies = parseArmies("day24/input")
    run(armies)


if __name__ == "__main__":
    main()