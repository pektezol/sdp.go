# DemoParser [![Go Report Card](https://goreportcard.com/badge/github.com/pektezol/demoparser)](https://goreportcard.com/report/github.com/pektezol/demoparser) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/pektezol/DemoParser/blob/main/LICENSE)

Work-In-Progress demo parser for Portal 2 written in Golang.

## Couldn't Do This Without Them

- [@UncraftedName](https://github.com/UncraftedName): For [UntitledParser](https://github.com/UncraftedName/UntitledParser)
- [@NeKzor](https://github.com/NeKzor): For [nekz.me/dem](https://nekz.me/dem)

## Usage

```bash
$ ./parser demo.dem
```

## Currently Supports

- File or folder input using the CLI.
- Basic parsing of demo headers each message type.
- Basic parsing of packet classes.
- Custom injected SAR data parsing.

## TODO

- StringTableEntry parsing. ([#8][i8])
- In-depth packet class parsing for each class. ([#7][i7])

[i8]: https://github.com/pektezol/DemoParser/issues/8
[i7]: https://github.com/pektezol/DemoParser/issues/7