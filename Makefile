.PHONY: help testdata

help:
	@echo "testdata - format, lint and generate code from testdata"

testdata:
	cd testdata && \
		rm -rf gen && \
		buf format --write && \
		buf lint && \
		buf generate
