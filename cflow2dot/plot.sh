cflow Tris.c -l > cflow.out

echo "digraph {"
LastLevel=0
LastFunc="init"
ContextFunc="init"
declare -a Stack
cat cflow.out|
while read line
do
	Level="$(echo $line | grep -Po '[0-9]+(?=})')"
	Func="$(echo $line | grep -Po '[a-zA-Z0-9_-]+(?=\()')"
	if [[ "$Level" > "$LastLevel" ]]
	then
		Stack[$LastLevel]=$LastFunc
		ContextFunc="$LastFunc"
	elif [[ "$Level" < "$LastLevel" ]] 
	then
		ContextFunc=${Stack[$(($Level-1))]}
	fi
	echo "$ContextFunc -> $Func;"
	LastFunc="$Func"
	LastLevel="$Level"
done
echo "}"
