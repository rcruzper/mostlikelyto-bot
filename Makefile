test:
	go test ./... -v

compile-arm: test
	CGO_ENABLED=0 GOARCH=arm GOOS=linux go build -a -installsuffix cgo -o app src/app/main.go

compile-macos: test
	CGO_ENABLED=0 go build -a -installsuffix cgo -o app src/app/main.go

clean:
	@docker rm mariadb-mostlikelyto -f

database: clean
	@docker run --name mariadb-mostlikelyto -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mariadb:10.3
