# CAST AI cli (Beta)

NOTE: CAST AI CLI is in it's early stage. Feel free to contribute, ask questions, give feedback.

![demo](https://raw.githubusercontent.com/castai/cast-cli/master/cluster_create.gif)

## Installation

### macOS

`cast` is available via [Homebrew][], and as a downloadable binary from the [releases page][].

#### Homebrew

| Install:          | Upgrade:          |
| ----------------- | ----------------- |
| `brew install castai/tap/cli` | `brew upgrade castai/tap/cli` |

### Linux

`cast` is available via [Homebrew](#homebrew), and as downloadable binaries from the [releases page][].

#### Homebrew

| Install:          | Upgrade:          |
| ----------------- | ----------------- |
| `brew install castai/tap/cli` | `brew upgrade castai/tap/cli` |

### Windows

`cast` is available as a downloadable binary from the [releases page][].

## Getting started

After installing CLI you need to configure API access token to access CAST AI public API.

### Quick configuration

```
cast configure
```

After done configuration file is saved to file system.
	
### Configure via environment variables
It is possible to override all configuration with environment variables.

| Variable          | Description          
| ----------------- | ----------------- |
| CASTAI_API_TOKEN | API access token |
| CASTAI_API_HOSTNAME | API hostname |
| CASTAI_DEFAULT_REGION | Default region for cluster creation |
| CASTAI_DEBUG | Enable debug mode | 
| CASTAI_CONFIG | Custom path to CLI configuration file |

## Usage

Run `cast` without any arguments to get help. Use --help on sub commands to get more help, eg. `cast cluster --help` 

```
CAST AI Command Line Interface

Usage:
  cast [command]

Available Commands:
  cluster     Manage clusters
  completion  Generate completion script
  configure   Setup initial configuration
  credentials Manage credentials
  help        Help about any command
  node        Manage clusters nodes
  region      Manage regions
  version     Print version

Flags:
  -h, --help   help for cast

Use "cast [command] --help" for more information about a command.
```

#### Autocompletion

Run complection command help and follow instructions.
```
cast completion --help
```

#### Create and access cluster

1. Create cluster
```
cast cluster create --name=my-cluster --credentials=aws --progress
```

2. List clusters

```
cast cluster list
```

3. Setup kubeconfig
```
cast cluster get-kubeconfig my-cluster-name
```

4. Verify Kubernetes nodes
```
kubectl get nodes
```

#### Add nodes

Interactive move
```
cast node add
```

Declarative move
```
cast -c=cluster-name node add --cloud=gcp --role=worker --shape=medium
```

#### Delete node

Interactive move
```
cast node delete
```

Declarative move
```
cast node add
```

#### Connect to node via SSH

Interactive move
```
cast node ssh
```

Declarative mode

```
cast -c=cluster-name node ssh my-node-name-123
```


[Homebrew]: https://brew.sh
[releases page]: https://github.com/castai/cli/releases/latest
