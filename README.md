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
| `-allow-redirect` | redirects will be allowed | `false`
| `-request-repeats` | number of request repeats | `1`
| `-allow-code-blocks` | checking links in code blocks | `false`
| `-timeout` | connection timeout (in seconds) | `30`
| `-ignore-external` | ignore external links | `false`
| `-ignore-internal` | ignore internal links | `false`
| `-v` | enable verbose logging | `false`
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
docker run --rm -v $PWD:/milv:ro magicmatatjahu/milv:stability -base-path=/milv
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

#### Advanced configuration

> **NOTE**: For this example tree of project is the same as above.

Config file can look like this:

```yaml
white-list-external: ["localhost", "abc.com"]
white-list-internal: ["LICENSE"]
black-list: ["./README.md"]
request-repeats: 5
timeout: 45
allow-redirect: false
allow-code-blocks: true
files:
  - path: "./src/foo.md"
    config:
      white-list-external: ["google.com"]
      white-list-internal: ["#contributing"]
      request-repeats: 3
      timeout: 30
      allow-code-blocks: false
    links:
      - path: "https://github.com/magicmatatjahu/milv"
        config:
          timeout: 15
          allow-redirect: true
```

In this example we can see that `milv` will globally check external links with 45 seconds timeout, also won't allow redirect and will allow checking links in code snippets and default times of request repeats is set 5.

`Milv` also allows to separately configurate files. Timeout in `./src/foo.md` file will be set to 30 seconds, links will be checking 3 times (if they will return error) and the links in code blocks won't be checked. However, a single link `https://github.com/magicmatatjahu/milv` will be checking with 15 seconds timeout with the possibility of redirection.

## Troubleshooting links

The below table describes the types of errors during checking links and examples of how to solve them:

| Error        | Solution example   |
|-------------|----------------|
| `404 Not Found` | Page doesn't exist - you have to change the external link to the correct one |
| Error with formatting link | Correct link or if link has a variables or it is a example, add this link to the `white-list-external` or `white-list-internal` |
| `The specified file doesn't exist` | Change the relative path to the file to the correct one or use a absolute path (second solution is not recommended) |
| `The specified header doesn't exist in file` | Change the anchor link in `.md` file to the correct one. Sometimes `milv` give a hint (`Did you mean about <similar header>?`) of which header (existing in the file) is very similar to the given. |
| `The specified anchor doesn't exist...` or `The specified anchor doesn't exist in website...` | Check which anchors are on the external website and correct the specified anchor or remove the redirection to the given anchor. Sometimes `milv` give a hint (`Did you mean about <similar anchor>?`) of which anchor (existing in the website) is very similar to the given. |
| `Get <external link>: net/http: request canceled (Client.Timeout exceeded while awaiting headers)` | Increase net timeout to the all files, specific file or specific link or increase times of request repeats ([Here's](#advanced-configuration) how to do it) |
| `Get <external link>: EOF ` | Same as above or change the link to the other one (probably website doesn't exist) |
| Other types of errors and errors with contains `no such host` or `timeout` words | Most likely, the website doesn't exist or you do not have access to it. Possible solutions: change the link to another, correct one, remove it or add it to the `white-list-external` or `white-list-internal` |

It is a good practice to add local or internal (in the local network) links to the global white list of external or internal links, such as `http://localhost`.

## Validate Pull Requests

`milv` can help you validate links in all `.md` files in whole repository when a pull request is created (or a commit is pushed).

### Jenkins

To use `milv` with Jenkins, connect your repo and create a [`Jenkinsfile`](https://jenkins.io/doc/book/pipeline/jenkinsfile/#creating-a-jenkinsfile) and add stage:

```groovy
stage("validate internal & external links") {
    workDir = pwd()
    sh "docker run --rm -v $workDir:/milv:ro magicmatatjahu/milv:0.0.4 -base-path=/milv"
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

* [ ] error handling 
* [ ] refactor (new architecture)
* [ ] documentations
* [ ] possibility to validation remote repositories hosted on **GitHub**
* [ ] parse other type of files
* [x] add more commands like a: timeout for http.Get(), allow redirects or SSL
* [ ] landing page for project