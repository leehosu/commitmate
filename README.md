# commitgen

🤖 AI 기반 Git 커밋 메시지 자동 생성 도구

## 특징

- ✨ **AI 기반 커밋 메시지 생성**: OpenAI GPT 및 Anthropic Claude 지원
- 📝 **Conventional Commits 포맷**: 업계 표준 커밋 메시지 형식
- 🎯 **간단한 사용법**: 한 번의 명령어로 커밋 완료
- ⚙️ **유연한 설정**: API 키 및 제공자 선택 가능
- 🚀 **크로스 플랫폼**: Linux, macOS, Windows 지원

## 설치

### 바이너리 다운로드 (추천)

최신 릴리즈에서 OS에 맞는 바이너리를 다운로드하세요:
[Releases](https://github.com/leehosu/commitgen/releases)

```bash
# macOS/Linux
chmod +x commitgen
sudo mv commitgen /usr/local/bin/

# Windows
# commitgen.exe를 PATH에 추가
```

### Go install

```bash
go install github.com/leehosu/commitgen@latest
```

## 빠른 시작

### 1. API 키 설정

**OpenAI 사용:**
```bash
commitgen config set-key openai sk-xxxxx
commitgen config set-provider openai
```

**Claude 사용:**
```bash
commitgen config set-key claude sk-ant-xxxxx
commitgen config set-provider claude
```

### 2. 커밋 생성

```bash
# 파일 변경 후
git add .

# AI가 자동으로 커밋 메시지 생성 및 커밋
commitgen
```

## 사용법

### 기본 명령어

```bash
# 기본 사용 (staged 변경사항 분석 및 자동 커밋)
commitgen

# 커밋 메시지만 생성하고 커밋하지 않음
commitgen --dry-run

# 특정 AI 제공자 사용 (일회성)
commitgen --provider openai
commitgen --provider claude

# git hooks 무시
commitgen --no-verify
```

### 설정 관리

```bash
# API 키 설정
commitgen config set-key openai sk-xxxxx
commitgen config set-key claude sk-ant-xxxxx

# 기본 제공자 설정
commitgen config set-provider openai

# 모델 변경
commitgen config set-model openai gpt-4o-mini
commitgen config set-model claude claude-3-5-haiku-20241022

# 현재 설정 확인
commitgen config show

# 버전 확인
commitgen version
```

### 환경변수

설정 파일 대신 환경변수로도 설정 가능:

```bash
export COMMITGEN_OPENAI_API_KEY=sk-xxxxx
export COMMITGEN_CLAUDE_API_KEY=sk-ant-xxxxx
export COMMITGEN_PROVIDER=openai
```

## 설정 파일

설정은 `~/.commitgen/config.yaml`에 저장됩니다:

```yaml
provider: openai
language: en
template: conventional

openai:
  api_key: sk-xxxxx
  model: gpt-4o
  max_tokens: 100

claude:
  api_key: sk-ant-xxxxx
  model: claude-3-5-sonnet-20241022
  max_tokens: 100
```

## Conventional Commits

commitgen은 [Conventional Commits](https://www.conventionalcommits.org/) 형식을 따릅니다:

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

## 예시

```bash
$ git add .
$ commitgen

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

## 개발

```bash
# 저장소 클론
git clone https://github.com/leehosu/commitgen.git
cd commitgen

# 의존성 설치
go mod download

# 빌드
go build -o commitgen

# 실행
./commitgen
```

## 라이선스

MIT License - [LICENSE](LICENSE) 파일 참조

## 기여

이슈와 PR을 환영합니다!

## 작성자

[@leehosu](https://github.com/leehosu)
