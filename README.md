## Development

Ensure you have `go 1.9.2` (or better) environment set up.
Also make sure `GOROOT`, `GOPATH` and `PATH` environment variables are set to their appropriate values.

For example on `CentOS 7 or Ubuntu 16` add the following lines to your `.bashrc`

```bash
export GOROOT="/usr/lib/go-1.9"
export GOPATH="$HOME/go"
export PATH="$PATH:$GOROOT/bin:$GOPATH/bin"
```

```bash
$ make code-gen
./hack/update-codegen.sh

```