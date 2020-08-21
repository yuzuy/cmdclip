# cmdclip
A CLI tool to clip output of a command

## Installation
```
go get -u github.com/yuzuy/cmdclip/cmd/cmdclip
```

## Usage
```
cmdclip echo 'hoge' // clip "hoge"
cmdclip cat main.go // clip content of main.go
```

### Flags
#### p - Print output of the command
```
cmdclip -p echo 'hoge' // clip "hoge" and print "hoge"
```
