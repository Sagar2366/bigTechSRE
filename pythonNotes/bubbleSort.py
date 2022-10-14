Space Complexity: o(1)
Time Complexity: O(n2), best case - o(n)
  
  
  def bubbleSort(array):
    isSorted = False
	counter = 0
	while not isSorted:
		isSorted = True
		for i in range(len(array) - 1 - counter):
			if array[i] > array[i+1]:
				array[i], array[i+1] = array[i+1], array[i]
				isSorted = False
		counter += 1
	return array
