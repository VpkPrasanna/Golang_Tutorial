to build a package as a executable , we do that with the help of golang easily 

- to build a package which is compatible for windows

```python
GOOS="windows" go build 
```

to build a package which is compatible for linux

```python
GOOS="linux" go build 
```

to build a package which is compatible for Mac

```python
GOOS="darwin" go build 
```
To find the list of possible platforms  that golang supports

```python
go tool dist list
```