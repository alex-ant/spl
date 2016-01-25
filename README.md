# spl

[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](https://gobuild.io/alex-ant/spl/master) [![Go Report Card Badge](http://goreportcard.com/badge/alex-ant/spl)](http://goreportcard.com/report/alex-ant/spl)

A small command-line tool for splitting strings using a substring as a delimiter.

### Features
* The input source can be either STDIN or a file
* You may request to output one or more parts for each input line
* A custom output delimiter can be set
* Provided delimiter can automatically be multiplied while splitting (so splitting string "123aaa456" by "a" with a flag -s you'll get two parts "123" and "456")

### Installation

```
go get github.com/alex-ant/spl
```
Or in case you don't have Go installation, just follow the [gorelease](https://gobuild.io/alex-ant/spl/master) link to download binaries for your OS.

### Usage

Reading from file:
```sh
$ cat test.txt
SaaaAaaaC
PaaaD
LaaaBaaaE

$ spl -d aaa -p 1 test.txt
S
P
L

$ spl -d aaa -p 2 test.txt
A
D
B

$ spl -d aaa -p 3 test.txt
C

E

$ spl -d aaa -p 1,2 test.txt
S,A
P,D
L,B

$ spl -d aaa -p 1,2,3 -od '|' test.txt
S|A|C
P|D|
L|B|E

$ spl -d aaa -p 3,2,1 -od '|' test.txt
C|A|S
|D|P
E|B|L

$ spl -d a -p 2 -s test.txt
A
D
B
```

Reading from STDIN:
```sh
$ echo "STRINGabcTOabcSPLIT" | spl -d abc -p 1
STRING

$ echo "STRINGabcTOabcSPLIT" | spl -d abc -p 2
TO

$ echo -e "STRINGabcTO \n abcSPLIT" | spl -d abc -p 2
TO
SPLIT

$ echo -e "STRINGabcTOabc11 \n abcSPLITabc22" | spl -d abc -p 2,3 -od /
TO/11
SPLIT/22

$ echo "a    b c   d" | spl -d ' ' -p 2 -s
b

$ echo "a    b c   d" | spl -d ' ' -p 3 -s
c
```
