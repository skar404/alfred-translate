FULL_REPO:=skar404/alfred-translate
VERSION:=1.0.0.test_github_apisss

clean:
	rm -rf ./bin

build: clean
	mkdir bin
	GOOS=darwin GOARCH=amd64 go build -o app_build
	GOOS=darwin GOARCH=arm64 go build -o app_build_arm
	cp -R ./assets/* ./bin

	cp app_build ./bin
	cd ./bin && zip -r translate-amd64.alfredworkflow ./* && mv translate-amd64.alfredworkflow ..

	rm ./bin/app_build
	cp app_build_arm ./bin/app_build
	cd ./bin && zip -r translate-arm64.alfredworkflow ./* && mv translate-arm64.alfredworkflow ..

	rm app_build
	rm app_build_arm
	rm -rf ./bin

build_deploy:
	go build .github/deploy/main.go
	mv main deploy

