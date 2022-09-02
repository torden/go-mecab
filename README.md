# go-mecab

[![Build Status](https://travis-ci.org/torden/go-mecab.svg?branch=develop)](https://travis-ci.org/torden/go-mecab)
[![Go Report Card](https://goreportcard.com/badge/github.com/torden/go-mecab)](https://goreportcard.com/report/github.com/torden/go-mecab)
[![GoDoc](https://godoc.org/github.com/torden/go-mecab?status.svg)](https://godoc.org/github.com/torden/go-mecab)
[![Coverage Status](https://coveralls.io/repos/github/torden/go-mecab/badge.svg?branch=develop)](https://coveralls.io/github/torden/go-mecab?branch=master)
[![GitHub version](https://badge.fury.io/gh/torden%2Fgo-mecab.svg)](https://badge.fury.io/gh/torden%2Fgo-mecab)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


# Install

You need to tell Go where MeCab has been installed.

```bash
CGO_CFLAGS=-I/path/to/include -I./ CGO_LDFLAGS=-L/path/to/lib -lmecab -lstdc++ -Wl,-rpath,/path/to/lib -lmecab go get github.com/torden/go-mecab
```

If you installed mecab-config, execute following comands.

```bash
CGO_CFLAGS="`mecab-config --cflags` -I./" CGO_LDFLAGS="`mecab-config --libs` -Wl,-rpath,`mecab-config --libs-only-L`" go get github.com/torden/go-mecab
```

If you installed mecab pkg in this repository, execute following comands.
```bash
go get github.com/torden/go-mecab
```


# Links

- [MeCab's GitHub Page](http://taku910.github.io/mecab/)
- [MeCab's Repository](https://github.com/taku910/mecab)
- [MeCab-Ko's Repository](https://bitbucket.org/eunjeon/mecab-ko)
- [MeCab-Ko's Dictionary Repository](https://bitbucket.org/eunjeon/mecab-ko-dic)


  ---

*Please feel free. I hope it is helpful for you*
