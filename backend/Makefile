build:
	./backend.sh "build"

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
	./backend.sh "test"

clean:
	./backend.sh "clean"

logs:
	./backend.sh "logs_backend"

lint:
	./backend.sh "lint_backend"

format:
	./backend.sh "format_backend"