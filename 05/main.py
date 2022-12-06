def readFile(file):
    file = open(file, "r")
    return file.readlines()


def find(s, ch):
    return [i for i, ltr in enumerate(s) if ltr == ch]


def firstTask(fileContent):
    columnCount = len(fileContent[0]) // 4
    columns = [[] for i in range(columnCount)]
    i = 0
    while len(fileContent[i]) != 1:
        indices = find(fileContent[i], "[")
        for index in indices:
            currentItem = fileContent[i][index : index + 3]
            columns[index // 4].insert(0, currentItem)
        i += 1
    i += 1
    for row in fileContent[i:]:
        currentInstruction = row.split()
        moveSize = int(currentInstruction[1])
        startStack = int(currentInstruction[3]) - 1
        endStack = int(currentInstruction[5]) - 1
        # print(moveSize, startStack, endStack)
        for i in range(moveSize):
            poped = columns[startStack].pop()
            columns[endStack].append(poped)
    endChars = [stack[-1].replace("[", "").replace("]", "") for stack in columns]
    return "".join(endChars)


def secondTask(fileContent):
    columnCount = len(fileContent[0]) // 4
    columns = [[] for i in range(columnCount)]
    i = 0
    while len(fileContent[i]) != 1:
        indices = find(fileContent[i], "[")
        for index in indices:
            currentItem = fileContent[i][index : index + 3]
            columns[index // 4].insert(0, currentItem)
        i += 1
    i += 1
    for row in fileContent[i:]:
        currentInstruction = row.split()
        moveSize = int(currentInstruction[1])
        startStack = int(currentInstruction[3]) - 1
        endStack = int(currentInstruction[5]) - 1
        # print(moveSize, startStack, endStack)
        poped = [columns[startStack].pop() for i in range(moveSize)]
        for po in poped[::-1]:
            columns[endStack].append(po)
    endChars = [stack[-1].replace("[", "").replace("]", "") for stack in columns]
    return "".join(endChars)


if __name__ == "__main__":
    fileContent = readFile("././05.txt")
    print(firstTask(fileContent))
    print(secondTask(fileContent))
