#!/bin/bash

fbif(){
	local INPUT=$1
	if [[ $(expr $INPUT % 15) == 0 ]] 
	then
		echo fizzbuzz
	elif [[ $(expr $INPUT % 3) == 0 ]] 
	then
		echo fizz
	elif [[ $(expr $INPUT % 5) == 0 ]] 
	then
		echo buzz
	else
		echo $INPUT
	fi
}
fizzbuzzoneline(){
	expr $1 % 15 == 0 > /dev/null && echo fizzbuzz && return || 
	expr $1 % 5 == 0 > /dev/null && echo buzz && return || 
	expr $1 % 3 == 0 > /dev/null && echo fizz && return || 
	echo $1 
}
for i in `seq 100`
do
	echo $(fizzbuzzoneline $i)
done
