clean:
	rm -rf ./bin

build: clean
	GOOS=darwin GOARCH=amd64 go build -o app_build
	GOOS=darwin GOARCH=arm64 go build -o app_build_arm
	cp -R ./assets/country-flags ./bin
	cp -R ./assets/* ./bin

	cp app_build ./bin
	cd ./bin && zip -r translate-amd64.alfredworkflow ./*

	rm ./bin/app_build
	cp app_build_arm ./bin/app_build
	cd ./bin && zip -r translate-arm64.alfredworkflow ./*

	rm app_build
	rm app_build_arm
	rm -rf ./bin
