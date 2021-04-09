
test: vendor
	@echo Testing: Account
	cd account && go test ./... && cd ..
	@echo Testing: Common
	cd common && go test ./... && cd ..
	@echo Testing: form
	cd form && go test ./... && cd ..
	@echo Testing: gateway
	cd gateway && go test ./... && cd ..
	@echo Testing: logAggregator
	cd logAggregator && go test ./... && cd ..
	@echo Testing: mailto
	cd mailto && go test ./... && cd ..


up: vendor
	@echo "==== Building / Starting ===="
	docker-compose up --build

# local build with mysql
up_mysql: vendor
	@echo "==== Building / Starting (MySQL | secrets/.env.mysql) ===="
	docker-compose --env-file ./secrets/.env.mysql -f docker-compose.mysql.yml up --build 


# build and push all containers for production
push_prod: vendor
	docker-compose -f docker-compose.prod.yml build
	docker-compose -f docker-compose.prod.yml push

vendor: vendor_gateway vendor_account vendor_form vendor_logaggregator vendor_mailto
	@echo "==== Vendoring all dependencies ===="


vendor_gateway:
	@echo "==== Vendor: Gateway ===="
	cd ./gateway && go mod vendor
	

vendor_account:
	@echo "==== Vendor: Account ===="
	cd ./account && go mod vendor
	

vendor_form:
	@echo "==== Vendor: Form ===="
	cd ./form && go mod vendor
	

vendor_logaggregator:
	@echo "==== Vendor: LogAggregator ===="
	cd ./logAggregator && go mod vendor


vendor_mailto:
	@echo "==== Vendor: mailto ===="
	cd ./mailto && go mod vendor
