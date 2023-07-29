SHELL := cmd.exe

dev:
	nodemon --exec go run main.go

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
	test -d "./_ui" || \
	mkdir "./_ui" && \
	cp -R "../front-end/build/"* "./_ui" && \
	echo "success to inject fe to _ui folder"
	
