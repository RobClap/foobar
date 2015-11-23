package easier

import (
	"strings"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Kb_alphabet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
const PRINTABLE_ASCII = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

//Function meant to give a more reliable % operation
//Uses the second parameter's absolute value and
//Treats negative numbers as numbers "rotating" the module
//in the other direction.
//i.e. -10 % 7 = 4
//This function ALWAYS returns a number between 0 and b
func rotatingMod(a, b int) int {
	if b < 0 {
		b = -b
	}
	if a < 0 {
		if c := b + a%b; c == b {
			return 0
		} else {
			return c
		}
	} else {
		return a % b
	}
}

//TODO docme
func Rot(str_in, charset string, shift int) (out string) {
	runes := make([]rune, len(str_in))
	for pos, c := range str_in {
		if i := strings.Index(charset, string(c)); i >= 0 {
			runes[pos] = rune(charset[rotatingMod(i+shift, len(charset))])
		} else {
			runes[pos] = rune(c)
		}
	}
	out = string(runes)
	return
}

//TODO docme
func Substitute(str_in, alpha1, alpha2 string) (out string) {
	if len(alpha2) != len(alpha1) {
		//TODO Panic
	}
	//TODO also check for duplicate in input?
	runes := make([]rune, len(str_in))
	for pos, c := range str_in {
		if i := strings.Index(alpha1, string(c)); i >= 0 {
			runes[pos] = rune(alpha2[i])
		} else {
			runes[pos] = rune(c)
		}
	}
	out = string(runes)
	return
}

//abcde → edcba
func Reverse(str_in string) (out string) {
	n := len(str_in)
	runes := make([]rune, n)
	for _, rune := range str_in {
		n--
		runes[n] = rune
	}
	out = string(runes[n:])
	return
}

//abcde,2 → cdeab
func Shift(str_in string, split_index int) (out string) {
	runes := make([]rune, len(str_in))
	for i := range str_in {
		runes[i] = rune(str_in[rotatingMod(i+split_index, len(str_in))])
	}
	out = string(runes)
	return
}

func PreviousAsKey(str_in, charset string, start int) (out string) {
	runes := make([]rune, len(str_in))
	prev_char := start
	for pos, c := range str_in {
		if i := strings.Index(charset, string(c)); i >= 0 {
			runes[pos] = rune(charset[rotatingMod(i+prev_char, len(charset))])
			prev_char = i
		} else {
			runes[pos] = rune(c)
		}
	}
	out = string(runes)
	return
}

func PreviousAsKeyDecode(str_in, charset string, start int) (out string) {
	runes := make([]rune, len(str_in))
	prev_char := start
	for pos, c := range str_in {
		if i := strings.Index(charset, string(c)); i >= 0 {
			tmp := rotatingMod(i-prev_char, len(charset))
			runes[pos] = rune(charset[tmp])
			prev_char = tmp
		} else {
			runes[pos] = rune(c)
		}
	}
	out = string(runes)
	return
}

func Cascade(str_in, charset string, start int) (out string) {
	runes := make([]rune, len(str_in))
	prev_char := start
	for pos, c := range str_in {
		if i := strings.Index(charset, string(c)); i >= 0 {
			tmp := rotatingMod(i+prev_char, len(charset))
			runes[pos] = rune(charset[tmp])
			prev_char = tmp
		} else {
			runes[pos] = rune(c)
		}
	}
	out = string(runes)
	return
}

func CascadeDecode(str_in, charset string, start int) (out string) {
	runes := make([]rune, len(str_in))
	prev_char := start
	for pos, c := range str_in {
		if i := strings.Index(charset, string(c)); i >= 0 {
			tmp := rotatingMod(i-prev_char, len(charset))
			runes[pos] = rune(charset[tmp])
			prev_char += tmp
		} else {
			runes[pos] = rune(c)
		}
	}
	out = string(runes)
	return
}

func Vigenere(str_in, key, charset string) (out string) {
	runes := make([]rune, len(str_in))
	for pos, c := range str_in {
		in_i := strings.Index(charset, string(c))
		key_i := strings.Index(charset, string(key[pos%len(key)]))
		if in_i >= 0 && key_i >= 0 {
			runes[pos] = rune(charset[rotatingMod(in_i+key_i, len(charset))])
		} else {
			runes[pos] = rune(c)
		}
	}
	out = string(runes)
	return
}

func VigenereDecode(str_in, key, charset string) (out string) {
	runes := make([]rune, len(str_in))
	for pos, c := range str_in {
		in_i := strings.Index(charset, string(c))
		key_i := strings.Index(charset, string(key[pos%len(key)]))
		if in_i >= 0 && key_i >= 0 {
			runes[pos] = rune(charset[rotatingMod(in_i-key_i, len(charset))])
		} else {
			runes[pos] = rune(c)
		}
	}
	out = string(runes)
	return
}

func Transpose(str_in string, row_length int, key []int) (out string) {
	//	runes := make([]rune, len(str_in))

	return
}

func TransposeDecode(str_in string, row_length int, key []int) (out string) {
	return
}

/*
#TODO add key
def transpose(str_in,row_length):
    out=""
    table=[]
    counter=0
    while(counter<len(str_in)):
        tmplst=[]
        for x in range(0,row_length):
            tmplst.append(str_in[counter])
            counter+=1
            if counter>=len(str_in):
                break
        table.append(tmplst)
    for x in range(0,row_length):
        for y in range(0,len(table)):
            if y*row_length+x<len(str_in):
                out+=table[y][x]
    return out

def transpose_decode(str_in,row_length):
    out=""
    incomplete_counter=len(str_in)%row_length
    row_length=int(len(str_in) / row_length)
    counter=0
    table=[]
    while counter<len(str_in):
        tmplst=[]
        for x in range(0,row_length):
            tmplst.append(str_in[counter])
            counter+=1
        if incomplete_counter>0:
            tmplst.append(str_in[counter])
            counter+=1
            incomplete_counter-=1
        table.append(tmplst)
    for x in range(0,len(table[0])):
        for y in range(0,len(table)):
            if counter>0:
                out+=table[y][x]
                counter-=1
    return out
*/
