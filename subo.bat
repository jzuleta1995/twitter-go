git add .
git commit -m "Last commit"
git push
go build -o main main.go
del main.zip
tar.exe -a -cf main.zip main
