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

#### Setup your shell to use dotx

```bash
eval "$(dotx init <shell>)"
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

Add one or more files or directories to your dotfiles. This command tracks configuration files or directories in your dotfiles by creating symlinks to their original locations.

```bash
dotx add <path>... [-d, --dir]
```

Options:
- `-d, --dir`: Optional directory to add the dotfile to

Example:
```bash
dotx add ~/.bashrc
dotx add ~/.config/nvim
dotx add ~/.bashrc ~/.zshrc
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

### `cd`

Go to your local dotfiles directory. You need to add `eval "$(dotx init <shell>)"` to your shell config to work properly. It has to be one of the following options: bash, zsh, fish.

Example:
```bash
dotx cd
```

### `doctor`

Inspect your system and repository for undeployed dotfiles and help you link them. This command scans your dotx repository for dotfiles that are not currently deployed (i.e., not symlinked in their destination paths) and presents a TUI to choose which ones to deploy.

```bash
dotx doctor [-f, --force]
```

Options:
- `-f, --force`: Never prompt for overwriting existing files when deploying selected dotfiles

Example:
```bash
dotx doctor
dotx doctor --force # Deploy without any overwrite prompts
```

### `sync`

Manage Git operations for your dotfiles repository. This command provides subcommands for initializing, pulling, and pushing changes to synchronize your dotfiles across systems.

#### `sync init`

Initialize by cloning a remote dotfiles repository. This command sets up your dotfiles environment by cloning an existing Git repository containing your configuration files.

```bash
dotx sync init <repository-url> [-d, --deploy] [-f, --force]
```

Options:
- `-d, --deploy`: Automatically deploy dotfiles after initialization
- `-f, --force`: Do not prompt for confirmation when overwriting existing files

Example:
```bash
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
- `-m, --message`: Specify a commit message (if not provided, you'll be prompted to enter one)

Example:
```bash
dotx sync push
dotx sync push -m "Update bash aliases"
```


## Configuration

dotx uses two configuration files that follow the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html):

### Application Configuration

Located at `$XDG_CONFIG_HOME/dotx/config.yaml` (typically `~/.config/dotx/config.yaml` on Linux and `~/Library/Application Support/dotx/config.yaml` on macOS):

```yaml
verbose: true                            # Enable verbose logging
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
  - path: "setup.sh"
    on: init           # one of: init, pull, deploy
    run: once          # optional: always (default) | once | changed
  - path: "scripts/bootstrap.sh"
    on: init
    run: changed
```

This file is automatically updated when you add new dotfiles using the `add` command. You can define scripts as list entries with a path and an on event (one of: init, pull, deploy). Optionally set run to control execution frequency: always (default), once, or changed (runs when file content hash changes). Scripts fire on the corresponding commands: init for `dotx sync init`, pull for `dotx sync pull`, and deploy for `dotx deploy`. 

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
│   setup.sh                   # Example script referenced by scripts entries
├── scripts/
│   └── bootstrap.sh           
└── dotx.yaml                  # Repository configuration file
```

### Synchronization

- `dotx sync pull` fetches changes from your remote repository
- `dotx sync push` commits and pushes your local changes to the remote
- This allows you to keep your dotfiles in sync across multiple machines

### Scripting (Hooks)

dotx allows you to automate custom setup steps by defining scripts (also known as hooks) in your repository configuration file (`dotx.yaml`). These scripts are especially useful for tasks such as installing dependencies, setting up environments, or running any initialization logic.

#### Script events and run conditions

Define scripts under `scripts` as a list of entries with a path, an `on` event, and an optional `run` condition:

```yaml
scripts:
  - path: "scripts/bootstrap.sh"
    on: init
    run: once
  - path: "scripts/post-pull.sh"
    on: pull
  - path: "scripts/post-deploy.sh"
    on: deploy
```

- `on` controls when the script runs. Allowed values: `init`, `pull`, `deploy`.
- `run` controls how often the script runs:
  - `always` (default): run every time the event occurs
  - `once`: run only once ever
  - `changed`: run only if the file's content hash has changed since the last run

Execution:
- `dotx sync init` triggers scripts with `on: init` (after cloning and setup)
- `dotx sync pull` triggers scripts with `on: pull` (after pulling updates)
- `dotx deploy` triggers scripts with `on: deploy` (after deployment)

##### Example use cases

- Installing required packages
- Setting up programming language environments
- Configuring system preferences
- Running custom shell commands

## License

This project is licensed under the [Apache License 2.0](LICENSE).
