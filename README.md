# docker-code-link

cli tool for attaching containers directly to vscode

## Dependencies

- Remote Development (VS Code Extension)

VS Marketplace link: https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack

## Usage

The basic usage is as follows:

```
# Display a list of containers
docker-code-link list

# Attach to a specific container
docker-code-link attach [container-name]
```

### Bind Keys

`bash`

```.bashrc
bind -x '"\C-p":"/path/to/docker-code-link list\n"'
```

`zsh`

```.zshrc
bindkey -s "^P" "/path/to/docker-code-link list\n"
```

## Features

- Display a list of Docker containers

- Attach to a selected container using VS Code

## Building from Source

This requires go >= 1.20

If you wish to build docker-code-link from source for your specific platform, follow these steps:

1. Clone the Repository:

```
git clone https://github.com/mizunara-tech/docker-code-link.git
cd docker-code-link
```

2. Build the Binary:

For Linux/Mac:

```
GOOS=linux GOARCH=amd64 go build -o docker-code-link ./cmd/docker-code-link
```

For Windows:

```
GOOS=windows GOARCH=amd64 go build -o docker-code-link.exe ./cmd/docker-code-link
```

Replace GOOS and GOARCH values as per your requirement.

## License

This project is provided under the MIT License.

## Contributing

If you would like to contribute, please send a pull request.
