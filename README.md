# env-vault

A simple tool to store env vars in keyring and execute programs with them injected.

> This tool is in early development, the api is likely to change

There are some of solutions to store secrets and inject them as env var during runtime like [vault](https://www.vaultproject.io/). Tho i like the idea of storing them into a single source of truth, i do not like managing a server instance for my private projects. I normally store API keys etc. in keepass and copy them into my terminal session. This tool makes it easier for me to manage my env var secrets. It is very similar to [99designs/aws-vault](https://github.com/99designs/aws-vault).

## Install

I will not provide binaries in the reposiories release section. Bare in mind, this tool manages your secrets, you will not want precompiled binaries to see them. Compile it yourself or use tools like [gobinaries.com](https://gobinaries.com/) which will download the source code and compile it on the fly (if you trust them).  

It's recommended to use `go install`. This will download and compile the program: 

```sh
go install github.com/janstuemmel/env-vault
```

If you trust services like gobinaries, type: 

```sh
curl -sf https://gobinaries.com/janstuemmel/env-vault | sh
```

## Usage

### Create

Creates a new environment

```
env-vault create my-project
```

### Add

Adds a var to environment

```
env-vault add my-project MY_SECRET=s3cret
```

### Remove Env

Removes a environment

```
env-vault remove my-project
```

### Remove var

Removes a environment

```
env-vault remove my-project MY_SECRET
```

### List Envs

List all environments

```
env-vault list
```

### List Vars

List all vars in environment

```
env-vault list my-project
```

### Exec

Execute command with env injected

```
env-vault exec my-project env
```