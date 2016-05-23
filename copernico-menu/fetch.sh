wget 2>/dev/null -O /tmp/menu.pdf $(wget -O- 'http://www.copernicomilano.it/cafe-copernico/' 2>/dev/null | 
grep  "SCOPRI IL MEN. DI OGGI" | 
grep -Po 'href="\K[^"]*')
xdg-open /tmp/menu.pdf
