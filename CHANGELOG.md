# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Nothing

### Changed
- Nothing

### Fixed
- Nothing

### Removed
- Nothing

## [1.1.4] - 2024-12-19

### Changed
- 🎨 **Improved edit screen layout**: Separated label and input field
  - First line: Label display ("Commit message (Ctrl+C to go back)")
  - Second line: Input field with 🤖 emoji
  - Much cleaner and more readable layout

### Improved
- ✨ **Visual enhancement**: Clear indication of AI-generated message
- 📝 **Better readability**: Label and input separated for intuitive UX

## [1.1.3] - 2024-12-19

### Fixed
- 🐛 **Corrected keyboard shortcut hint**: ESC → Ctrl+C
  - promptui.Prompt only supports Ctrl+C (ESC not supported)
  - Updated input screen label with correct shortcut
  - Fixed for both Korean and English

## [1.1.2] - 2024-12-19

### Changed
- 🚀 **Dramatically improved edit workflow**: Much faster and more intuitive editing experience
  - Edit → Direct input screen (removed intermediate menu)
  - Display "ESC to go back" hint on input screen
  - Provide 3 actions after edit: Use / Edit again / Back
  - Immediate editing without unnecessary steps

### Improved
- ✨ **UX optimization**: Faster workflow by removing unnecessary intermediate steps
- 📝 **Clear guidance**: Navigation instructions explicitly shown in UI
- 🔄 **Flexible iteration**: Edit again option for multiple edits until satisfied

## [1.1.1] - 2024-12-19

### Added
- 🎨 **Edit screen navigation menu**: 3 options when Edit is selected
  - ✎ Edit - Edit message
  - ✓ Use - Commit with current message
  - ↩️ Back - Return to main menu
- ↩️ **Back navigation feedback during edit**: Clear cancellation message when ESC/Ctrl+C is pressed

### Changed
- 🔄 **Auto-return after edit**: Return to selection screen instead of immediate commit after editing
- 📝 **Repeatable editing**: Edit multiple times until satisfied

### Fixed
- 🐛 Fixed golangci-lint S1023 warning (removed redundant break statement)

### Improved
- ✨ Overall UX improvement: More flexible navigation and clear feedback

## [1.1.0] - 2024-12-19

### Changed (Breaking Changes)
- 🎯 **Fully automated JIRA integration**: Always works automatically without configuration
  - Removed `set-jira-integration` command
  - Removed `JiraIntegration` config field
  - Removed `COMMITMATE_JIRA_INTEGRATION` environment variable
  - If branch has JIRA pattern → Automatically adds `[ISSUE-123]` prefix
  - If branch doesn't have JIRA pattern → Skips prefix
  - No configuration needed, works based on branch name pattern

### Improved
- 🚀 **Simplified user experience**: No need to toggle settings
- ✨ **Intelligent automation**: Automatically detects and applies JIRA issue keys

## [1.0.1] - 2024-12-19

⚠️ **Deprecated**: This version is deprecated due to JIRA integration configuration complexity. Please upgrade to v1.1.0 or later.

### Added
- 🎫 **JIRA integration**: Automatically extract JIRA issue from branch name and prepend to commit message
  - Pattern detection: `DEVOPS2-430`, `PROJ-123` etc.
  - Configurable via `set-jira-integration` command
  - Configurable via `COMMITMATE_JIRA_INTEGRATION` environment variable

### Changed
- 📝 **Documentation restructuring**: Main README in English, Korean docs moved to `docs/ko.md`

### Improved
- 🎨 **Automatic JIRA prefix**: Reduces manual work and ensures consistency

## [1.0.0] - 2024-12-19

### Changed (Breaking Changes)
- 🔄 **Project renamed**: `commitgen` → `commitmate`
  - Repository name: `leehosu/commitgen` → `leehosu/commitmate`
  - Module path: `github.com/leehosu/commitgen` → `github.com/leehosu/commitmate`
  - Binary name: `commitgen` → `commitmate`
  - Homebrew package: `brew install commitgen` → `brew install commitmate`
  - Config directory: `~/.commitgen` → `~/.commitmate`
  - Environment variables: `COMMITGEN_*` → `COMMITMATE_*`

### Migration Guide
```bash
# Uninstall old version
brew uninstall commitgen

# Install new version
brew tap leehosu/tap
brew install commitmate

# Migrate config (optional)
mv ~/.commitgen ~/.commitmate
```

### Fixed
- 🐛 Fixed i18n support for error messages
  - All error messages now respect UI language setting
  - Fixed Ctrl+C error message not respecting language

## [0.3.0] - 2024-12-16 (commitgen)

### Added
- 🌍 **Multilingual support**: Separate language settings for commit messages and UI
  - `CommitLanguage`: Language for AI-generated commit messages (Korean/English)
  - `UILanguage`: Language for CLI UI messages (Korean/English)
  - Set via `set-commit-language` and `set-ui-language` commands
  - Set via `COMMITGEN_COMMIT_LANGUAGE` and `COMMITGEN_UI_LANGUAGE` environment variables
- 📚 **Bilingual documentation**: English README with Korean docs in `docs/ko.md`

### Changed
- 🔄 Replaced single `Language` setting with `CommitLanguage` and `UILanguage`

### Improved
- ✨ **Flexible language control**: Independent control of commit message and UI language
- 🌐 **Better internationalization**: Complete i18n system implementation

## [0.2.1] - 2024-12-15 (commitgen)

### Fixed
- 🔧 Fixed GoReleaser GitHub token permission issues
- 📦 Fixed Homebrew formula publishing

## [0.2.0] - 2024-12-15 (commitgen)

### Added
- 🍺 **Homebrew tap integration**: Automatic formula updates on release
  - `brew tap leehosu/tap`
  - `brew install commitgen`
- 🤖 **Automated releases**: GitHub Actions with GoReleaser
  - Cross-platform binary builds
  - Automatic GitHub release creation
  - Automatic Homebrew formula updates

### Improved
- 📦 **Simplified installation**: One-line Homebrew install
- 🔄 **Automated distribution**: No manual release process

## [0.1.1] - 2024-12-14 (commitgen)

### Fixed
- 🐛 Fixed error handling in commit command

## [0.1.0] - 2024-12-13 (commitgen)

### Added
- 🎉 **Initial release**
- 🤖 **AI-powered commit message generation**
  - OpenAI GPT support
  - Anthropic Claude support
- 📝 **Conventional Commits format**
- 🎨 **Interactive UI with prompt selection**
- ⚙️ **Configuration management**
  - API key setup
  - Provider selection
  - Language settings
  - Custom templates
- 🔄 **Commit workflow**
  - Generate message
  - Edit message
  - Regenerate message
  - Cancel operation
- 🌍 **Multilingual support** (Korean/English)
