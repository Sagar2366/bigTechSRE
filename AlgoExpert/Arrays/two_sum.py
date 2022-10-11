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

## Solution 2: Hash Table
## TC - O(N), SC - 0(N)

def twoNumberSum(array, targetSum):
    # Write your code here.

    nums = {}
    for num in array:
        secondNum = targetSum - num
        if secondNum in nums:
            return [secondNum, num]
        else:
            nums[num] = True
    return []
    pass


## Solution 3: Two Pointers
## TC - O(NLOGN), SC - O(1)
def twoNumberSum(array, targetSum):
    # Write your code here.

    array.sort()
    left = 0
    right = len(array) - 1

    while left < right:
        currentSum = array[left] + array[right]
        if currentSum == targetSum:
            return [array[left], array[right]]
        elif currentSum < targetSum:
            left += 1
        elif currentSum > targetSum:
            right -= 1
    return []