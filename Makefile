.PHONY: default serve

default: serve

serve:
	gin --port 3151 --appPort 3150
