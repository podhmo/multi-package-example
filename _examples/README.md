```console
$ make -n

# python
make -C 00use-python/
.venv/bin/python -m pip install -r requirements.txt
.venv/bin/python main.py

# go
make -C 01use-go/
go run main.go
```
