# Problem Link - https://www.geeksforgeeks.org/sort-an-array-of-0s-1s-and-2s/
# Dutch National Flag

# Solution :

'''
Given an array A[] consisting of only 0s, 1s, and 2s. The task is to write a 
function that sorts the given array. The functions should put all 0s first, 
then all 1s and all 2s in last.

This problem is also the same as the famous “Dutch National Flag problem”. 
The problem was proposed by Edsger Dijkstra. The problem is as follows:

Given N balls of colour red, white or blue arranged in a line in random order.
You have to arrange all the balls such that the balls with the same colours are 
adjacent with the order of the balls, with the order of the colours being red,
white and blue (i.e., all red coloured balls come first then the white coloured 
balls and then the blue coloured balls). 
'''
'''

TC :- O(N)
SC :- O(N)

arr = [0,1,2,0,0,1,1,1,2]

count0 = 0
count1 = 0
count2 = 0

for i in range(len(arr)):
    if arr[i] == 0:
        count0 += 1
    elif arr[i] == 1:
        count1 += 1
    else:
        count2 += 1
print(count0, count1, count2)
sorted_arr = []
while count0 > 0 :
    sorted_arr.append(int(0))
    count0 -= 1
while count1 > 0:
    sorted_arr.append(int(1))
    count1 -= 1
while count2 > 0 :
    sorted_arr.append(int(2))
    count2 -= 1
print(sorted_arr)



'''
##############################################################################################################################################################
'''
TC :- O(N)
SC :- O(1)
'''

arr = [0,1,2,0,0,1,1,1,2]

ptr0 = 0
ptr1 = 0
ptr2 = len(arr)-1

while(ptr1 != ptr2):
    if arr[ptr1] == 0:
        arr[ptr0], arr[ptr1] = arr[ptr1], arr[ptr0]
        ptr0 += 1
        ptr1 += 1
    elif arr[ptr1] == 1:
        ptr1 += 1
    else:
        arr[ptr1], arr[ptr2] = arr[ptr2], arr[ptr1]
        ptr2 -= 1

print(arr)