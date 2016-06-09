wget 2>/dev/null -O- $(wget -O- 'http://www.copernicomilano.it/cafe-copernico/' 2>/dev/null | 
grep  "SCOPRI IL MEN. DI OGGI" | 
grep -Po 'href="\K[^"]*')|
grep -i "tartare" &&
echo "C'è la tartare" ||
echo "Non c'è la tartare"
