# go-simple-frw
go simple file reader and writer<br>
<br><br>

<h2>Installation / usage</h2>
<pre>
go get -u "github.com/amirhoseinjfri/go-simple-frw" 


import gsf "github.com/amirhoseinjfri/go-simple-frw"
data := gsf.FileData{}
file, err := os.OpenFile("list.txt", os.O_CREATE, os.ModePerm)
CheckError(err)
defer file.Close()
gsf.ReadFile(file)
</pre>

function writetofile has 4 thing to be done : <br>
 1 : data to write into a file as string <br>
 2 : data called with typed FileData <br>
 3 : true|false => if true it will write to the end of file <br>
 4 : file path as os package
