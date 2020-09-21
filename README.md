# test-http-server-aws-info
Minimal http(s) server for testing connections, that responds with AWS info

# Create RPM Amazon Linux2
```bash
docker build -t amazonlinux:2-go-rpmbuild .
GOOS=linux GOARCH=amd64 go build -o build/amd64/tstwbsrv main.go
docker run --rm -v ${PWD}:/build amazonlinux:2-go-rpmbuild sh /build/pkg-test.sh
```

Links
https://github.com/mh-cbon/go-bin-rpm
https://yourbasic.org/golang/http-server-example/
https://stackoverflow.com/questions/26090301/run-both-http-and-https-in-same-program
https://gist.github.com/denji/12b3a568f092ab951456
