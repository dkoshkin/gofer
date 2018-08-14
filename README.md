# Gofer

A CLI utility to help you keep your project's ever-changing dependency versions up to date.

## The Name
According to the dictionary: [gofer](https://www.merriam-webster.com/dictionary/gofer) - an employee whose duties include running errands 

The tool is also written in Go, a language with a recognizable [gopher mascot](https://blog.golang.org/gopher)

## Motivation

For the last 2 years I maintained the [kismatic](https://github.com/apprenda/kismatic) project, a Kubernetes cluster lifecycle management tool.

As the project grew the process of checking for new docker image versions and Github project releases became pretty time consuming. The idea for a tool that would automate that process was born.

## Prerequisites

### Using a Docker container

* `docker` installed
* if desired add command to your path:
```
echo 'docker run                                    \
  -v $(pwd):/gofer                                  \
  -e GOOGLE_ACCESS_TOKEN=$GOOGLE_ACCESS_TOKEN       \
  dkoshkin/gofer:stable $@ ' > /usr/local/bin/gofer \
&& chmod +x /usr/local/bin/gofer
```

### Using a binary

* a Linux or macOS
```
# for Linux
wget https://github.com/dkoshkin/gofer/releases/download/$VERSION/gofer-linux-amd64 -O /usr/local/bin/gofer
# for macOS
wget https://github.com/dkoshkin/gofer/releases/download/$VERSION/gofer-darwin-amd64 -O /usr/local/bin/gofer
```

## Usage

1) Initialize an empty config file in `./.gofer/config.yaml`

```
gofer init -f ./.gofer/config.yaml
```

Will result in `./.gofer/config.yaml`:

```
apiversion: v0.1
dependencies: []
```

2) Add a new docker dependency with an optional `--mask` (a regular expression) and `--type`.   
The `--type` will be inferred from the `name` but can be set explicitly or set to `manual` to for unsupported dependencies.  
*Note* when set to `manual` the `dig` command will skip the dependency when fetching latest version.

```
gofer add busybox 1.28.1 --mask "1.28.[0-9]+" --type docker
```

Will result in `./.gofer/config.yaml`:

```
apiversion: v0.1
dependencies:
- name: busybox
  type: docker
  version: 1.28.1
  mask: 1.28.[0-9]+
```

**IMPORTANT when fetching versions for gcr.io docker images set:** 
```
export GOOGLE_ACCESS_TOKEN=`gcloud auth print-access-token`
```

3) Fetch the latest versions of all dependencies
```
gofer dig
```

Will result in `./.gofer/config.yaml`:

```
apiversion: v0.1
dependencies:
- name: busybox
  type: docker
  version: 1.28.1
  latestVersion: 1.28.4
  mask: 1.28.[0-9]+
```

### Example
A more complete `config.yaml` example available [here](https://raw.githubusercontent.com/dkoshkin/gofer/master/examples/config.yaml).

## Supported Dependencies

* docker - depending on the implementation of the registry the mechanism to provide the credentials will be different
  * [dockerhub](https://hub.docker.com/)
  * gcr.io
  * quay.io
  * private registry - (in progress)
* [github](https://github.com/) - (in progress)
* *manual* - the `dig` command will skip this dependency when fetching latest version

## Development

```
make vendor
make test
# build a docker container
make build-container
# or build a binary to bin/
make build-binary
```