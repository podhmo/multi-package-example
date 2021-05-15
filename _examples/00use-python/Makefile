PYTHON ?= .venv/bin/python

default: $(PYTHON)
	$(PYTHON) main.py

.venv/bin/python:
	python -m venv .venv
	.venv/bin/pip install -r requirements.txt

