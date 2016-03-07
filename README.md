# daqdata

[![GoDoc][godoc image]][godoc link]
[![License Badge][license image]][LICENSE.txt]

Go library to read/write binary data and header information from a DAQ,
such as the [USB-1608FS-Plus][]. The binary data is stored in its raw
format in one or more binary files. The header file contains information
about the binary data such as calibration data, number of channels, etc.

## Installation

```bash
$ go get github.com/gotmc/daqdata
```

## Dependencies

At this time there are no dependencies beyond the [Go standard
library][go-std].

## Documentation

Documentation can be found at either:

- <https://godoc.org/github.com/gotmc/daqdata>
- <http://localhost:6060/pkg/github.com/gotmc/daqdata/> after running `$
  godoc -http=:6060`

## Contributing

[daqdata][] is developed using [Scott Chacon][]'s [GitHub Flow][]. To
contribute, fork [daqdata][], create a feature branch, and then
submit a [pull request][].  [GitHub Flow][] is summarized as:

- Anything in the `master` branch is deployable
- To work on something new, create a descriptively named branch off of
  `master` (e.g., `new-oauth2-scopes`)
- Commit to that branch locally and regularly push your work to the same
  named branch on the server
- When you need feedback or help, or you think the branch is ready for
  merging, open a [pull request][].
- After someone else has reviewed and signed off on the feature, you can
  merge it into master.
- Once it is merged and pushed to `master`, you can and *should* deploy
  immediately.

## Testing

Prior to submitting a [pull request][], please run:

```bash
$ gofmt
$ golint
$ go vet
$ go test
```

To update and view the test coverage report:

```bash
$ go test -coverprofile coverage.out
$ go tool cover -html coverage.out
```

## License

[daqdata][] is released under the MIT license.  Please see the
[LICENSE.txt][] file for more information.

[GitHub Flow]: http://scottchacon.com/2011/08/31/github-flow.html
[go-std]: https://golang.org/pkg/
[godoc image]: https://godoc.org/github.com/gotmc/mccdaq?status.svg
[godoc link]: https://godoc.org/github.com/gotmc/mccdaq
[jasper]: https://textiles.ncsu.edu/blog/team/warren-jasper/
[libusb]: https://github.com/gotmc/libusb
[libusb-c]: http://libusb.info
[LICENSE.txt]: https://github.com/gotmc/mccdaq/blob/master/LICENSE.txt
[license image]: https://img.shields.io/badge/license-MIT-blue.svg
[mccdaq]: https://github.com/gotmc/mccdaq
[mcc-linux]: http://www.mccdaq.com/daq-software/Linux-Support.aspx
[pull request]: https://help.github.com/articles/using-pull-requests
[Scott Chacon]: http://scottchacon.com/about.html
[usb-1608fs-plus]: http://www.mccdaq.com/usb-data-acquisition/USB-1608FS-Plus.aspx
