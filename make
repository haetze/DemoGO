go build $1
git add $1 -f 
git commit -m $2 -a
git push -u origin master 

