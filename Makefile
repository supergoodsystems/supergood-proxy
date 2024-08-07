.PHONY: help
help:
	@echo 'Makefile for `supergood-proxy` project'
	@echo ''
	@echo 'Development supergood-proxy targets:'
	@echo '   make run-local                    Run `supergood-proxy` on the host'

################################################################################
# Development supergood-proxy targets
################################################################################

.PHONY: run-local
run-local:
	@go run \
	  ./cmd/ \
	  --path ./_config/dev.yml \
