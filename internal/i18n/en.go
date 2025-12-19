package i18n

func getEnglishMessages() Messages {
	return Messages{
		// Git operations
		CheckingRepository:     "🔍 Checking repository...",
		AnalyzingStagedChanges: "📝 Analyzing staged changes...",
		LargeChangesWarning:    "⚠️  Large changes detected (%d bytes). Only summarized information will be sent.",
		NoChanges:              "No changes to commit. Please run 'git add' first",

		// AI operations
		GeneratingMessage:   "✨ Generating commit message with AI...",
		RegeneratingMessage: "✨ Regenerating commit message...",
		MessageGenerated:    "✨ Generated commit message:",
		NewMessageGenerated: "✨ Newly generated commit message:",

		// Commit operations
		Committing:      "🚀 Committing...",
		CommitSuccess:   "✓ Commit completed successfully!",
		CommitCancelled: "Commit cancelled",
		DryRunMode:      "ℹ️  --dry-run mode: not committing",

		// Prompts
		PromptLabel:      "Do you want to use this commit message?",
		PromptYes:        "✓ Yes - commit",
		PromptEdit:       "✎ Edit - edit and commit",
		PromptRegenerate: "↻ Regenerate - generate again",
		PromptCancel:     "✗ Cancel",
		EditPromptLabel:  "Commit message",

		// Config
		ConfigTitle:           "📋 Current configuration:",
		ConfigProvider:       "Default provider: %s",
		ConfigCommitLanguage: "Commit message language: %s",
		ConfigUILanguage:     "UI language: %s",
		ConfigTemplate:       "Template: %s",
		OpenAISettings:       "OpenAI settings:",
		ClaudeSettings:        "Claude settings:",
		APIKeyLabel:           "  API key: %s",
		APIKeyNotSet:          "  API key: (not set)",
		ModelLabel:            "  Model: %s",
		MaxTokensLabel:        "  Max tokens: %d",

		// Success messages
		APIKeySaved:        "✓ %s API key saved successfully",
		ProviderSet:       "✓ Default provider set to %s",
		ModelSet:          "✓ %s model set to %s",
		CommitLanguageSet: "✓ Commit message language set to %s",
		UILanguageSet:     "✓ UI language set to %s",

		// Error messages
		ErrorNotGitRepo:      "Not a git repository",
		ErrorNoStagedChanges: "Failed to get staged changes",
		ErrorLoadConfig:      "Failed to load configuration",
		ErrorCreateClient:    "Failed to create AI client: %v",
		ErrorGenerateMessage: "Failed to generate commit message",
		ErrorCommitFailed:    "Commit failed",
		ErrorInvalidProvider: "Invalid provider. Use 'openai' or 'claude'",
		ErrorInvalidLanguage: "Invalid language. Use 'ko' or 'en'",
		ErrorSaveConfig:      "Failed to save configuration: %v",
		ErrorSelectFailed:     "Selection failed",
		ErrorInputFailed:      "Input failed",
		ErrorGetConfigPath:    "Failed to load configuration: %v",
		ErrorUnknownSelection: "Unknown selection: %s",
		ErrorInvalidBoolValue: "Invalid value. Use 'true' or 'false'",

		// Hint messages
		HintSetAPIKey: "Hint: Set API key with 'commitmate config set-key %s <API_KEY>'",

		// JIRA integration
		JiraIssueDetected: "🎫 JIRA issue detected: %s",
		JiraIssueAdded:    "   JIRA issue added to commit message",
	}
}
