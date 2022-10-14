## Arrays in Go

- Fixed size single type array
- Syntax
  `var variablename = [size] datatype {comma separated elements}`
  `var variable_name = [size] datatype  <----- Declaring an empty array`

- Adding & accessing element by their index in array

## Slices in Go

- Slice is an abstraction of an array
- Variable sized array
- Can get subarray of its own
- Resized when needed
- Add new element in slice using append() at the end of an slice
- Syntax
  `
  var variablename [] datatype
  var variablename = [] datatype {}
  variablename := [] datatype {}
  `