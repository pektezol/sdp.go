# sdp.go [![Go Reference](https://pkg.go.dev/badge/github.com/pektezol/sdp.go.svg)](https://pkg.go.dev/github.com/pektezol/sdp.go) [![Go Report Card](https://goreportcard.com/badge/github.com/pektezol/sdp.go)](https://goreportcard.com/report/github.com/pektezol/sdp.go) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/pektezol/sdp.go/blob/main/LICENSE)

Source Demo Parser written in Golang.
Currently supports Portal 2.

## Couldn't Do This Without Them

- [@UncraftedName](https://github.com/UncraftedName): For [UntitledParser](https://github.com/UncraftedName/UntitledParser)
- [@NeKzor](https://github.com/NeKzor): For [dem.nekz.me](https://dem.nekz.me)

## Usage

```bash
$ ./parser demo.dem

$ ./parser demos/
```

## Currently Supports

- File or folder input using the CLI.
- API calls over library import.
- Either pretty text or JSON output.
- Parsing of demo headers and each (most) message type.
- Parsing of packet classes.
- Custom injected SAR data parsing.
- Multi-threaded parsing over multiple demos.

## TODO

- In-depth packet class parsing for each class. ([#7][i7])
- Entity parsing. ([#17][i17])

[i7]: https://github.com/pektezol/sdp.go/issues/7
[i17]: https://github.com/pektezol/sdp.go/issues/17