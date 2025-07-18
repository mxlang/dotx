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
        <a href="#overview"><b>Overview</b></a> •
        <a href="#installation"><b>Installation</b></a> •
        <a href="#usage"><b>Usage</b></a> •
        <a href="#commands"><b>Commands</b></a> •
        <a href="#configuration"><b>Configuration</b></a> •
        <a href="#how-it-works"><b>How It Works</b></a> •
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

- **Track** configuration files in a Git repository
- **Deploy** your dotfiles to new systems
- **Synchronize** changes across multiple machines
- **Maintain** a clean and organized dotfiles setup

## Installation

### Using Go

```bash
go install github.com/mxlang/dotx/cmd/dotx@latest
```

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

All commands support the following flags:

- `--verbose` (`-v`): Enable more detailed output, which can be helpful for debugging.
- `--version`: Display the current version information for dotx.

### `add`

Add a file or directory to your dotfiles. This command tracks a configuration file or directory in your dotfiles by creating a symlink to its original location.

```bash
dotx add <path> [-d, --dir]
```

Options:
- `-d, --dir`: Optional directory to add the dotfile to

Example:
```bash
dotx add ~/.bashrc
dotx add ~/.config/nvim
dotx add -d starship ~/.config/starship.toml
```

### `deploy`

Deploy your dotfiles to the current system. This command creates symbolic links from your dotfiles to their appropriate locations in your home directory.

```bash
dotx deploy [-f, --force]
```

Options:
- `-f, --force`: Do not prompt for confirmation when overwriting existing files

Example:
```bash
dotx deploy
dotx deploy --force
```

### `sync`

Manage Git operations for your dotfiles repository. This command provides subcommands for initializing, pulling, and pushing changes to synchronize your dotfiles across systems.

#### `sync init`

Initialize by cloning a remote dotfiles repository. This command sets up your dotfiles environment by cloning an existing Git repository containing your configuration files and running your configured scripts. You can run `dotx sync init` without a URL if you have already cloned a remote repository. This will execute the init scripts again.

```bash
dotx sync init [repository-url] [-d, --deploy] [-f, --force]
```

Options:
- `-d, --deploy`: Automatically deploy dotfiles after initialization
- `-f, --force`: Do not prompt for confirmation when overwriting existing files

Example:
```bash
dotx sync init
dotx sync init https://github.com/username/dotfiles.git
dotx sync init https://github.com/username/dotfiles.git --deploy --force
```

#### `sync pull`

Update local dotfiles by pulling changes from the remote repository. This command fetches and merges the latest changes from your remote dotfiles repository to keep your local copy up-to-date.

```bash
dotx sync pull [-d, --deploy] [-f, --force]
```

Options:
- `-d, --deploy`: Automatically deploy dotfiles after pulling
- `-f, --force`: Do not prompt for confirmation when overwriting existing files

Example:
```bash
dotx sync pull
dotx sync pull --deploy --force
```

#### `sync push`

Save and upload local dotfile changes to the remote repository. This command commits local changes to your dotfiles and pushes them to the remote repository for backup and sharing.

```bash
dotx sync push [-m, --message <commit-message>]
```

Options:
- `-m, --message`: Specify a commit message (if not provided, uses the default commit message from config)

Example:
```bash
dotx sync push
dotx sync push -m "Update bash aliases"
```

#### `sync status`

Show if your dotfiles repository is up to date with the remote.

Options:
- `-p, --prompt`: Return a boolean that can be used in your shell prompt to indicate your dotfiles repository status

Example:
```bash
dotx sync status
dotx sync status --prompt
```

## Configuration

dotx uses two configuration files that follow the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html):

### Application Configuration

Located at `$XDG_CONFIG_HOME/dotx/config.yaml` (typically `~/.config/dotx/config.yaml` on Linux and `~/Library/Application Support/dotx/config.yaml` on macOS):

```yaml
verbose: true                            # Enable verbose logging
commitMessage: "default commit message"  # Default commit message for sync push
deployOnInit: true                       # Automatically deploy dotfiles after initialization
deployOnPull: true                       # Automatically deploy dotfiles after pulling
```

You can create or edit this file manually to customize dotx's behavior. If the file doesn't exist, dotx will use default values.

### Repository Configuration

Located at `$XDG_DATA_HOME/dotx/dotfiles/dotx.yaml` (typically `~/.local/share/dotx/dotfiles/dotx.yaml` on Linux and `~/Library/Application Support/dotx/dotfiles/dotx.yaml` on macOS):

```yaml
dotfiles:
  - source: "/.bashrc"
    destination: "$HOME/.bashrc"
  - source: "/.config/nvim"
    destination: "$HOME/.config/nvim"
scripts:
  init:
    - setup.sh
    - scripts/bootstrap.sh
```

This file is automatically updated when you add new dotfiles using the `add` command. You can manually add scripts to the init property these scripts will be executed every time you run `dotx sync init`.

## How It Works

### File Management

1. When you add a file with `dotx add`:
   - The original file is moved to the dotfiles repository
   - A symbolic link is created at the original location, pointing to the file in the repository
   - The mapping between the repository file and the original location is stored in the configuration

2. When you deploy with `dotx deploy`:
   - Symbolic links are created based on the mappings in the configuration
   - If a file already exists at the target location, you'll be prompted to overwrite it

### Repository Structure

```
~/.local/share/dotx/dotfiles/  # Default repository location
├── .bashrc                    # Your actual configuration files
├── .vimrc
├── .config/
│   └── nvim/                  # Directories are preserved
│       ├── init.vim
│       └── ...
│   setup.sh                   # Scripts that runs everytime you execute `dotx sync init`
├── scripts/
│   └── bootstrap.sh           
└── dotx.yaml                  # Repository configuration file
```

### Synchronization

- `dotx sync pull` fetches changes from your remote repository
- `dotx sync push` commits and pushes your local changes to the remote
- This allows you to keep your dotfiles in sync across multiple machines

### Scripting (Hooks)

dotx allows you to automate custom setup steps by defining scripts (also known as hooks) in your repository configuration file (`dotx.yaml`). These scripts are especially useful for tasks such as installing dependencies, setting up environments, or running any initialization logic after cloning or updating your dotfiles.

#### Init Scripts

You can specify scripts to be executed automatically every time you run `dotx sync init`. To do this, add them under the `scripts.init` property in your `dotx.yaml` file:

```yaml
scripts:
  init:
    - setup.sh
    - scripts/bootstrap.sh
```

When you run:

```bash
dotx sync init
```

dotx will execute each script listed in the `init` section, in the order they appear. This allows you to automate any setup or bootstrapping tasks required for your environment.

##### Example use cases

- Installing required packages
- Setting up programming language environments
- Configuring system preferences
- Running custom shell commands

## License

This project is licensed under the [Apache License 2.0](LICENSE).
