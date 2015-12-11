# spl

spl is a small command-line tool for splitting strings using a substring as a delimiter.
The input source can be either STDIN or a file.

### Installation

```go get github.com/alex-ant/spl```

### Usage

Reading from file:
```sh
$ cat test.txt
SaaaAaaaC
PaaaD
LaaaBaaaE
$ spl -d aaa -p 1 test
S
P
L
$ spl -d aaa -p 2 test
A
D
B
$ spl -d aaa -p 3 test
C

E
```

Reading from STDIN:
```sh
$ echo "string 12ab to 12ab split" | spl -d 12ab -p 1
string
$ echo "string 12ab to 12ab split" | spl -d 12ab -p 2
 to
$ echo -e "string 12ab \nto 12ab split" | spl -d 12ab -p 1
string
to
```
