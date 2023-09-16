# DemoParser [![Go Report Card](https://goreportcard.com/badge/github.com/pektezol/demoparser)](https://goreportcard.com/report/github.com/pektezol/demoparser) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/pektezol/DemoParser/blob/main/LICENSE)

Work-In-Progress demo parser for Portal 2 written in Golang.

Usage:

```bash
$ ./parser demo.dem
```

Currently supports:

- File or folder input using the CLI.
- Basic parsing of demo headers each message type.
- Basic parsing of packet classes.
- Custom injected SAR data parsing.

TODO:

- StringTableEntry parsing.
- In-depth packet class parsing for each class.
