VERSION = 0.1.1

publish-go:
	git tag -f go/foo/v$(VERSION)
	git tag -f go/bar/v$(VERSION)
	git push -f --tag
.PHONY: publish-go

publish-python:
	echo nothing to do
.PHONY: publish-python

publish-root:
	git tag -f v$(VERSION)
.PHONY: publish-root
