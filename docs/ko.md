<div align="center">

# 🤖 commitmate

**AI 기반 Git 커밋 메시지 자동 생성 도구**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/leehosu/commitmate)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/leehosu/commitmate)](https://github.com/leehosu/commitmate/releases)

[English](../README.md) | [한국어](ko.md)

[특징](#특징) • [빠른 시작](#빠른-시작) • [설치](#설치) • [설정](#설정) • [고급 기능](#고급-기능)

</div>

---

## 특징

- 🤖 **AI 기반** - OpenAI GPT & Anthropic Claude 지원
- 📝 **Conventional Commits** - 업계 표준 커밋 메시지 형식
- 🌏 **다국어 지원** - 한글/영어 메시지 및 UI 지원
- 🎫 **JIRA 통합** - 브랜치 이름에서 이슈 번호 자동 감지
- 🎨 **인터랙티브 UI** - 간단한 프롬프트로 수정, 재생성, 취소
- ⚙️ **유연한 설정** - CLI 또는 환경변수로 설정
- 🚀 **크로스 플랫폼** - Linux, macOS, Windows

## 빠른 시작

```bash
# 1. API 키 설정
commitmate config set-key openai sk-xxxxx
commitmate config set-provider openai

# 2. 변경사항 스테이징
git add .

# 3. 커밋 생성 및 실행
commitmate
```

## 예시

```bash
$ git add .
$ commitmate

🔍 Staged 변경사항 분석 중...
✨ AI가 커밋 메시지를 생성했습니다:

feat(auth): add JWT authentication middleware

? 이 커밋 메시지를 사용하시겠습니까? 
  ▸ Yes - 커밋 실행
    Edit - 메시지 수정
    Regenerate - 다시 생성
    Cancel - 취소

✓ 커밋이 완료되었습니다!
```

## 설치

### Homebrew (추천)

```bash
brew tap leehosu/tap
brew install commitmate
```

### 바이너리 다운로드

[Releases](https://github.com/leehosu/commitmate/releases)에서 다운로드

```bash
# macOS/Linux
tar -xzf commitmate_*.tar.gz
chmod +x commitmate
sudo mv commitmate /usr/local/bin/

# Windows
# commitmate.exe를 압축 해제 후 PATH에 추가
```

## 설정

### 기본 설정

```bash
# API 키
commitmate config set-key openai sk-xxxxx
commitmate config set-key claude sk-ant-xxxxx

# 제공자
commitmate config set-provider openai  # 또는 claude

# 모델 (선택사항)
commitmate config set-model openai gpt-4o-mini
commitmate config set-model claude claude-3-5-haiku-20241022
```

### 언어 설정

```bash
commitmate config set-commit-language ko  # 커밋 메시지 언어 (ko/en)
commitmate config set-ui-language ko      # UI 언어 (ko/en)
```

### 설정 확인

```bash
commitmate config show
```

### 환경변수

```bash
export COMMITMATE_OPENAI_API_KEY=sk-xxxxx
export COMMITMATE_CLAUDE_API_KEY=sk-ant-xxxxx
export COMMITMATE_PROVIDER=openai
export COMMITMATE_COMMIT_LANGUAGE=ko
export COMMITMATE_UI_LANGUAGE=ko
```

## 사용법

```bash
commitmate                    # 분석 및 커밋
commitmate --dry-run          # 메시지만 생성
commitmate --provider openai  # 특정 제공자 사용
commitmate --no-verify        # git hooks 무시
```

## 고급 기능

### Conventional Commits

commitmate는 [Conventional Commits](https://www.conventionalcommits.org/) 형식을 따릅니다:

```
<type>(<scope>): <subject>
```

**지원 타입:** `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`, `perf`, `ci`, `build`, `revert`

### JIRA 통합

브랜치 이름에서 JIRA 이슈 번호를 자동으로 감지합니다:

```bash
# 브랜치: DEVOPS2-430-add-feature
commitmate
# 출력: [DEVOPS2-430] feat: add user authentication

# 브랜치: feature/add-auth
commitmate
# 출력: feat: add user authentication
```

**지원 패턴:** `PROJECT-123`, `ABC-456`, `DEVOPS2-430`

**참고:** `main`, `master`, `develop` 브랜치에서는 JIRA 접두사가 추가되지 않습니다.

### 다국어 지원

커밋 메시지와 UI 언어를 각각 설정할 수 있습니다:

```bash
# 영어 커밋 메시지, 한글 UI (글로벌 팀의 한국인 개발자)
commitmate config set-commit-language en
commitmate config set-ui-language ko

# 한글 커밋 메시지, 영어 UI (한국 회사의 외국인 개발자)
commitmate config set-commit-language ko
commitmate config set-ui-language en
```

## 기여하기

이슈와 PR을 환영합니다! [CONTRIBUTING.md](../CONTRIBUTING.md) 참조

## 변경 이력

릴리즈 히스토리는 [CHANGELOG.md](../CHANGELOG.md) 참조

## 라이선스

MIT License - [LICENSE](../LICENSE) 참조

## 작성자

[@leehosu](https://github.com/leehosu)

---

<div align="center">

**⭐ 도움이 되셨다면 Star를 눌러주세요!**

</div>
