build: compile copy-resources

compile:
	go build -o out/charchapoint main/main.go

copy-resources:
	mkdir -p out/public
	cp public/index.html out/public/
	cp config.json out/