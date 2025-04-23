<div align="center">
    <p>
        <picture>
            <img alt="dotx logo" src="./docs/assets/logo.svg" width="200">
        </picture>
    </p>
    <h1>dotx</h1>
    <p>
        <b>A modern dotfile manager for tracking and syncing configuration files</b>
    </p>
    <p>
        <a href="#installation">Installation</a> •
        <a href="#usage">Usage</a> •
        <a href="#commands">Commands</a> •
        <a href="#configuration">Configuration</a> •
        <a href="#license">License</a>
    </p>
    <hr>
    <p>
        <h3>⚠️ This project is still under development ⚠️</h3>
    </p>
</div>

## Overview

**dotx** helps you manage, version control, and synchronize your configuration files (dotfiles) across multiple systems. It provides a simple and intuitive CLI to:

- Track configuration files in a Git repository
- Deploy your dotfiles to new systems
- Synchronize changes across multiple machines
- Maintain a clean and organized dotfiles setup

## Installation

### Prerequisites

- Go 1.18 or higher
- Git

### From Source

```bash
# Clone the repository
git clone https://github.com/mxlang/dotx.git

# Navigate to the project directory
cd dotx

# Build and install
go install ./cmd/dotx
```

## Usage

### Getting Started

1. **Initialize your dotfiles repository**:

```bash
# Clone an existing dotfiles repository
dotx sync init https://github.com/username/dotfiles.git

# Or create a new repository and push it later
```

2. **Add configuration files**:

```bash
# Add your shell configuration
dotx add ~/.bashrc

# Add a directory of configuration files
dotx add ~/.config/nvim
```

3. **Deploy your dotfiles on a new system**:

```bash
# After initializing your repository
dotx deploy
```

4. **Synchronize changes**:

```bash
# Pull latest changes from remote
dotx sync pull

# Push your changes to remote
dotx sync push -m "Update nvim configuration"
```

## Commands

### `add`

Track a configuration file or directory in your dotfiles repository by creating a symlink to its original location.

```bash
dotx add <path>
```

Example:
```bash
dotx add ~/.bashrc
dotx add ~/.config/nvim
```

### `deploy`

Create symbolic links from your dotfiles repository to their appropriate locations in your home directory.

```bash
dotx deploy
```

### `sync`

Manage Git operations for your dotfiles repository.

#### `sync init`

Set up your dotfiles environment by cloning an existing Git repository containing your configuration files.

```bash
dotx sync init <repository-url>
```

Example:
```bash
dotx sync init https://github.com/username/dotfiles.git
```

#### `sync pull`

Fetch and merge the latest changes from your remote dotfiles repository to keep your local copy up-to-date.

```bash
dotx sync pull
```

#### `sync push`

Commit local changes to your dotfiles and push them to the remote repository for backup and sharing.

```bash
dotx sync push [-m <commit-message>]
```

Example:
```bash
dotx sync push -m "Update bash aliases"
```

## Configuration

dotx uses two configuration files that follow the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html):

### Application Configuration

Located at `$XDG_CONFIG_HOME/dotx/config.yaml` (typically `~/.config/dotx/config.yaml` on Unix-like systems):

```yaml
verbose: false                    # Enable verbose logging
commitMessage: "update dotfiles"  # Default commit message for sync push
```

### Repository Configuration

Located at `$XDG_DATA_HOME/dotx/dotfiles/dotx.yaml` (typically `~/.local/share/dotx/dotfiles/dotx.yaml` on Unix-like systems):

```yaml
dotfiles:
  - source: "bashrc"
    destination: "$HOME/.bashrc"
  - source: "config/nvim"
    destination: "$HOME/.config/nvim"
```

This file is automatically updated when you add new dotfiles.

## How It Works

1. When you add a file with `dotx add`, the original file is moved to the dotfiles repository.
2. A symbolic link is created at the original location, pointing to the file in the repository.
3. The mapping between the repository file and the original location is stored in the configuration.
4. When you deploy with `dotx deploy`, symbolic links are created based on this mapping.

## License

This project is licensed under the [Apache License 2.0](LICENSE).
