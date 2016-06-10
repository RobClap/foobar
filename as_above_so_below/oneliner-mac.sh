# Requires networksetup -setairportpower airport on
2> /dev/null curl "https://maps.google.com/maps/api/staticmap?maptype=satellite&center=$(
2> /dev/null curl "https://maps.googleapis.com/maps/api/browserlocation/json?browser=firefox&sensor=true" \
	--data-urlencode "$(/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport -s | 
		grep -Eo "[^ ].* ..:..:..:..:..:.. -[0-9]{2}" | 
		awk '{printf "&wifi=mac:%s|ssid:",$(NF-1); for(i=1;i<NF-2;i++){printf "%s ", $i}; printf "%s|ss:%s",$(NF-2),substr( $NF, 2, 3)}')" | 
		grep -Eo "[0-9]+\.[0-9]{7},?" | 
		tr -d '\n'
)&zoom=15&scale=2&size=640x360" > ~/as_above_so_below.png 
osascript -e 'tell application "Finder" to set desktop picture to POSIX file "~/as_above_so_below.png"'

