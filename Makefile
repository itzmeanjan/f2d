graphql_gen:
	pushd app
	gqlgen generate
	popd

build:
	go build -o f2d

run: build
	./f2d
