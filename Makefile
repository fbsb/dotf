build:
	docker build -t dotf .

shell: build
	docker run --rm -it -v $$(pwd)/test:/root dotf bash
