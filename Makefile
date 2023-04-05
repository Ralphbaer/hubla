build:
	./make.sh "build"

help:
	@echo "Management commands"
	@echo ""
	@echo "Usage:"
	@echo "  ## Root Commands"
	@echo "    make build           Build all project services."
	@echo "    make test            Run tests on all projects."
	@echo "    make clean           Clean the directory tree of produced artifacts."
	@echo "    make lint            Run static code analysis (lint)."
	@echo "    make format          Run code formatter."
	@echo ""
	@echo "  ## Utility Commands"
	@echo "    make setup-env       Setup github hooks locally."
	@echo ""

test:
	./make.sh "test"

clean:
	./make.sh "clean"

lint:
	./make.sh "lint"

format:
	./make.sh "format"

run:
	docker-compose up --build