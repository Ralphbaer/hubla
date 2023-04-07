default: build

package-name=transaction
service-name=transaction-service
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

build_linux: clean_build
	GOOS=linux GOARCH=amd64 go build -o $(bin-dir)/$(service-name)

run:
	go run main.go
	
test:
	make unit-test && make integration-test

unit-test:
	mkdir -p $(artifacts-dir) && \
	gotestsum --format short-verbose --junitfile $(artifacts-dir)/junit-report.xml \
	./... github.com/Ralphbaer/hubla/backend/common... \
	-coverprofile=$(artifacts-dir)/coverage.out

integration-test:
	go test ./usecase -v -run TestSalesUseCase_GetNearestSales
	
cover-html:
	go tool cover -html=$(artifacts-dir)/coverage.out -o $(artifacts-dir)/coverage.html

cover: test cover-html
	open $(artifacts-dir)/coverage.html

clean:
	rm -rf $(bin-dir)
	rm -rf $(artifacts-dir)

image: build_linux
	docker build -t hubla/$(service-name):v1 .

container:
	docker run --env SERVER_ADDRESS=:3000 -it -p 127.0.0.1:3000:3000 hubla/$(service-name):v1 bash

localenv-withData:
	rm -rf .localenv/withData/.data && \
	docker-compose -f .localenv/withData/docker-compose.yml up --force-recreate
	
localenv:
	rm -rf .localenv/.data && \
	docker-compose -f .localenv/docker-compose.yml up --force-recreate

doc:
	swagger init spec \
		--title "Hubla Transaction API" \
		--description "The Hubla Transaction API efficiently imports creator and affiliate data \
		from the Hubla ecosystem, including balances, transactions, and product names. \
		This secure and streamlined service simplifies financial data handling and integration \
		within your application." \
		--version 1.0.0 \
		--scheme http \
		--consumes application/json \
		--produces application/json
