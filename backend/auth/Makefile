default: build

package-name=auth
service-name=auth-service
artifacts-dir=./.artifacts
bin-dir=./.bin

.PHONY: gen
gen:
	go generate ./...

clean_build:
	echo "Cleaning build directory..." && \
	rm -rf $(bin-dir) && \
    mkdir -p $(bin-dir)

build: clean_build
	go build -o $(bin-dir)/$(service-name)

run:
	go run main.go
	
test:
	make unit-test && make integration-test

unit-test:
	mkdir -p $(artifacts-dir) && \
	gotestsum --format short-verbose --junitfile $(artifacts-dir)/junit-report.xml \
	./... github.com/Ralphbaer/hubla/common... \
	-coverprofile=$(artifacts-dir)/coverage.out

cover-html:
	go tool cover -html=$(artifacts-dir)/coverage.out -o $(artifacts-dir)/coverage.html

cover: test cover-html
	open $(artifacts-dir)/coverage.html

clean:
	rm -rf $(bin-dir)
	rm -rf $(artifacts-dir)

localenv:
	rm -rf .localenv/.data && \
	docker-compose -f .localenv/docker-compose.yml up --force-recreate

doc:
	swagger init spec \
		--title "Hubla Auth API" \
		--description "The Hubla Auth API is a secure and reliable authentication service 
			specifically designed for users within the Hubla company. It is responsible for managing user credentials, 
			verifying user identities, and issuing access tokens to grant appropriate permissions across the company's applications. 
			By adhering to industry-standard security practices and protocols,
			the Hubla Auth API ensures a seamless and secure user authentication experience, which helps maintain trust and safeguard sensitive user information." \
		--version 1.0.0 \
		--scheme http \
		--consumes application/json \
		--produces application/json
