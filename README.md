# Markdown internal & external links validation (MILV)

`MILV` is a bot that parses, checks and validates internal & external URLs links in markdown files. It can be used for [verification pull requests](#validate-pull-requests).

## Installation

    $ go get github.com/magicmatatjahu/milv

This library use [`gopkg.in/yaml.v2`](https://github.com/go-yaml/yaml), so if you haven't it in your `GOPATH` use `go get gopkg.in/yaml.v2` or `dep ensure -vendor-only`

## Validate Pull Requests

### Jenkins

To use `MILV` with Jenkins, connect your repo and create a [`Jenkinsfile`](https://jenkins.io/doc/book/pipeline/jenkinsfile/#creating-a-jenkinsfile) and add stage:

```groovy
stage("validate internal & external links") {
    workDir = pwd()
    sh "docker run --rm -v $workDir:/milv/mds magicmatatjahu/milv:latest -docker=true"
}
```

## Contact

- [github.com/magicmatatjahu](https://github.com/magicmatatjahu)

## License

This project is available under the MIT license. See the [LICENSE](LICENSE) file for more info.