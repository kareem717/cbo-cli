# K7-CBO: Personal CBO AI Agent

K7-CBO is an open-source AI agent that acts as your personal Chief Business Officer, helping with lead prospecting, cold outreach, and various business development tasks.

## Features

- 🎯 Lead Prospecting
- 📧 Cold Outreach Automation
- 💼 Business Development Tools
- 🔒 Local Data Storage
- 🤖 AI-Powered Decision Making

## Installation

### Prerequisites

- GCC (for SQLite support)
- CGO enabled

### Build from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/k7-cbo.git
cd k7-cbo

# Install dependencies
make install

# Build the binary
make build
```

### System-specific Requirements

#### Linux

```bash
sudo apt-get install gcc
```

#### macOS

```bash
xcode-select --install
```

#### Windows

Install MinGW or MSYS2 for GCC support

### Common Commands

```bash
# Get help
k7-cbo --help

# View version
k7-cbo version

# Start lead prospecting
k7-cbo prospect
```

## Configuration

K7-CBO stores its configuration and data in:

- Linux: `~/.k7-cbo/`
- macOS: `~/Library/Application Support/k7-cbo/`
- Windows: `%APPDATA%\k7-cbo\`

## Development

```
...TBA
```

### Project Structure

```
...TBA
```

### Adding New Commands

```go
func init() {
    rootCmd.AddCommand(newCommand)
}
```

### Running Tests

```bash
make test
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feat/some-command`)
3. Commit your changes (`git commit -m 'feat: add some-command'`) [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
4. Push to the branch (`git push origin feat/some-command`)
5. Open a Pull Request

## Future Plans

- Hosted version with additional features
- API integration capabilities
- Advanced AI models support
- Team collaboration features

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- Create an issue for bug reports
- Start a discussion for feature requests
- Join our community (Coming soon)

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra)
- Built with [Goose](https://github.com/pressly/goose)
- SQLite for local storage
- Special thanks to all contributors
