cat file.txt | awk 'BEGIN{i=1;}{i+=1;temp=split($0,arr," ");for(j=1;j<=temp;j++){if(!ans[j]){ans[j]=arr[j]}else{ans[j]=ans[j]" "arr[j];}}}END{for(j=1;j<=temp;j++){printf ans[j]"\n"}}'
