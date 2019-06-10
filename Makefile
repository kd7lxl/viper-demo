demo: build
	./viper-demo --tags=location=Tacoma
	./viper-demo --tags=location=Tacoma,size=medium
	./viper-demo --tags=location=Tacoma --tags=county=US
	./viper-demo --tags=location=Tacoma,size=medium --tags=county=US

.PHONY: build
build:
	go build -o viper-demo .
