.DEFAULT_GOAL = help
.SILENT:
ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:)

build-linux:  ## Build de la version linux
	env GOOS=linux GOARCH=amd64 go build -o build/linux/copyfast ./cmd/copyfast.go
	
build-darwin: ## Build de la version darwin
	env GOOS=darwin GOARCH=amd64 go build -o build/darwin/copyfast ./cmd/copyfast.go

build-windows: ## Build de la version windows
	env GOOS=windows GOARCH=amd64 go build -o build/windows/copyfast ./cmd/copyfast.go

build-all: build-linux build-darwin build-windows ## build de toutes les versions 

test: build-linux ## Test de la version linux sur container
	docker build --force-rm -t locals/copyfast .
	docker run locals/copyfast

run: ## run sans compilation du projet GO
		go run cmd/copyfast.go $(ARGS)

ftp-server:
	docker run -p 21:21 -p 21000-21010:21000-21010 -e USERS="toto|toto" -e ADDRESS=127.0.0.1 delfer/alpine-ftp-server

help: #Pour gérer automatiquement l'aide ## Display all comands available
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

git: 
	git add -A 
	git commit -m "Auto Commit"
	git push