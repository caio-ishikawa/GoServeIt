GoServeIt v1.0.0
----------------
<img src="https://i.imgur.com/L5sXufv.png">

GoServeIt is a zero-dependency command-line tool that quickly starts a barebones HTTP server to serve files in localhost. The main usecase is to quickly serve scripts for CTF challenges. 

Installation
---------------
### Using ```go install```
If you have a Go environment ready to go, it's as easy as:

```sh
go install github.com/caio-ishikawa/GoServeIt/@1.0.0
```

### Building From Source
This tool is written in [Go](https://golang.org/), and you will need to install the Go language/compiler/toolkit if you don't already have it. Full details of installation and set up can be found [on the Go language website](https://golang.org/doc/install). Once installed you can run the following command:

```sh
git clone https://github.com/caio-ishikawa/goserveit.git && cd goserveit && make build
```

### Uninstalling
To uninstall GoServeIt, you can navigate to the cloned repository (or clone it if yo have deleted it), and run:

```sh
make uninstall
```

Getting Started
---------------
You can start a GoServeIt server by typing 'gsi' in your terminal followed by the relevant flags:
- -p (Optional): Specifies the port on which the HTTP server will listen (defaults to 8080).
- -f (Required): Specifies the name of the file to serve (e.g., -f test.sh).
- -v (Optional): If set, it will dump request information for visibility.

### Example:
```sh
gsi -p 9000 -f test.sh -v
```

