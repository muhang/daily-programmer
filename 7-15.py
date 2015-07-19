def problem(testStr, problemStr):
    for probLetter in problemStr:
        if probLetter in testStr:
            testStr = testStr[testStr.index(probLetter)+1:len(testStr)]
        else:
            return False
    return True

if __name__ == "__main__":
    testA = problem("synchronized", "snond")
    testB = problem("misfunctioned", "snond")
    testC = problem("mispronounced", "snond")
    testD = problem("shotgunned", "snond")
    testE = problem("snond", "snond")
    print(testA)
    print(testB)
    print(testC)
    print(testD)
    print(testE)
