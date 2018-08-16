# Cloning Repository

Skip any unnecessary steps if already complete.

```
# set GOPATH
$ export GOPATH=$HOME/go

#  make src directory
$ mkdir $GOPATH/src

# clone into src
$ git clone $GIT_URL $GOPATH/src/telnetgo

# move to project directory
$ cd $GOPATH/src/telnetgo

```

# Installing Dependencies

```
# Automatically
$ make deps

# Manually
$ go get github.com/reiver/go-oi
$ go get github.com/reiver/go-telnet
$ go get github.com/reiver/go-telnet/telsh
```
