package i18n

func getKoreanMessages() Messages {
	return Messages{
		// Git 작업
		CheckingRepository:     "🔍 Git 저장소 확인 중...",
		AnalyzingStagedChanges: "📝 Staged 변경사항 분석 중...",
		LargeChangesWarning:    "⚠️  변경사항이 큽니다 (%d 바이트). 요약된 정보만 전달됩니다.",
		NoChanges:              "커밋할 변경사항이 없습니다. 'git add' 명령어를 먼저 실행하세요",

		// AI 작업
		GeneratingMessage:   "✨ AI가 커밋 메시지를 생성하는 중...",
		RegeneratingMessage: "✨ 커밋 메시지를 다시 생성하는 중...",
		MessageGenerated:    "✨ 생성된 커밋 메시지:",
		NewMessageGenerated: "✨ 새로 생성된 커밋 메시지:",
		EditedMessage:       "✏️  수정된 커밋 메시지:",

		// 커밋 작업
		Committing:      "🚀 커밋 실행 중...",
		CommitSuccess:   "✓ 커밋이 완료되었습니다!",
		CommitCancelled: "커밋이 취소되었습니다",
		EditCancelled:   "↩️  수정이 취소되었습니다. 이전 메뉴로 돌아갑니다.",
		DryRunMode:      "ℹ️  --dry-run 모드: 커밋하지 않습니다",

		// 프롬프트
		PromptLabel:           "이 커밋 메시지를 사용하시겠습니까?",
		PromptYes:             "✓ Yes - 커밋 실행",
		PromptEdit:            "✎ Edit - 수정",
		PromptRegenerate:      "↻ Regenerate - 다시 생성",
		PromptCancel:          "✗ Cancel - 취소",
		EditPromptLabel:       "🤖 커밋 메시지 (Ctrl+C로 뒤로가기)",
		EditActionLabel:       "수정된 메시지를 어떻게 하시겠습니까?",
		EditActionUseMessage:  "✓ Use - 이 메시지로 커밋",
		EditActionEditAgain:   "✎ Edit again - 다시 수정",
		EditActionBack:        "↩️  Back - 뒤로 가기",

		// Config
		ConfigTitle:          "📋 현재 설정:",
		ConfigProvider:       "기본 제공자: %s",
		ConfigCommitLanguage: "커밋 메시지 언어: %s",
		ConfigUILanguage:     "UI 언어: %s",
		ConfigTemplate:       "템플릿: %s",
		OpenAISettings:       "OpenAI 설정:",
		ClaudeSettings:       "Claude 설정:",
		APIKeyLabel:          "  API 키: %s",
		APIKeyNotSet:         "  API 키: (설정되지 않음)",
		ModelLabel:           "  모델: %s",
		MaxTokensLabel:       "  Max Tokens: %d",

		// 성공 메시지
		APIKeySaved:        "✓ %s API 키가 저장되었습니다",
		ProviderSet:       "✓ 기본 제공자가 %s로 설정되었습니다",
		ModelSet:          "✓ %s 모델이 %s로 설정되었습니다",
		CommitLanguageSet: "✓ 커밋 메시지 언어가 %s로 설정되었습니다",
		UILanguageSet:     "✓ UI 언어가 %s로 설정되었습니다",

		// 에러 메시지
		ErrorNotGitRepo:      "Git 저장소가 아닙니다",
		ErrorNoStagedChanges: "Staged 변경사항을 가져올 수 없습니다",
		ErrorLoadConfig:      "설정을 로드할 수 없습니다",
		ErrorCreateClient:    "AI 클라이언트를 생성할 수 없습니다: %v",
		ErrorGenerateMessage: "커밋 메시지를 생성할 수 없습니다",
		ErrorCommitFailed:    "커밋 실패",
		ErrorInvalidProvider: "잘못된 제공자입니다. 'openai' 또는 'claude'를 사용하세요",
		ErrorInvalidLanguage: "잘못된 언어입니다. 'ko' 또는 'en'을 사용하세요",
		ErrorSaveConfig:      "설정 저장 실패: %v",
		ErrorSelectFailed:    "선택 중 오류 발생",
		ErrorInputFailed:     "입력 중 오류 발생",
		ErrorGetConfigPath:   "설정을 불러올 수 없습니다: %v",
		ErrorUnknownSelection: "알 수 없는 선택: %s",
		ErrorInvalidBoolValue: "잘못된 값입니다. 'true' 또는 'false'를 사용하세요",

		// Hint 메시지
		HintSetAPIKey: "힌트: 'commitmate config set-key %s <API_KEY>' 명령어로 API 키를 설정하세요",

		// JIRA 통합
		JiraIssueDetected: "🎫 JIRA 이슈 감지: %s",
		JiraIssueAdded:    "   커밋 메시지에 JIRA 이슈 추가됨",
	}
}
