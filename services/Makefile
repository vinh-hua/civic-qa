

up: vendor
	@echo "==== Building / Starting ===="
	docker-compose up --build


vendor: vendor_gateway vendor_account vendor_form vendor_logaggregator
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



