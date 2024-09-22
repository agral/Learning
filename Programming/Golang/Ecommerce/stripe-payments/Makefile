FRONTEND_PORT=4000
BACKEND_PORT=4001

.PHONY: check_env clean build build_frontend build_backend start start_frontend start_backend stop stop_frontend stop_backend

# Note: makes check_env somewhat obsolete, but it can still be used
# to quickly check that set_env.sh has been sourced.
include set_env.sh

start: start_backend start_frontend
build: check_env clean build_backend build_frontend
stop: stop_frontend stop_backend

build_frontend:
	@echo "Building front end..."
	@go build -o dist/gostripe ./cmd/web
	@echo "Successfully built front end."

build_backend:
	@echo "Building back end..."
	@go build -o dist/gostripe_api ./cmd/api
	@echo "Successfully built back end."

clean:
	@rm -rf "dist/*"
	@go clean
	@echo "Cleaned."

start_frontend: build_frontend
	./dist/gostripe -port=${FRONTEND_PORT} &
	@echo "Started the front end."

start_backend: build_backend
	./dist/gostripe_api -port=${BACKEND_PORT} &
	@echo "Started the back end."

stop_frontend:
	@-pkill -SIGTERM -f "gostripe -port=${FRONTEND_PORT}"
	@echo "Stopped the front end server."

stop_backend:
	@-pkill -SIGTERM -f "gostripe_api -port=${BACKEND_PORT}"
	@echo "Stopped the back end server."

check_env:
ifndef STRIPE_API_TEST_KEY
$(error STRIPE_API_TEST_KEY is unset)
endif
	@echo "The required environment variables are set."
