# Markdown internal & external links validation

`MILV` is a bot that parses, checks and validates internal & external URLs links in markdown files. It can be used for [verification pull requests](#validate-pull-requests) and also as standalone library.

## Installation

```bash
$ go get -u -v github.com/magicmatatjahu/milv
```

For the above command to work you must have [GoLang](https://golang.org/doc/install) installed

After installation, run the program with `milv` from anywhere in the file system.

### Run the code from source

If you want run the code without installation, run the following commands to get the source code, resolve external dependencies and build the project. 

For this operations you must have also installed package manager [Dep](https://github.com/golang/dep).

```bash
git clone https://github.com/magicmatatjahu/milv.git
cd milv
dep ensure
go build
```

## Usage

### Command line parameters

You can use the following parameters while using `milv` binary:

| Name        | Description    | Default Value       |
|-------------|----------------|---------------------|
| `-base-path` | root directory of repository | `""`
| `-config-file` | configuration file for bot. See [more](#config-file) | `milv.config.yaml`
| `-white-list-ext` | comma separate external links which will not be checked | `[]`
| `-white-list-int` | comma separate internal links which will not be checked  | `[]`
| `-black-list` | comma separate files which will not be checked | `[]`
| `-help` or `-h` | Show available parameters | n/a

Files to be checked are given as free parameters.

### Examples

* Checks all links, without matching `github.com` in external links, in `.md` files in current directory+subdirectories without files matching `vendor` in path:

```bash
milv -black-list="vendor" -white-lis-ext="github.com"
```

* Checks links only in `./README.md` and `./foo/bar.md` files:

```bash
milv ./README.md ./foo/bar.md
```

### Docker image

If you do not want to install `milv` and it's dependencies you can simply use Docker and Docker image:

```bash
docker run --rm -v $PWD:/milv/mds magicmatatjahu/milv:latest -base-path=./mds
```

## Config file

The configuration file allows for quick parameterization of the `milv` works. Config file must be a `.yaml` file.

Parameterization is very similar to using parameters in the `CLI`. However, you can configure files, located in subdirectories relative to the configuration file, separately with different config.

### Examples

If your tree of your project look like this:

```
├── README.md
├── LICENSE
├── main.go
├── milv.config.yaml
└── src
    ├── file.go
    ├── file_test.go
    ├── foo.md
    └── some_dir
        └── bar.md
```

your config file can look like this:

```yaml
white-list-external: ["localhost", "abc.com"]
white-list-internal: ["LICENSE"]
black-list: ["./README.md"]
files:
  - path: "./src/foo.md"
    config:
      white-list-external: ["github.com"]
      white-list-internal: ["#contributing"]
```

Before run validation, `milv` remove from files list `./README.md` file to check and connect global `white-list-external` with file `white-list-external` and `white-list-external` for `./src/foo.md` file will look that:

```yaml
white-list-external: ["localhost", "abc.com", "github.com"]
```

Similarly will be with `white-list-internal`.

If you have a config file and you use a `CLI`, then `milv` will automatically combine the parameters from file and consol.

## Validate Pull Requests

`milv` can help you validate links in all `.md` files in whole repository when a pull request is created (or a commit is pushed).

### Jenkins

To use `milv` with Jenkins, connect your repo and create a [`Jenkinsfile`](https://jenkins.io/doc/book/pipeline/jenkinsfile/#creating-a-jenkinsfile) and add stage:

```groovy
stage("validate internal & external links") {
    workDir = pwd()
    sh "docker run --rm -v $workDir:/milv/mds magicmatatjahu/milv:latest -base-path=./mds"
}
```

## Other validators

In opensource community is available other links validation libraries written in JS, Ruby and others languages. Here are a few of note:

* [awesome_bot](https://github.com/dkhamsing/awesome_bot): validator written in Ruby. Allows for validation external and internal links in `.md` files.
* [remark-validate-links](https://github.com/remarkjs/remark-validate-links): validator written in JS. Allows for validation internal links in `.md` files.

## Contact

- [github.com/magicmatatjahu](https://github.com/magicmatatjahu)

## Contributing

If you want contribute this project, firstly read [CONTRIBUTING.md](CONTRIBUTING.md) file for details of submitting pull requests.

## License

This project is available under the MIT license. See the [LICENSE](LICENSE) file for more info.

## ToDo

* [ ] possibility to validation remote repositories hosted on **GitHub**
* [ ] parse other type of files
* [ ] add more commands like a: timeout for http.Get(), allow redirects or SSL
* [ ] landing page for project