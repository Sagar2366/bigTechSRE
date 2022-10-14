def isPalindrome(string):
    if string == string[::-1]:
		return True
	else:
		return False
  
  
 
Iterative:
	
def isPalindrome(String):
	l = 0
	r = len(String) - 1
	
	while l <  r:
		if String[l] != String[r]:
			return False
		l += 1
		r -= 1
	return True
