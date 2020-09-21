set -e
set -x

cd /build
rm -fr /build/pkg-build *.rpm

go-bin-rpm generate -a amd64 --version 0.0.1 -b pkg-build/amd64/ -o pkg-build/tstwbsrv-amd64.rpm
