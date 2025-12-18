# 빠른 시작 가이드

commitgen을 5분 안에 시작해보세요! 🚀

## 1단계: 설치

### 옵션 A: 바이너리 다운로드 (권장)

[Releases 페이지](https://github.com/leehosu/commitgen/releases)에서 OS에 맞는 파일을 다운로드하세요.

**macOS/Linux:**
```bash
# 다운로드 (예시: macOS ARM64)
wget https://github.com/leehosu/commitgen/releases/download/v0.1.0/commitgen_0.1.0_Darwin_arm64.tar.gz

# 압축 해제
tar -xzf commitgen_0.1.0_Darwin_arm64.tar.gz

# 설치
chmod +x commitgen
sudo mv commitgen /usr/local/bin/

# 확인
commitgen version
```

**Windows:**
```powershell
# 다운로드 후 압축 해제
# commitgen.exe를 PATH에 추가

# 확인
commitgen version
```

### 옵션 B: Go install

```bash
go install github.com/leehosu/commitgen@latest
```

## 2단계: API 키 설정

### OpenAI 사용

1. [OpenAI API Keys](https://platform.openai.com/api-keys)에서 API 키 생성
2. 키 설정:
   ```bash
   commitgen config set-key openai sk-xxxxx
   commitgen config set-provider openai
   ```

### Claude 사용

1. [Anthropic Console](https://console.anthropic.com/)에서 API 키 생성
2. 키 설정:
   ```bash
   commitgen config set-key claude sk-ant-xxxxx
   commitgen config set-provider claude
   ```

### 환경변수로 설정 (선택사항)

```bash
# OpenAI
export COMMITGEN_OPENAI_API_KEY=sk-xxxxx
export COMMITGEN_PROVIDER=openai

# Claude
export COMMITGEN_CLAUDE_API_KEY=sk-ant-xxxxx
export COMMITGEN_PROVIDER=claude
```

## 3단계: 첫 커밋 생성

```bash
# Git 저장소로 이동
cd your-project

# 파일 수정
echo "console.log('Hello, World!');" > index.js

# Staging
git add index.js

# commitgen 실행!
commitgen
```

출력 예시:
```
🔍 Git 저장소 확인 중...
📝 Staged 변경사항 분석 중...
✨ AI가 커밋 메시지를 생성하는 중...

✨ 생성된 커밋 메시지:

feat(core): add hello world console log

? 이 커밋 메시지를 사용하시겠습니까? 
  ▸ ✓ Yes - 커밋 실행
    ✎ Edit - 수정 후 커밋
    ↻ Regenerate - 다시 생성
    ✗ Cancel - 취소
```

## 4단계: 고급 사용법

### 커밋하지 않고 메시지만 확인
```bash
commitgen --dry-run
```

### 특정 AI 제공자 사용 (일회성)
```bash
commitgen --provider claude
```

### Git hooks 무시
```bash
commitgen --no-verify
```

### 모델 변경
```bash
# 더 저렴한 모델 사용
commitgen config set-model openai gpt-4o-mini

# Claude Haiku (빠르고 저렴)
commitgen config set-model claude claude-3-5-haiku-20241022
```

## 5단계: 설정 확인

```bash
# 현재 설정 확인
commitgen config show
```

출력 예시:
```
📋 현재 설정:

기본 제공자: openai
언어: en
템플릿: conventional

OpenAI 설정:
  API 키: sk-pr...xK4r
  모델: gpt-4o
  Max Tokens: 150

Claude 설정:
  API 키: (설정되지 않음)
  모델: claude-3-5-sonnet-20241022
  Max Tokens: 150
```

## 트러블슈팅

### "Git 저장소가 아닙니다"
```bash
git init  # Git 저장소 초기화
```

### "커밋할 변경사항이 없습니다"
```bash
git add .  # 파일을 staging area에 추가
```

### "API 키가 설정되지 않았습니다"
```bash
commitgen config set-key openai sk-xxxxx
```

### API 오류 (401 Unauthorized)
- API 키가 올바른지 확인
- API 키에 사용 권한이 있는지 확인
- OpenAI의 경우 크레딧이 남아있는지 확인

## 다음 단계

- [README.md](README.md) - 전체 문서
- [CONTRIBUTING.md](CONTRIBUTING.md) - 기여 가이드
- [GitHub Issues](https://github.com/leehosu/commitgen/issues) - 버그 리포트 및 기능 요청

## 팁

1. **알리아스 설정**
   ```bash
   alias cg='commitgen'
   ```

2. **설정 파일 위치**
   ```
   ~/.commitgen/config.yaml
   ```

3. **여러 프로젝트에서 다른 제공자 사용**
   ```bash
   # 프로젝트 A
   commitgen --provider openai
   
   # 프로젝트 B
   commitgen --provider claude
   ```

4. **CI/CD에서 사용 금지**
   - commitgen은 로컬 개발 도구입니다
   - CI/CD에서는 사용하지 마세요 (API 비용 및 보안)

즐거운 커밋하세요! 🎉
