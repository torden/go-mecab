# go-mecab

Go(golang) Binding to the mecab-ko

----

[![Build Status](https://github.com/torden/go-mecab/actions/workflows/go-mecab.yml/badge.svg)](https://github.com/torden/go-mecab/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/torden/go-mecab)](https://goreportcard.com/report/github.com/torden/go-mecab)
[![GoDoc](https://godoc.org/github.com/torden/go-mecab?status.svg)](https://godoc.org/github.com/torden/go-mecab)
[![codecov](https://codecov.io/gh/torden/go-mecab/branch/master/graph/badge.svg?token=04152a42-5140-4337-b82b-c50655ada485)](https://codecov.io/gh/torden/go-mecab)
[![Coverage Status](https://coveralls.io/repos/github/torden/go-mecab/badge.svg?branch=master)](https://coveralls.io/github/torden/go-mecab?branch=master)
[![CodeQL](https://github.com/torden/go-mecab/workflows/CodeQL/badge.svg)](https://github.com/torden/go-mecab/actions/workflows/codeql-analysis.yml)
[![GitHub version](https://img.shields.io/github/v/release/torden/go-mecab)](https://github.com/torden/go-mecab)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


# Install

You need to tell Go where MeCab has been installed.

## Install from Pre-build Packages

- [Ubuntu Package](https://github.com/torden/go-mecab/tree/develop/pkg/ubuntu)
- [RedHat(CentOS) Package](https://github.com/torden/go-mecab/tree/develop/pkg/rhel)


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
