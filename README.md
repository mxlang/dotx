<div align="center">
    <p>
        <img width="180" alt="dotx logo" src="./docs/assets/logo.svg">
    </p>
    <h1>✨ dotx ✨</h1>
    <p>
        <b>A modern dotfile manager for tracking and syncing configuration files</b>
    </p>
    <p>
        <a href="https://github.com/mxlang/dotx/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-Apache%202.0-blue.svg" alt="License"></a>
        <img src="https://img.shields.io/badge/status-under%20development-yellow.svg" alt="Status">
        <img src="https://img.shields.io/badge/platform-Linux%20%7C%20macOS-lightgrey.svg" alt="Platform">
    </p>
    <br>
    <p>
        <a href="#installation"><b>Installation</b></a> •
        <a href="#usage"><b>Usage</b></a> •
        <a href="#commands"><b>Commands</b></a> •
        <a href="#configuration"><b>Configuration</b></a> •
        <a href="#license"><b>License</b></a>
    </p>
    <hr>
    <p>
        <h3>⚠️ This project is still under development ⚠️</h3>
    </p>
    <br>
</div>

## Overview

**dotx** helps you manage, version control, and synchronize your configuration files (dotfiles) across multiple systems. It provides a simple and intuitive CLI to:

- Track configuration files in a Git repository
- Deploy your dotfiles to new systems
- Synchronize changes across multiple machines
- Maintain a clean and organized dotfiles setup

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/mxlang/dotx.git

# Navigate to the project directory
cd dotx

# Build and install
go install ./cmd/dotx
```

### Go

```bash
go install github.com/mxlang/dotx.git
```

## Usage

### Getting Started

1. **Initialize your dotfiles repository**:

```bash
# Clone an existing dotfiles repository
dotx sync init https://github.com/username/dotfiles.git
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

Located at `$XDG_CONFIG_HOME/dotx/config.yaml` (typically `~/.config/dotx/config.yaml` on Linux and `~/Library/Application Support/dotx/config.yaml` on macOS):

```yaml
verbose: true                            # Enable verbose logging
commitMessage: "default commit message"  # Default commit message for sync push
```

### Repository Configuration

Located at `$XDG_DATA_HOME/dotx/dotfiles/dotx.yaml` (typically `~/.local/share/dotx/dotfiles/dotx.yaml` on Linux and `~/Library/Application Support/dotx/dotfiles/dotx.yaml` on macOS):

```yaml
dotfiles:
  - source: "/.bashrc"
    destination: "$HOME/.bashrc"
  - source: "/.config/nvim"
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
