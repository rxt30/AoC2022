import numpy as np


def readFile(fileName):
    file = open(fileName, "r")
    content = file.readlines()
    return [line.rstrip("\n") for line in content]


def readTreeGrid(fileContent):
    treeGrid = []
    for row in fileContent:
        treeRow = []
        for tree in row:
            treeRow.append(tree)
        treeGrid.append(np.array(treeRow))
    return np.array(treeGrid)


def firstTask(fileContent):
    treeGrid = readTreeGrid(fileContent)
    visibleTreeCount = len(treeGrid) * 2 + len(treeGrid[0]) * 2 - 4
    for i in range(1, len(treeGrid) - 1):
        for j in range(1, len(treeGrid[0]) - 1):
            currentTreeHeight = treeGrid[i][j]
            leftSide = sorted(treeGrid[i][0:j], reverse=True)
            rightSide = sorted(treeGrid[i][j + 1 :], reverse=True)
            upSide = sorted(
                [treeHeight[j] for treeHeight in treeGrid[0:i]], reverse=True
            )
            downSide = sorted(
                [treeHeight[j] for treeHeight in treeGrid[i + 1 :]], reverse=True
            )
            if (
                leftSide[0] < currentTreeHeight
                or rightSide[0] < currentTreeHeight
                or upSide[0] < currentTreeHeight
                or downSide[0] < currentTreeHeight
            ):
                visibleTreeCount += 1
    return visibleTreeCount


def getVisionScore(trees, currentTreeHeight):
    i = 0
    for i, tree in enumerate(trees):
        if tree >= currentTreeHeight:
            return i + 1
    return i + 1


def secondTask(fileContent):
    treeGrid = readTreeGrid(fileContent)
    visionScores = []
    for i in range(1, len(treeGrid) - 1):
        for j in range(1, len(treeGrid[0]) - 1):
            currentTreeHeight = treeGrid[i, j]
            upScore = getVisionScore(
                treeGrid[:i, [j]].flatten()[::-1], currentTreeHeight
            )
            downScore = getVisionScore(
                treeGrid[i + 1 :, [j]].flatten(), currentTreeHeight
            )
            leftScore = getVisionScore(
                treeGrid[i, 0:j].flatten()[::-1], currentTreeHeight
            )
            rightScore = getVisionScore(
                treeGrid[i, j + 1 :].flatten(), currentTreeHeight
            )
            visionScores.append(upScore * downScore * leftScore * rightScore)
    return sorted(visionScores, reverse=True)[0]


if __name__ == "__main__":
    # fileContent = readFile("././test.txt")
    fileContent = readFile("././08.txt")
    print(firstTask(fileContent))
    print(secondTask(fileContent))
