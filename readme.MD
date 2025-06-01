# Termevent

A Go-based terminal event tracking system that monitors user input in a shell environment and triggers configurable actions based on predefined rules.

## Overview

Termevent creates an interactive terminal session that wraps around a shell (default: `sh`) and monitors all user input. When specific patterns are detected in the commands entered by the user, it can trigger various actions such as logging, notifications, or custom behaviors.

## Features

- **Real-time Input Monitoring**: Tracks all user input in the terminal session
- **Pattern Matching**: Configurable rules to match specific command patterns
- **Action System**: Extensible action framework (currently supports logging)
- **PTY Integration**: Full terminal emulation with proper signal handling
- **YAML Configuration**: Easy-to-configure rules via YAML files

## Installation

### Prerequisites

- Go 1.23.0 or later
- Unix-like operating system (Linux, macOS)

### Build from Source

```bash
git clone <repository-url>
cd termevent
go mod download
go build -o termevent
```

## Usage

### Basic Usage

Run the application:

```bash
./termevent
```

This will:

1. Display a welcome message
2. Start a new shell session
3. Monitor all commands entered
4. Apply configured rules to detect patterns
5. Execute actions when patterns match

### Configuration

Create or modify the `termevent.yml` file to define your rules:

```yaml
# Termevent Configuration
version: "1.0"
name: "termevent-config"

# Input Rules - Text patterns only
input_rules:
  - name: "ls detector"
    pattern: "ls"
    category: "text_pattern"
    type: "exact_match"
    actions:
      - name: "log"
        message: "ls command triggered"
```

#### Rule Structure

- **name**: Descriptive name for the rule
- **pattern**: The text pattern to match against user input
- **category**: Rule category (currently supports "text_pattern")
- **type**: Matching type (currently supports "exact_match")
- **actions**: List of actions to execute when pattern matches

#### Available Actions

- **log**: Logs a message when the pattern is detected
  ```yaml
  actions:
    - name: "log"
      message: "Custom log message here"
  ```

## Architecture

### Components

- **Terminal Package** (`terminal/`):

  - `term.go`: Main terminal creation and management
  - `ansi.go`: ANSI escape sequence handling
  - `chars.go`: Character constants

- **Rules Package** (`rules/`):
  - `rules.go`: Rule definition and configuration loading
  - `tracker.go`: Input tracking and pattern matching
  - `actions.go`: Action execution system

### How It Works

1. **Terminal Creation**: Creates a pseudo-terminal (PTY) and starts a shell
2. **Input Capture**: Intercepts all user keystrokes before they reach the shell
3. **Pattern Matching**: Compares completed commands against configured rules
4. **Action Execution**: Triggers associated actions when patterns match
5. **Shell Forwarding**: Forwards input to the actual shell for normal operation

## Development

### Project Structure

```
termevent/
├── main.go              # Application entry point
├── main2.go             # Additional functionality (ANSI demo)
├── go.mod               # Go module definition
├── termevent.yml        # Configuration file
├── terminal/            # Terminal management
│   ├── term.go
│   ├── ansi.go
│   └── chars.go
└── rules/               # Rule and action system
    ├── rules.go
    ├── tracker.go
    └── actions.go
```

### Dependencies

- `github.com/creack/pty`: PTY creation and management
- `golang.org/x/term`: Terminal control and raw mode
- `golang.org/x/sys`: System calls
- `gopkg.in/yaml.v3`: YAML configuration parsing

### Adding New Actions

To add a new action type:

1. Define the action struct in `rules/actions.go`
2. Implement the `ActionExecutor` interface
3. Add the action name constant
4. Update the `GetSpecificAction()` method

Example:

```go
type NotifyAction struct {
    Name    string `yaml:"name"`
    Title   string `yaml:"title"`
    Message string `yaml:"message"`
}

func (a *NotifyAction) Execute() {
    // Implementation here
}
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Roadmap

- [ ] Support for regex patterns
- [ ] Additional action types (notifications, webhooks, etc.)
- [ ] Command history analysis
- [ ] Performance metrics
- [ ] Web dashboard
- [ ] Plugin system
