.PHONY: default serve exec

default: serve

serve:
	gin --port 3151 --appPort 3150

exec:
	go build main.go && ./main
