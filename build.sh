#! /bin/bash
current=`pwd`
echo 'const project_folder = "'$current'"' >> $current/engine/engine.go
echo 'const project_folder = "'$current'"' >> $current/parse/readExcel.go
go build main.go && mv main $GOPATH/bin/gcase
echo "now you can use command: gcase"


