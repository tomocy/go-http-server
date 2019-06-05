# server
a command to serve static files

### Installation
```
go install github.com/tomocy/server
```

### Usage
```
server:
  -addr string
    	address of server (default ":80")
  -root string
    	root of static files to be served (default "./")
```

### Example
Taking [tomocy.github.io](https://github.com/tomocy/tomocy.github.io) as an example,
```
$ ls
README.md	css		image		index.html	js sitemap.xml
# start serving files in current directory on :80
$ server
```
http://localhost can be accessed.  