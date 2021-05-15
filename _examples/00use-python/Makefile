PYTHON ?= .venv/bin/python

default: $(PYTHON)
	$(PYTHON) -m pip install -r requirements.txt
	$(PYTHON) main.py

updated: $(PYTHON)
	$(PYTHON) -m pip install -U -r requirements.txt
	$(PYTHON) main.py

.venv/bin/python:
	python -m venv .venv
	.venv/bin/pip install -r requirements.txt

