GP=$(shell echo $(GOPATH))

spl_build:
	@go install; \
	mv $(GP)/bin/spl .

install:
	@mv spl /usr/bin;
