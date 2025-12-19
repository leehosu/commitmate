# commitmate

🤖 AI 기반 Git 커밋 메시지 자동 생성 도구

[English](../README.md) | **[한국어](ko.md)**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/leehosu/commitmate)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/leehosu/commitmate)](https://github.com/leehosu/commitmate/releases)

>  **[기여하기](../CONTRIBUTING.md)** | **[변경 이력](../CHANGELOG.md)**

## 특징

- ✨ **AI 기반 커밋 메시지 생성**: OpenAI GPT 및 Anthropic Claude 지원
- 📝 **Conventional Commits 포맷**: 업계 표준 커밋 메시지 형식
- 🌏 **다국어 지원**: 한글/영어 커밋 메시지 및 UI 지원
- 🎯 **간단한 사용법**: 한 번의 명령어로 커밋 완료
- ⚙️ **유연한 설정**: API 키 및 제공자 선택 가능
- 🚀 **크로스 플랫폼**: Linux, macOS, Windows 지원

## 설치

### Homebrew (추천) 🍺

```bash
# Tap 추가
brew tap leehosu/tap

# 설치
brew install commitmate

# 확인
commitmate version
```

### 바이너리 다운로드

최신 릴리즈에서 OS에 맞는 바이너리를 다운로드하세요:
[Releases](https://github.com/leehosu/commitmate/releases)

```bash
# macOS/Linux
tar -xzf commitmate_*_*.tar.gz
chmod +x commitmate
sudo mv commitmate /usr/local/bin/

# Windows
# commitmate.exe를 압축 해제 후 PATH에 추가
```

## 빠른 시작

### 1. API 키 설정

**OpenAI 사용:**
```bash
commitmate config set-key openai sk-xxxxx
commitmate config set-provider openai
```

**Claude 사용:**
```bash
commitmate config set-key claude sk-ant-xxxxx
commitmate config set-provider claude
```

### 2. 커밋 생성

```bash
# 파일 변경 후
git add .

# AI가 자동으로 커밋 메시지 생성 및 커밋
commitmate
```

## 사용법

### 기본 명령어

```bash
# 기본 사용 (staged 변경사항 분석 및 자동 커밋)
commitmate

# 커밋 메시지만 생성하고 커밋하지 않음
commitmate --dry-run

# 특정 AI 제공자 사용 (일회성)
commitmate --provider openai
commitmate --provider claude

# git hooks 무시
commitmate --no-verify
```

### 설정 관리

```bash
# API 키 설정
commitmate config set-key openai sk-xxxxx
commitmate config set-key claude sk-ant-xxxxx

# 기본 제공자 설정
commitmate config set-provider openai

# 모델 변경
commitmate config set-model openai gpt-4o-mini
commitmate config set-model claude claude-3-5-haiku-20241022

# 언어 설정
commitmate config set-commit-language ko  # 커밋 메시지 언어 (ko/en)
commitmate config set-ui-language en      # UI 언어 (ko/en)

# 현재 설정 확인
commitmate config show

# 버전 확인
commitmate version
```

### 환경변수

설정 파일 대신 환경변수로도 설정 가능:

```bash
export COMMITMATE_OPENAI_API_KEY=sk-xxxxx
export COMMITMATE_CLAUDE_API_KEY=sk-ant-xxxxx
export COMMITMATE_PROVIDER=openai
export COMMITMATE_COMMIT_LANGUAGE=ko  # 커밋 메시지 언어
export COMMITMATE_UI_LANGUAGE=en      # UI 언어
```

## Conventional Commits

commitmate는 [Conventional Commits](https://www.conventionalcommits.org/) 형식을 따릅니다:

```
<type>(<scope>): <subject>

[optional body]

[optional footer]
```

**지원하는 타입:**
- `feat`: 새로운 기능
- `fix`: 버그 수정
- `docs`: 문서 변경
- `style`: 코드 포맷팅 (기능 변경 없음)
- `refactor`: 리팩토링
- `test`: 테스트 추가/수정
- `chore`: 빌드, 설정 등 기타 변경
- `perf`: 성능 개선
- `ci`: CI 설정 변경
- `build`: 빌드 시스템 변경
- `revert`: 이전 커밋 되돌리기

## 다국어 지원

commitmate는 한글과 영어를 지원합니다:

### 커밋 메시지 언어
AI가 생성하는 커밋 메시지의 언어를 설정할 수 있습니다:

```bash
# 영어 커밋 메시지 (기본값, 글로벌 협업에 적합)
commitmate config set-commit-language en

# 한글 커밋 메시지
commitmate config set-commit-language ko
```

### UI 언어
CLI 인터페이스의 언어를 설정할 수 있습니다:

```bash
# 한글 UI (기본값)
commitmate config set-ui-language ko

# 영어 UI
commitmate config set-ui-language en
```

### 사용 시나리오

**시나리오 1: 한국인 개발자, 글로벌 팀**
```bash
commitmate config set-commit-language en  # 영어 커밋 메시지
commitmate config set-ui-language ko      # 한글 UI
```

**시나리오 2: 외국인 개발자, 한국 회사**
```bash
commitmate config set-commit-language ko  # 한글 커밋 메시지
commitmate config set-ui-language en      # 영어 UI
```

**시나리오 3: 모두 영어**
```bash
commitmate config set-commit-language en  # 영어 커밋 메시지
commitmate config set-ui-language en      # 영어 UI
```

## JIRA 통합

commitmate는 브랜치 이름에서 JIRA 이슈 번호를 **자동으로** 감지하여 커밋 메시지에 추가합니다 - 별도 설정 불필요!

### 작동 방식

JIRA 이슈 패턴(예: `PROJECT-123`, `DEVOPS2-430`)이 포함된 브랜치를 생성하기만 하면 commitmate가 자동으로 감지하여 커밋 메시지에 추가합니다.

**JIRA 패턴이 있는 경우:**
```bash
# JIRA 이슈가 포함된 브랜치 생성
git checkout -b DEVOPS2-430-add-user-feature

# 커밋 생성
git add .
commitmate

# 결과: [DEVOPS2-430] feat: add user authentication
```

**JIRA 패턴이 없는 경우:**
```bash
# 일반 브랜치 이름
git checkout -b feature/add-auth

# 커밋 생성
git add .
commitmate

# 결과: feat: add user authentication (JIRA 접두사 없음)
```

### 자동 동작

commitmate는 **항상** 자동으로 작동합니다:
- ✅ 브랜치 이름에서 JIRA 패턴 자동 감지
- ✅ 패턴이 있으면 `[ISSUE-123]` 접두사 추가
- ✅ 패턴이 없으면 접두사 생략
- ✅ 설정이나 환경변수 불필요
- ✅ 특수 브랜치(main, master, develop) 제외

### 지원하는 패턴

- `PROJECT-123`
- `ABC-456`
- `DEVOPS2-430`
- 일반적인 JIRA 프로젝트 키 + 번호 조합

**참고:** main, master, develop 브랜치에서는 JIRA 이슈 번호가 추가되지 않습니다.

## 예시

```bash
$ git add .
$ commitmate

🔍 Staged 변경사항 분석 중...
✨ AI가 커밋 메시지를 생성했습니다:

feat(auth): add JWT authentication middleware

? 이 커밋 메시지를 사용하시겠습니까? 
  ▸ Yes - 커밋 실행
    Edit - 수정 후 커밋
    Regenerate - 다시 생성
    Cancel - 취소

✓ 커밋이 완료되었습니다!
```


## 라이선스

MIT License - [LICENSE](../LICENSE) 파일 참조

## 기여

이슈와 PR을 환영합니다!

## 작성자

[@leehosu](https://github.com/leehosu)
