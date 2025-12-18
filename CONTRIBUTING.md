# 기여 가이드

commitgen 프로젝트에 기여해주셔서 감사합니다! 🎉

## 개발 환경 설정

### 필수 요구사항

- Go 1.21 이상
- Git
- (선택) GoReleaser (릴리즈 테스트용)

### 로컬 개발

```bash
# 저장소 클론
git clone https://github.com/leehosu/commitgen.git
cd commitgen

# 의존성 설치
go mod download

# 빌드
go build -o commitgen

# 실행
./commitgen --help
```

## 코드 기여하기

### 브랜치 전략

- `main`: 안정 버전
- `feature/*`: 새로운 기능
- `fix/*`: 버그 수정
- `docs/*`: 문서 변경

### Pull Request 프로세스

1. 이슈를 먼저 생성하여 변경사항을 논의합니다
2. Fork 후 브랜치를 생성합니다
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. 변경사항을 커밋합니다 (Conventional Commits 형식)
   ```bash
   git commit -m "feat: add amazing feature"
   ```
4. 브랜치를 Push 합니다
   ```bash
   git push origin feature/amazing-feature
   ```
5. Pull Request를 생성합니다

### 커밋 메시지 규칙

Conventional Commits 형식을 따릅니다:

```
<type>(<scope>): <subject>

[optional body]

[optional footer]
```

**타입:**
- `feat`: 새로운 기능
- `fix`: 버그 수정
- `docs`: 문서 변경
- `style`: 코드 포맷팅
- `refactor`: 리팩토링
- `test`: 테스트 추가/수정
- `chore`: 빌드, 설정 등

**예시:**
```
feat(ai): add support for GPT-4o-mini
fix(config): resolve config file parsing error
docs(readme): update installation instructions
```

## 테스트

```bash
# 단위 테스트 실행
go test ./...

# 특정 패키지 테스트
go test ./internal/ai

# 커버리지 확인
go test -cover ./...
```

## 코드 스타일

- `gofmt`로 포맷팅
- `golangci-lint`로 린팅
- 명확하고 의미있는 변수/함수명 사용
- 주석은 한국어 또는 영어로 작성

## 이슈 리포팅

버그를 발견하셨나요? 이슈를 생성해주세요!

**포함할 내용:**
- 예상 동작 vs 실제 동작
- 재현 단계
- 환경 정보 (OS, Go 버전 등)
- 에러 메시지 및 로그

## 새로운 AI 제공자 추가하기

새로운 AI 제공자를 추가하려면:

1. `internal/ai/` 디렉토리에 새 파일 생성 (예: `gemini.go`)
2. `Client` 인터페이스 구현
3. `NewClient()` 함수에 새 제공자 추가
4. 설정 구조체에 새 제공자 추가
5. 테스트 작성
6. README 업데이트

## 라이선스

기여하신 코드는 MIT 라이선스 하에 배포됩니다.

## 질문이 있으신가요?

이슈를 생성하거나 토론을 시작해주세요!
