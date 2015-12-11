# spl

spl is a small command-line tool for splitting strings using a substring as a delimiter.

### Features
* The input source can be either STDIN or a file.
* You may request to output one or more parts for each input line.
* A custom output delimiter can be set.

### Installation

```go get github.com/alex-ant/spl```

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

$ echo -e "STRINGabcTOabc11 \n abcSPLITabc22" | ./spl -d abc -p 2,3 -od /
TO/11
SPLIT/22
```
