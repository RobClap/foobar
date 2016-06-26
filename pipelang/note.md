# Features
Splitting input in tokens, put it in structs
split pipes, mux pipes, merge pipes
execute block based on signal emitted by previous one

# Example syntax
splitComment(commentchar string){
	in | re "^\s*" + commentchar | 1
	in | re -v "^\s*" + commentchar | 2
}

cat inputFile -> splitComment -1> > comments.txt
										-2> > non-comments.txt

Same as

cat inputFile 
splitComment
-1> > comments.txt
-2> > non-comments.txt

Same as

cat inputFile | splitComment "#" |1 > comments.txt |2 > non-comments.txt
