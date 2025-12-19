# commitmate

🤖 AI-powered Git commit message generator

**[English](README.md)** | [한국어](docs/ko.md)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/leehosu/commitmate)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/leehosu/commitmate)](https://github.com/leehosu/commitmate/releases)

>  **[Contributing](CONTRIBUTING.md)** | **[Changelog](CHANGELOG.md)**

## Features

- ✨ **AI-powered commit message generation**: Supports OpenAI GPT and Anthropic Claude
- 📝 **Conventional Commits format**: Industry-standard commit message convention
- 🌏 **Multilingual support**: Korean/English commit messages and UI
- 🎫 **JIRA integration**: Auto-add JIRA issue numbers from branch names
- 🎯 **Simple usage**: Commit with a single command
- ⚙️ **Flexible configuration**: Choose API keys and providers
- 🚀 **Cross-platform**: Linux, macOS, Windows support

## Installation

### Homebrew (Recommended) 🍺

```bash
# Add tap
brew tap leehosu/tap

# Install
brew install commitmate

# Verify
commitmate version
```

### Binary Download

Download the binary for your OS from the latest release:
[Releases](https://github.com/leehosu/commitmate/releases)

```bash
# macOS/Linux
tar -xzf commitmate_*_*.tar.gz
chmod +x commitmate
sudo mv commitmate /usr/local/bin/

# Windows
# Extract commitmate.exe and add to PATH
```

## Quick Start

### 1. Set API Key

**Using OpenAI:**
```bash
commitmate config set-key openai sk-xxxxx
commitmate config set-provider openai
```

**Using Claude:**
```bash
commitmate config set-key claude sk-ant-xxxxx
commitmate config set-provider claude
```

### 2. Generate Commit

```bash
# After making changes
git add .

# AI automatically generates commit message and commits
commitmate
```

## Usage

### Basic Commands

```bash
# Basic usage (analyze staged changes and commit)
commitmate

# Generate message only without committing
commitmate --dry-run

# Use specific AI provider (one-time)
commitmate --provider openai
commitmate --provider claude

# Skip git hooks
commitmate --no-verify
```

### Configuration Management

```bash
# Set API key
commitmate config set-key openai sk-xxxxx
commitmate config set-key claude sk-ant-xxxxx

# Set default provider
commitmate config set-provider openai

# Change model
commitmate config set-model openai gpt-4o-mini
commitmate config set-model claude claude-3-5-haiku-20241022

# Language settings
commitmate config set-commit-language ko  # Commit message language (ko/en)
commitmate config set-ui-language en      # UI language (ko/en)

# JIRA integration
commitmate config set-jira-integration true   # Enable JIRA integration
commitmate config set-jira-integration false  # Disable JIRA integration

# Show current configuration
commitmate config show

# Check version
commitmate version
```

### Environment Variables

You can also configure using environment variables:

```bash
export COMMITMATE_OPENAI_API_KEY=sk-xxxxx
export COMMITMATE_CLAUDE_API_KEY=sk-ant-xxxxx
export COMMITMATE_PROVIDER=openai
export COMMITMATE_COMMIT_LANGUAGE=ko  # Commit message language
export COMMITMATE_UI_LANGUAGE=en      # UI language
```

## Conventional Commits

commitmate follows the [Conventional Commits](https://www.conventionalcommits.org/) format:

```
<type>(<scope>): <subject>

[optional body]

[optional footer]
```

**Supported types:**
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style`: Code formatting (no functional changes)
- `refactor`: Code refactoring
- `test`: Adding/modifying tests
- `chore`: Build, config, and other changes
- `perf`: Performance improvements
- `ci`: CI configuration changes
- `build`: Build system changes
- `revert`: Revert previous commit

## Multilingual Support

commitmate supports both Korean and English:

### Commit Message Language
Configure the language for AI-generated commit messages:

```bash
# English commit messages (default, suitable for global teams)
commitmate config set-commit-language en

# Korean commit messages
commitmate config set-commit-language ko
```

### UI Language
Configure the CLI interface language:

```bash
# Korean UI (default)
commitmate config set-ui-language ko

# English UI
commitmate config set-ui-language en
```

### Usage Scenarios

**Scenario 1: Korean developer, global team**
```bash
commitmate config set-commit-language en  # English commit messages
commitmate config set-ui-language ko      # Korean UI
```

**Scenario 2: International developer, Korean company**
```bash
commitmate config set-commit-language ko  # Korean commit messages
commitmate config set-ui-language en      # English UI
```

**Scenario 3: All in English**
```bash
commitmate config set-commit-language en  # English commit messages
commitmate config set-ui-language en      # English UI
```

## JIRA Integration

commitmate **automatically** detects JIRA issue numbers from branch names and adds them to commit messages - no configuration needed!

### How it works

Simply create a branch with a JIRA issue pattern (e.g., `PROJECT-123`, `DEVOPS2-430`) and commitmate will automatically detect and add it to your commit message.

**With JIRA pattern:**
```bash
# Create branch with JIRA issue
git checkout -b DEVOPS2-430-add-user-feature

# Generate commit
git add .
commitmate

# Result: [DEVOPS2-430] feat: add user authentication
```

**Without JIRA pattern:**
```bash
# Regular branch name
git checkout -b feature/add-auth

# Generate commit  
git add .
commitmate

# Result: feat: add user authentication (no JIRA prefix)
```

### Automatic behavior

commitmate **always** works automatically:
- ✅ Detects JIRA patterns in branch names
- ✅ Adds `[ISSUE-123]` prefix if pattern is found
- ✅ Skips prefix if no JIRA pattern exists
- ✅ No configuration or environment variables needed
- ✅ Excludes special branches (main, master, develop)

### Supported patterns

- `PROJECT-123`
- `ABC-456`
- `DEVOPS2-430`
- Standard JIRA project key + number combinations

**Note:** JIRA issue numbers are not added when on main, master, or develop branches.

## Example

```bash
$ git add .
$ commitmate

🔍 Analyzing staged changes...
✨ AI generated commit message:

feat(auth): add JWT authentication middleware

? Do you want to use this commit message? 
  ▸ Yes - commit
    Edit - edit and commit
    Regenerate - generate again
    Cancel

✓ Commit completed successfully!
```


## License

MIT License - See [LICENSE](LICENSE) file

## Contributing

Issues and PRs are welcome!

## Author

[@leehosu](https://github.com/leehosu)
