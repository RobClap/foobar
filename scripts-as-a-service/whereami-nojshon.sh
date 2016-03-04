whereami(){
	local output=$(curl "https://maps.googleapis.com/maps/api/browserlocation/json?browser=firefox&key=AIzaSyDBgL8fm9bD8RLShATOLI1xKmFcZ4ieMkM&sensor=true" \
		--data-urlencode "$(nmcli -f SSID,BSSID,SIGNAL dev wifi list | sed -rn 's/(.*[^ ])[ ]*(..:..:..:..:..:..)  ([0-9]*) .*/\&wifi=mac:\2|ssid:\1|ss:\3/p' | tr -d '\n')" 2>/dev/null |
	egrep "(lat|lng|accuracy)" | 
	tr -d '" ,' | 
	sed -r -e 's/lat:/latitude = /' -e 's/lng:/longitude = /' -e 's/accuracy:([0-9]*)/accuracy = \1 m/g')
	local lat=$(echo -en "$output" | grep -oP "latitude = \K[0-9]+\.[0-9]{2}")
	local long=$(echo -en "$output" | grep -oP "longitude = \K[0-9]+\.[0-9]{2}")
	local acc=$(echo -en "$output" | grep -oP "accuracy = \K[0-9]+")
	local address=$(curl "http://maps.googleapis.com/maps/api/geocode/json?latlng=$lat,$long&sensor=false" 2>/dev/null| 
	grep -oP "\"formatted_address\" : \"\K[^\"]+" | head -n 1)
	#jshon -e "results" -e 0 -e "formatted_address")
	echo -e "You are here: \n\tlatitude: $lat\n\tlongitude: $long\n\taddress:$address \n\taccuracy: $acc m" 
}


whereami
