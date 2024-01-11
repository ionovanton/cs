import os

readmeName = "README.md"
headerTemplate = ['## ', 'FOLDER_NAME', ' &nbsp;&nbsp;![', 'FOLDER_NAME', '](https://progress-bar.dev/', 'UNIT_PERCENT', '/?title=', 'UNIT_DONE', '/', 'UNIT_TOTAL', ')\n']
unitTemplate = ['### ', 'FOLDER_NAME', '\n![', 'FOLDER_NAME', '](https://progress-bar.dev/', 'UNIT_PERCENT', '/?title=', 'UNIT_DONE', '/', 'UNIT_TOTAL', ')\n']
# ## cs &nbsp;&nbsp;![Progress](https://progress-bar.dev/0/?title=111/5563)
### aads
# ![Progress](https://progress-bar.dev/2/?title=31/1610)

def generate(currentPath: str, currentFolder: str) -> list[int, int]:
    readmePath = currentPath + '/' + readmeName
    listdirs = os.listdir(currentPath)
    dirs = [
        x for x in listdirs
        if os.path.isdir(currentPath + '/' + x) is True
        and x[0] != '.'
        and x != 'archive'
    ]
    unitDone, unitTotal = 0, 0
    if len(dirs) != 0:
        listing = []
        for x in dirs:
            res = generate(currentPath + '/' + x, x)
            listing.append([x, res[0], res[1]])
            unitDone += res[0]
            unitTotal += res[1]
        try:
            os.remove(readmePath)
        except OSError:
            pass
        headerList = headerTemplate[:]
        for i, v in enumerate(headerList):
            if v == 'FOLDER_NAME':
                headerList[i] = currentFolder
            elif v == 'UNIT_DONE':
                headerList[i] = str(unitDone)
            elif v == 'UNIT_TOTAL':
                headerList[i] = str(unitTotal)
            elif v == 'UNIT_PERCENT':
                headerList[i] = str(int((unitDone / unitTotal) * 100))
        header = "".join(str(x) for x in headerList)
        with open(readmePath, "a") as file:
            file.write(header)
            for x in listing:
                folderName, done, total = x
                unitList = unitTemplate[:]
                for i, v in enumerate(unitList):
                    if v == 'FOLDER_NAME':
                        unitList[i] = folderName
                    elif v == 'UNIT_DONE':
                        unitList[i] = str(done)
                    elif v == 'UNIT_TOTAL':
                        unitList[i] = str(total)
                    elif v == 'UNIT_PERCENT':
                        unitList[i] = str(int((done / total) * 100))
                unit = "".join(str(x) for x in unitList)
                file.write(unit)
    else:
        with open(readmePath) as file:
            head = [next(file) for _ in range(2)]
            splitString = str.split(head[1])
        # parse total and done
        unitDone, unitTotal = int(splitString[1]), int(splitString[3])
    return [unitDone, unitTotal]

abspath = os.path.abspath(os.getcwd())
folderName, progress = generate(abspath, 'cs')
