#================================
#== DOCKER ENVIRONMENT
#================================
COMPOSE := @docker-compose

dcb:
	${COMPOSE} build

dcuf:
ifdef f
	${COMPOSE} up -d --${f}
endif

dcubf:
ifdef f
	${COMPOSE} up -d --build --${f}
endif

dcu:
	${COMPOSE} up -d --build

dcd:
	${COMPOSE} down

#================================
#== GOLANG ENVIRONMENT
#================================
GO := @go
GIN := @gin

goinstall:
	${GO} get .

godev:
	${GIN} -a 3001 -p 3001 -b bin/main run main.go

goprod:
	${GO} build -o main .

gotest:
	${COMPOSE} -f test.docker-compose.yaml up -d pos_mysql_test
	${COMPOSE} -f test.docker-compose.yaml up --build --abort-on-container-exit 
	${COMPOSE} -f test.docker-compose.yaml down --volumes
	

goformat:
	${GO} fmt ./...

full-test:
	# @echo "Running the full test..."
	${GO} test -v

unittest:
	${GO} test -v
