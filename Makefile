initial_generate:
	go run cmd/colorsgenerate/main.go --baseFile cmd/colorsgenerate/colors.yml -o ./gen/colors

gogenerate:
	go generate -v ./...

