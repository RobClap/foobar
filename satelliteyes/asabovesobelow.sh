#On macosx depends on https://github.com/victor/whereami
#On linux depends on nmcli
#Optional dependecy imagemagick

[[ "$(uname -s)" =~ Linux ]] && 
whereami () {
curl "https://maps.googleapis.com/maps/api/browserlocation/json?browser=firefox&sensor=true" --data-urlencode "$(nmcli -f SSID,BSSID,SIGNAL dev wifi list | sed -rn 's/(.*[^ ])[ ]*(..:..:..:..:..:..)  ([0-9]*) .*/\&wifi=mac:\2|ssid:\1|ss:\3/p' | tr -d '\n')" | grep -Po '"l(at|ng)" : \K[0-9.,]+' | tr -d '\n'
}
wget "https://maps.google.com/maps/api/staticmap?maptype=satellite&center=$(whereami)&zoom=15&scale=2&size=640x360" -O ~/as_above_so_below.png
#This line expands the image to a wallpaper size. Depends on imagemagick and is gonna be probably useless since you window manager is going to do that anyway.
#convert wallpaper.png -scale 150% out.png
[[ "$(uname -s)" =~ Darwin ]] && osascript -e 'tell application "Finder" to set desktop picture to POSIX file "~/as_above_so_below.png"'
