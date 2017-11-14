.PHONY: default serve run

default: serve

serve:
	gin --port 3151 --appPort 3150

run:
	go build main.go && ./main
