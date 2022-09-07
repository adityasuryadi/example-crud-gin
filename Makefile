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
