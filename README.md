# commitgen

🤖 AI-powered Git commit message generator

**[English](README.md)** | [한국어](docs/ko.md)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/leehosu/commitgen)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/leehosu/commitgen)](https://github.com/leehosu/commitgen/releases)

>  **[Contributing](CONTRIBUTING.md)** | **[Changelog](CHANGELOG.md)**

## Features

- ✨ **AI-powered commit message generation**: Supports OpenAI GPT and Anthropic Claude
- 📝 **Conventional Commits format**: Industry-standard commit message convention
- 🌏 **Multilingual support**: Korean/English commit messages and UI
- 🎯 **Simple usage**: Commit with a single command
- ⚙️ **Flexible configuration**: Choose API keys and providers
- 🚀 **Cross-platform**: Linux, macOS, Windows support

## Installation

### Homebrew (Recommended) 🍺

```bash
# Add tap
brew tap leehosu/tap

# Install
brew install commitgen

# Verify
commitgen version
```

### Binary Download

Download the binary for your OS from the latest release:
[Releases](https://github.com/leehosu/commitgen/releases)

```bash
# macOS/Linux
tar -xzf commitgen_*_*.tar.gz
chmod +x commitgen
sudo mv commitgen /usr/local/bin/

# Windows
# Extract commitgen.exe and add to PATH
```

## Quick Start

### 1. Set API Key

**Using OpenAI:**
```bash
commitgen config set-key openai sk-xxxxx
commitgen config set-provider openai
```

**Using Claude:**
```bash
commitgen config set-key claude sk-ant-xxxxx
commitgen config set-provider claude
```

### 2. Generate Commit

```bash
# After making changes
git add .

# AI automatically generates commit message and commits
commitgen
```

## Usage

### Basic Commands

```bash
# Basic usage (analyze staged changes and commit)
commitgen

# Generate message only without committing
commitgen --dry-run

# Use specific AI provider (one-time)
commitgen --provider openai
commitgen --provider claude

# Skip git hooks
commitgen --no-verify
```

### Configuration Management

```bash
# Set API key
commitgen config set-key openai sk-xxxxx
commitgen config set-key claude sk-ant-xxxxx

# Set default provider
commitgen config set-provider openai

# Change model
commitgen config set-model openai gpt-4o-mini
commitgen config set-model claude claude-3-5-haiku-20241022

# Language settings
commitgen config set-commit-language ko  # Commit message language (ko/en)
commitgen config set-ui-language en      # UI language (ko/en)

# Show current configuration
commitgen config show

# Check version
commitgen version
```

### Environment Variables

You can also configure using environment variables:

```bash
export COMMITGEN_OPENAI_API_KEY=sk-xxxxx
export COMMITGEN_CLAUDE_API_KEY=sk-ant-xxxxx
export COMMITGEN_PROVIDER=openai
export COMMITGEN_COMMIT_LANGUAGE=ko  # Commit message language
export COMMITGEN_UI_LANGUAGE=en      # UI language
```

## Conventional Commits

commitgen follows the [Conventional Commits](https://www.conventionalcommits.org/) format:

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

commitgen supports both Korean and English:

### Commit Message Language
Configure the language for AI-generated commit messages:

```bash
# English commit messages (default, suitable for global teams)
commitgen config set-commit-language en

# Korean commit messages
commitgen config set-commit-language ko
```

### UI Language
Configure the CLI interface language:

```bash
# Korean UI (default)
commitgen config set-ui-language ko

# English UI
commitgen config set-ui-language en
```

### Usage Scenarios

**Scenario 1: Korean developer, global team**
```bash
commitgen config set-commit-language en  # English commit messages
commitgen config set-ui-language ko      # Korean UI
```

**Scenario 2: International developer, Korean company**
```bash
commitgen config set-commit-language ko  # Korean commit messages
commitgen config set-ui-language en      # English UI
```

**Scenario 3: All in English**
```bash
commitgen config set-commit-language en  # English commit messages
commitgen config set-ui-language en      # English UI
```

## Example

```bash
$ git add .
$ commitgen

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
