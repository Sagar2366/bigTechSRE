'''
Move all negative numbers to beginning and positive to end with constant extra space

An array contains both positive and negative numbers in random order. 
Rearrange the array elements so that all negative numbers appear before all positive numbers.

Examples : 

Input: -12, 11, -13, -5, 6, -7, 5, -3, -6
Output: -12 -13 -5 -7 -3 -6 11 6 5
'''

arr = [-1, 0,1,-2,0,0,-1,1,1,2,-1,-2]

leftPtr = 0
rightPtr = len(arr)-1

while leftPtr <= rightPtr:
    while arr[leftPtr] < 0 :
        leftPtr += 1
    while arr[rightPtr] > 0:
        rightPtr -= 1
    arr[leftPtr], arr[rightPtr] = arr[rightPtr], arr[leftPtr]
    leftPtr += 1
    rightPtr -= 1
print(arr)
        
