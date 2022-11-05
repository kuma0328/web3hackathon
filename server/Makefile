# help ターゲットをデフォルトのターゲットにする
.DEFAULT_GOAL := help

# env
ENV_LOCAL_FILE := .env.local
ENV_LOCAL       = $(shell cat $(ENV_LOCAL_FILE))

# go
GO           := go run
GO_MAIN_PATH := ./src

# docker
DOCKER_FILE_DIR                 := ./docker
DOCKER_COMPOSE_LOCAL            := $(DOCKER_FILE_DIR)/docker-compose.local.yml
DOCKER_COMPOSE_LOCAL_DATABASE   := $(DOCKER_FILE_DIR)/docker-compose.local.database.yml
DOCKER_COMPOSE_LOCAL_SERVER     := $(DOCKER_FILE_DIR)/docker-compose.local.server.yml
DOCKER_COMPOSE_LOCAL_REDIS     := $(DOCKER_FILE_DIR)/docker-compose.local.redis.yml
DOCKER_EXEC                     := docker exec -it
GITHUB_REPOSITORY_NAME          := web3hackathon
DB_CONTAINER_NAME               := $(GITHUB_REPOSITORY_NAME)-database

# dir
SRC_DIR := ./src
DATA_DIR := ./database/data

# rm
RM:=rm -rf

.PHONY: up
up: ## docker環境を立ち上げる
	$(ENV_LOCAL) docker-compose \
	-f $(DOCKER_COMPOSE_LOCAL) \
	-f $(DOCKER_COMPOSE_LOCAL_DATABASE) \
	-f $(DOCKER_COMPOSE_LOCAL_REDIS) \
	-f $(DOCKER_COMPOSE_LOCAL_SERVER) up

.PHONY: down
down: ## dockerイメージを削除し、docker環境を閉じる
	docker-compose \
	-f $(DOCKER_COMPOSE_LOCAL) \
	-f $(DOCKER_COMPOSE_LOCAL_DATABASE) \
	-f $(DOCKER_COMPOSE_LOCAL_REDIS) \
	-f $(DOCKER_COMPOSE_LOCAL_SERVER) down \
	--rmi all --volumes --remove-orphans

.PHONY: fclean
fclean:down del-volumes ## マウントしたデータを削除、またdockerイメージも削除する

.PHONY: re
re:fclean up ## 完全に初期化した状態でdocker環境を立ち上げる

.PHONY: del-volumes
del-volumes:del-data

.PHONY: del-data
del-data:
	$(RM) $(DATA_DIR)


.PHONY: attach-db
attach-db: ## dockerのdbコンテナにアクセスする
	$(DOCKER_EXEC) $(DB_CONTAINER_NAME) bash

.PHONY: lint
lint: ## goのコードを整形する
	cd $(SRC_DIR) && gofmt -l -w .

.PHONY: deploy
deploy:


.PHONY: help
help: ## コマンドの一覧を標示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'