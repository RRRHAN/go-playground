dev:
	nodemon --exec "go run" main.go --signal SIGTERM

install-migration:
	go install -tags "sqlite3" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

install-fe:
	cd ../front-end && \
	echo "Changed directory to front-end" && \
	yarn install && \
	cd ../back-end && \
	echo "success to install fe" && \
	make build-fe

build-fe:
	cd ../front-end && \
	echo "Changed directory to front-end" && \
	yarn run build && \
	cd ../back-end && \
	echo "success to build fe" && \
	make inject-fe

inject-fe:
	rm -r "./ui" & \
	mkdir "./ui" && \
	cp -R "../front-end/build/"* "./ui" && \
	echo "success to inject fe to ui folder"
	
