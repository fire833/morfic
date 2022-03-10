
all:
	@echo "Hello there"

n:
	@echo "Building node server test binary..."
	@go build -o node/node-test node/main.go 
	@echo "Built."

nrun:
	@./node/node-test

a:
	@echo "Building API server test binary..."
	@go build -o api/api-test api/main.go
	@echo "Built."

arun:
	@./api/api-test

c:
	@echo "Building controller server test binary..."
	@go build -o controller/controller-test controller/main.go
	@echo "Built."

crun: 
	@./controller/controller-test