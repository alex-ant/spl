# spl

![split it](https://raw.githubusercontent.com/alex-ant/spl/master/split.jpg "split it")

spl is a small command-line tool for splitting string passed over STDIN by another multicharacter string passed as a parameter.

### Installation

```sh
$ git clone git@github.com:alex-ant/spl.git
$ cd spl
$ make
$ sudo make install
```

Binaries are also available in the release section!

### Usage

```sh
$ echo "string 12ab to 12ab split" | spl -d 12ab -p 1
string
$ echo "string 12ab to 12ab split" | spl -d 12ab -p 2
 to
$ echo "string 12ab to 12ab split" | spl -d 12ab -p 3
 split
```
