mkdir -p output/{conf,log,bin}
cp config/*.yaml output/conf
go build  -o output/bin/main