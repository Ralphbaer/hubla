build:
	./make.sh "build"

help:
	@echo "Management commands"
	@echo ""
	@echo "Usage:"
	@echo "  ## Root Commands"
	@echo "    make build           Build all project services."
	@echo "    make test            Run tests for both backend and frontend."
	@echo "    make clean           Clean the directory tree of produced artifacts."
	@echo "    make lint            Run static code analysis for both backend and frontend (lint)."
	@echo "    make format          Run code formatter for both backend and frontend."
	@echo ""
	@echo "  ## Utility Commands"
	@echo "    make setup-env       Setup github hooks locally."
	@echo ""

test:
	./make.sh "test"

clean:
	./make.sh "clean"

logs:
	./make.sh "logs"

lint:
	./make.sh "lint"

format:
	./make.sh "format"
	
setup-env:
	cp -R ./github/hooks/* .git/hooks/
	chmod +x .git/hooks/pre-commit
	chmod +x .git/hooks/pre-push

run:
	docker-compose up --build