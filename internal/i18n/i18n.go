package i18n

// Messages는 UI에 표시되는 모든 메시지를 담는 구조체입니다
type Messages struct {
	// Git 작업
	CheckingRepository     string
	AnalyzingStagedChanges string
	LargeChangesWarning    string
	NoChanges              string

	// AI 작업
	GeneratingMessage   string
	RegeneratingMessage string
	MessageGenerated    string
	NewMessageGenerated string

	// 커밋 작업
	Committing    string
	CommitSuccess string
	CommitCancelled string
	DryRunMode    string

	// 프롬프트
	PromptLabel      string
	PromptYes        string
	PromptEdit       string
	PromptRegenerate string
	PromptCancel     string
	EditPromptLabel  string

	// Config
	ConfigTitle          string
	ConfigProvider       string
	ConfigCommitLanguage string
	ConfigUILanguage     string
	ConfigTemplate       string
	OpenAISettings       string
	ClaudeSettings       string
	APIKeyLabel          string
	APIKeyNotSet         string
	ModelLabel           string
	MaxTokensLabel       string

	// 성공 메시지
	APIKeySaved       string
	ProviderSet       string
	ModelSet          string
	CommitLanguageSet string
	UILanguageSet     string

	// 에러 메시지
	ErrorNotGitRepo         string
	ErrorNoStagedChanges    string
	ErrorLoadConfig         string
	ErrorCreateClient       string
	ErrorGenerateMessage    string
	ErrorCommitFailed       string
	ErrorInvalidProvider    string
	ErrorInvalidLanguage    string
	ErrorSaveConfig         string
	ErrorSelectFailed       string
	ErrorInputFailed        string
	ErrorGetConfigPath      string
	ErrorUnknownSelection   string
	ErrorInvalidBoolValue   string

	// Hint 메시지
	HintSetAPIKey string

	// JIRA 통합
	JiraIssueDetected string
	JiraIssueAdded    string
}

// GetMessages는 지정된 언어에 맞는 메시지를 반환합니다
func GetMessages(lang string) Messages {
	if lang == "ko" {
		return getKoreanMessages()
	}
	return getEnglishMessages()
}
