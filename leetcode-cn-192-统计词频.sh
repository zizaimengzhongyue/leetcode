cat words.txt | tr " " "\n" | grep -E '[a-z]{1,}' | sort | uniq -c | sort -rn -k 1 | awk -F ' ' '{print $2" "$1}'
