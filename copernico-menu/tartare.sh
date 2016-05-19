wget 2>/dev/null -O /tmp/menu.pdf $(wget -O- 'http://www.copernicomilano.it/cafe-copernico/' 2>/dev/null | 
grep  "SCOPRI IL MEN. DI OGGI" | 
grep -Po 'href=" \K[^"]*')
pdftotext /tmp/menu.pdf - |
grep -i "tartare" &&
echo "C'è la tartare" ||
echo "Non c'è la tartare"
