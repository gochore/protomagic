.PHONY: help testdata

help:
	@echo "testdata - format, lint and generate code from testdata"

testdata:
	cd testdata && \
		buf format --write && \
		buf lint && \
		buf generate
