package simple_cypher

import (
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const kb_alphabet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func rot(str_in, charset string, shift int) (out string) {

	for _, c := range str_in {
		if strings.Contains(charset, string(c)) {
			out += string((int(c) + shift) % 256)
		} else {
			out += string(c)
		}
	}
	return
}

/*
def substitute(str_in,alpha1,alpha2):
    out=""
    for c in str_in:
        out+=alpha2[alpha1.index(c)]
    return out

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

def reverse(str_in):
    return str_in[::-1]

def prev_as_key(str_in,charset,start=0):
    out=""
    prev_char=start
    for c in str_in:
        out+=charset[(charset.index(c)+prev_char)%len(charset)]
        prev_char=charset.index(c)+1
    return out

def prev_as_key_decode(str_in,charset,start=0):
    out=""
    prev_char=start
    for c in str_in:
        out+=charset[(charset.index(c)-prev_char)%len(charset)]
        prev_char=charset.index(out[-1])+1
    return out

def cascade(str_in,charset,start=0):
    out=""
    shift=start
    for c in str_in:
        out+=charset[(charset.index(c)+shift)%len(charset)]
        shift+=charset.index(c)+1
    return out

def cascade_decode(str_in,charset,start=0):
    out=""
    shift=start
    for c in str_in:
        out+=charset[(charset.index(c)-shift)%len(charset)]
        shift+=charset.index(out[-1])+1
    return out
*/
