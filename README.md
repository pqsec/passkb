# passkb

## Paste your passwords safely anywhere
**passkb** is a "virtual keyboard", which will type any string after a small delay. Its primary use-case is to paste passwords into fields, which otherwise disable paste functionality.

Disabling paste functionality for password fields is a [questionable security practice][ncsc-let-paste], because it makes it harder to use password managers for these fields and encourages users to come up with insecure passwords. Since such password fields still need to allow users to type the password in, **passkb** emulates a virtual keyboard and "types" the provided password for you. So instead of pasting your secure password directly, you paste it into **passkb** and **passkb** "types" it in instead.

### Installation
**passkb** can be downloaded and compiled using standard `go get` approach. Assuming you have [Go](https://golang.org/doc/install) installed, just
```
go get github.com/pqsec/passkb/cmd/passkb
```
The `passkb` binary should appear in your `$GOPATH/bin` directory.

Works on Linux, Windows and MacOS (on MacOS *Command Line Tools for Xcode* are needed to compile the binary).

[ncsc-let-paste]: https://www.ncsc.gov.uk/blog-post/let-them-paste-passwords
