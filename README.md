iggy
====

[![GoDoc Widget]][GoDoc]

Iggy is a tool to generate `.gitignore` files from the command line.
Templates can be generated for multiple languages and/or frameworks.

A list of supported languages/frameworks can be found [here](https://github.com/github/gitignore).


Installation
------------

```
$ go get github.com/while-loop/iggy
```

Usage
-----

```
iggy [-a] <language/framework>...

Usage of iggy:
  -a    append ignores to current file

```

Example
-------

```bash
$ iggy go jetbrains macos vim
```

License
-------
iggy is licensed under the MIT License.
See [LICENSE](LICENSE) for details.

Author
------

Anthony Alves


[GoDoc]: https://godoc.org/github.com/while-loop/iggy
[GoDoc Widget]: https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square