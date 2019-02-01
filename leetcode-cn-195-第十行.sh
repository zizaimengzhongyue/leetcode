cat file.txt | awk 'BEGIN{line=0;}{line+=1;if(line==10){print $0}}'
