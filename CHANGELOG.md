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

## [0.2.0] - 2024-12-19

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

## [0.1.1] - 2024-12-18

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

## [0.1.0] - 2025-12-18

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
