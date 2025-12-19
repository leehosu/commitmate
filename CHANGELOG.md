# Changelog

이 프로젝트의 모든 주목할만한 변경사항이 이 파일에 문서화됩니다.

형식은 [Keep a Changelog](https://keepachangelog.com/ko/1.0.0/)를 기반으로 하며,
이 프로젝트는 [Semantic Versioning](https://semver.org/lang/ko/)을 따릅니다.

## [Unreleased]

### 추가
- 없음

### 변경
- 없음

### 수정
- 없음

### 제거
- 없음

## [1.1.4] - 2024-12-18

### 변경
- 🎨 **Edit 화면 레이아웃 개선**: 라벨과 입력 필드 분리
  - 첫 줄: 라벨 표시 ("커밋 메시지 (Ctrl+C로 뒤로가기)")
  - 둘째 줄: 🤖 이모지와 함께 입력 필드 표시
  - 훨씬 깔끔하고 읽기 쉬운 레이아웃

### 개선
- ✨ **시각적 개선**: AI 생성 메시지임을 명확히 표시
- 📝 **가독성 향상**: 라벨과 입력이 분리되어 더 직관적

## [1.1.3] - 2024-12-18

### 수정
- 🐛 **키보드 단축키 힌트 수정**: ESC → Ctrl+C
  - promptui.Prompt는 Ctrl+C만 지원 (ESC 미지원)
  - 입력 화면 레이블을 정확한 단축키로 업데이트
  - 한글/영어 모두 수정

## [1.1.2] - 2024-12-18

### 변경
- 🚀 **Edit 플로우 대폭 개선**: 훨씬 빠르고 직관적인 수정 경험
  - Edit 선택 → 바로 입력 화면으로 이동 (중간 메뉴 제거)
  - 입력 화면에 "ESC로 뒤로가기" 힌트 표시
  - 수정 후 3가지 액션 제공: Use / Edit again / Back
  - 더 이상 불필요한 단계 없이 즉시 수정 가능

### 개선
- ✨ **UX 최적화**: 불필요한 중간 단계 제거로 더 빠른 워크플로우
- 📝 **명확한 가이드**: 입력 중 뒤로가기 방법을 UI에 명시
- 🔄 **유연한 반복**: Edit again으로 만족할 때까지 여러 번 수정 가능

## [1.1.1] - 2024-12-18

### 추가
- 🎨 **수정 화면 네비게이션 메뉴**: Edit 선택 시 3가지 옵션 제공
  - ✎ Edit - 메시지 수정
  - ✓ Use - 현재 메시지로 커밋
  - ↩️ Back - 메인 메뉴로 돌아가기
- ↩️ **수정 중 뒤로가기 피드백**: ESC/Ctrl+C 입력 시 명확한 취소 메시지 표시

### 변경
- 🔄 **수정 후 자동 복귀**: 메시지 수정 후 즉시 커밋하지 않고 선택 화면으로 복귀
- 📝 **반복 수정 가능**: 만족할 때까지 여러 번 수정 가능

### 수정
- 🐛 golangci-lint S1023 경고 수정 (불필요한 break 문 제거)

### 개선
- ✨ 전반적인 UX 향상: 더 자유로운 네비게이션과 명확한 피드백

## [1.1.0] - 2024-12-19

### 변경 (Breaking Changes)
- 🎯 **JIRA 통합 완전 자동화**: 설정 없이 항상 자동 동작
  - `set-jira-integration` 명령어 제거
  - `JiraIntegration` 설정 필드 제거
  - `COMMITMATE_JIRA_INTEGRATION` 환경변수 제거
  - 브랜치에 JIRA 패턴 있음 → 자동으로 `[ISSUE-123]` 추가
  - 브랜치에 JIRA 패턴 없음 → 접두사 생략
  - 별도 설정 불필요, 그냥 작동

### 개선
- 🌏 **완전한 i18n 구현**: 모든 하드코딩된 한글 메시지를 i18n으로 전환
  - `ErrorLoadConfig`, `ErrorUnknownSelection`, `ErrorInvalidBoolValue` 추가
  - Ctrl+C 취소 메시지가 이제 `COMMITMATE_UI_LANGUAGE` 설정을 존중
  - 모든 에러 메시지가 UI 언어 설정에 따라 표시됨

## [1.0.1] - 2024-12-19 (deprecated)

### 개선
- 🎯 **JIRA 통합 완전 자동화**: 브랜치 이름에서 JIRA 패턴을 자동 감지하여 추가
  - JIRA 패턴이 있는 브랜치: 자동으로 `[ISSUE-123]` 접두사 추가
  - JIRA 패턴이 없는 브랜치: 접두사 없이 일반 커밋 메시지
  - 별도 환경변수 설정 불필요
  - `set-jira-integration false`로 감지 비활성화 가능

### 수정
- 🌏 **완전한 i18n 구현**: 모든 하드코딩된 한글 메시지를 i18n으로 전환
  - `ErrorLoadConfig`, `ErrorUnknownSelection`, `ErrorInvalidBoolValue` 추가
  - Ctrl+C 취소 메시지가 이제 `COMMITMATE_UI_LANGUAGE` 설정을 존중
  - 모든 에러 메시지가 UI 언어 설정에 따라 표시됨

## [1.0.0] - 2024-12-19

### 변경 (Breaking Changes)
- 🔄 **프로젝트 이름 변경: commitgen → commitmate**
- 모듈 경로: `github.com/leehosu/commitgen` → `github.com/leehosu/commitmate`
- 환경변수 접두사: `COMMITGEN_` → `COMMITMATE_`
- 설정 폴더: `~/.commitgen` → `~/.commitmate`
- Homebrew 패키지: `brew install commitgen` → `brew install commitmate`
- 이진 파일 이름: `commitgen` → `commitmate`

### 추가
- 🎫 JIRA 통합: 브랜치 이름에서 JIRA 이슈 번호 자동 감지 및 추가
- `set-jira-integration` 명령어: JIRA 통합 활성화/비활성화
- 환경변수 지원: `COMMITMATE_JIRA_INTEGRATION`
- JIRA 이슈 자동 감지 알림 메시지

### 참고
- main, master, develop 브랜치는 JIRA 통합 제외
- 이전 버전(commitgen)에서 마이그레이션 필요

## [0.3.0] - 2024-12-19 (commitgen)

### 추가
- 🎫 JIRA 통합: 브랜치 이름에서 JIRA 이슈 번호 자동 감지 및 추가
- `set-jira-integration` 명령어: JIRA 통합 활성화/비활성화
- 환경변수 지원: `COMMITGEN_JIRA_INTEGRATION`
- JIRA 이슈 자동 감지 알림 메시지

### 변경
- 브랜치 이름에서 JIRA 패턴 감지 시 자동으로 커밋 메시지에 추가
- main, master, develop 브랜치는 JIRA 통합 제외

## [0.2.1] - 2024-12-19 (commitgen)

### 수정
- fmt.Errorf의 비상수 포맷 문자열 오류 수정
- 에러 처리 코드 개선

## [0.2.0] - 2024-12-19 (commitgen)

### 추가
- 🌏 다국어 지원: 커밋 메시지와 UI 언어를 독립적으로 설정 가능
- `set-commit-language` 명령어: AI가 생성하는 커밋 메시지 언어 설정 (ko/en)
- `set-ui-language` 명령어: CLI 인터페이스 언어 설정 (ko/en)
- i18n 패키지: 한글/영어 메시지 체계적 관리
- 환경변수 지원: `COMMITGEN_COMMIT_LANGUAGE`, `COMMITGEN_UI_LANGUAGE`
- 영어 README 추가 (기본)
- `docs/` 폴더: 한글 문서를 docs/ko.md로 이동

### 변경
- GoReleaser 자동화: Homebrew tap 자동 업데이트 활성화
- Config 구조: `Language` → `CommitLanguage` + `UILanguage`로 분리
- 기본 언어 설정: 커밋 메시지(en), UI(ko)
- README 구조: 영어 기본, 한글은 docs/ko.md로 분리

### 개선
- 글로벌 사용자를 위한 완전한 영어 지원
- 한국 개발자를 위한 유연한 언어 조합 설정
- 4가지 언어 시나리오 지원 (en+en, en+ko, ko+en, ko+ko)

## [0.1.1] - 2024-12-18 (commitgen)

### 추가
- Rate limit 에러 감지 및 친절한 에러 메시지
- 큰 diff 자동 요약 기능 (10KB 이상)
- Diff 크기 경고 메시지
- README에 트러블슈팅 가이드 추가

### 변경
- 기본 max_tokens를 150에서 100으로 감소 (rate limit 방지)

### 수정
- OpenAI 및 Claude rate limit 처리 개선
- 큰 변경사항으로 인한 토큰 초과 문제 해결

## [0.1.0] - 2024-12-18 (commitgen)

### 추가
- 초기 릴리즈
- OpenAI GPT 지원
- Anthropic Claude 지원
- Conventional Commits 형식 지원
- 인터랙티브 UI (승인/수정/재생성/취소)
- 설정 관리 시스템
- 다중 플랫폼 지원 (Linux, macOS, Windows)
- GoReleaser 기반 자동 릴리즈
- GitHub Actions CI/CD
- Homebrew tap 지원
