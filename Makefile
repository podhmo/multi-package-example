VERSION = 0.0.0

publish-go:
	git tag -f go/foo/v$(VERSION)
	git tag -f go/bar/v$(VERSION)
	git push --tag
.PHONY: publish-go

publish-python:
	echo nothing to do
.PHONY: publish-python
