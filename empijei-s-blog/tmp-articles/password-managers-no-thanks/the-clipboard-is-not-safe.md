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

I will not talk about windows because it is so easy to [read other applications' memory](https://msdn.microsoft.com/en-us/library/ms680553%28VS.85%29.aspx) with user privileges that reading the clipboard becomes purposeless.
