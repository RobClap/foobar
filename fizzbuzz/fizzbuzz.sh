#!/bin/bash

fizzbuzz(){
	local INPUT=$1
	local OUTPUT=""
	if [[ $(($INPUT % 3)) == 0 ]] 
	then
		OUTPUT="fizz"
	fi
	if [[ $(($INPUT % 5)) == 0 ]] 
	then
		OUTPUT=$OUTPUT"buzz"
	fi
	if [[ -z $OUTPUT ]]
	then
		OUTPUT=$INPUT
	fi
	echo $OUTPUT
}
wtffizzbuzz(){
	expr $1 % 15 == 0 > /dev/null && echo fizzbuzz && return || 
	expr $1 % 3 == 0 > /dev/null && echo fizz && return || 
	expr $1 % 5 == 0 > /dev/null && echo buzz && return || 
	echo $1 
}
for i in `seq 100`
do
	fizzbuzz $i
done
