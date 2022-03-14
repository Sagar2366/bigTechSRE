Selection Sort:
Space Complexity: o(n)
Time Complexity: o(n^2)


Program:

def selectionSort(array):
    current = 0
	while current < len(array) - 1 :
		smallInd = current
		for i in range(current + 1, len(array)):
			if array[smallInd] > array[i]:
				smallInd = i
		swap(current, smallInd, array)
		current += 1

	return array

def swap(i, j, arr):
	arr[i], arr[j] = arr[j], arr[i]
	

 
