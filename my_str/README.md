# Strings

Messing with strings in go and also string encoding in genera.


## Structure

Structured as a go library. Code is run using ```$go test```.  


## Testing in go

You need a filename, ending with *_test.go*, you need a function inside that file with a signature:
```go
func (t *testing.T)
```  
Put the code you want to test inside that function. You can use ```t.Error``` or ```t.Fail``` to determine 
code success or failure.  
Inside the package, run ```$go test```.  


