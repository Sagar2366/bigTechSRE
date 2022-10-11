## Solution 1: Double For Loop
## TC - O(N2), SC - O(1)

def twoNumberSum(array, targetSum):
    # Write your code here.

    for i in range(len(array) - 1):
        firstNum = array[i]
        for j in range(i+1, len(array)):
            secondNum = array[j]
            if firstNum + secondNum == targetSum:
                return [firstNum, secondNum]
                
    return []
    pass

