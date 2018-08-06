# Gofer

A CLI utility to help you keep your project's ever-changing dependency versions up to date.

## The Name
According to the [dictionary](https://www.merriam-webster.com/dictionary/gofer): *gofer* - an employee whose duties include running errands 

The tool is also written in Go, a language with a recognizable *gofer* [mascot](https://blog.golang.org/gopher)

## Motivation

For the last 2 years I maintained the [kismatic](https://github.com/apprenda/kismatic) project, a Kubernetes cluster lifecycle management tool.

As the project grew the process of checking for new docker image versions and Github project releases became pretty time consuming. The idea for a tool that would automate that process was born.

## Usage

1) Initialize an empty config file in `./gofer/config.yaml`

```
docker run -v $(pwd):/gofer dkoshkin/gofer init -f ./.gofer/config.yaml
```

Will result in `./gofer/config.yaml`:

```
apiversion: v0.1
dependecies: []
```

2) Add a new docker dependency with an optional `--mask` (a regular expression) and `--type`.   
The `--type` will be inferred from the `name` but can be set explicitly or set to `manual` to for unsupported dependencies.  
*Note* when set to `manual` the `dig` command will skip the dependency when fetching latest version.

```
docker run -v $(pwd):/gofer dkoshkin/gofer add busybox 1.28.1 --mask "1.28.[0-9]+" --type docker
```

Will result in `./gofer/config.yaml`:

```
apiversion: v0.1
dependecies:
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
docker run -v $(pwd):/gofer -e GOOGLE_ACCESS_TOKEN=$GOOGLE_ACCESS_TOKEN dkoshkin/gofer dig
```

Will result in `./gofer/config.yaml`:

```
apiversion: v0.1
dependecies:
- name: busybox
  type: docker
  version: 1.28.1
  latestVersion: 1.28.4
  mask: 1.28.[0-9]+
```

### Example
A more complete `config.yaml` example available [here](https://raw.githubusercontent.com/dkoshkin/gofer/master/config.yaml.sample).

## Development

```
make vendor
make test
make build
```