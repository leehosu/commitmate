# commitgen 프로젝트 요약

## 📊 프로젝트 정보

- **프로젝트명**: commitgen
- **설명**: AI 기반 Git 커밋 메시지 자동 생성 도구
- **언어**: Go 1.21+
- **라이선스**: MIT
- **저장소**: https://github.com/leehosu/commitgen

## ✨ 주요 기능

### 핵심 기능
1. **AI 기반 커밋 메시지 생성**
   - OpenAI GPT-4o 지원
   - Anthropic Claude 3.5 Sonnet 지원
   - 자동 diff 분석 및 메시지 생성

2. **Conventional Commits 형식**
   - feat, fix, docs, style, refactor 등 표준 타입 지원
   - 구조화된 커밋 메시지 생성

3. **인터랙티브 UI**
   - 승인/수정/재생성/취소 옵션
   - promptui 기반 사용자 친화적 인터페이스

4. **유연한 설정 관리**
   - YAML 기반 설정 파일
   - 환경변수 지원
   - 다중 AI 제공자 지원

5. **크로스 플랫폼**
   - Linux (amd64, arm64)
   - macOS (amd64, arm64)
   - Windows (amd64)

## 🏗️ 아키텍처

### 디렉토리 구조
```
commitgen/
├── cmd/                    # CLI 명령어
│   ├── root.go            # 루트 명령어
│   ├── commit.go          # 커밋 생성
│   ├── config.go          # 설정 관리
│   └── version.go         # 버전 정보
├── internal/              # 내부 패키지
│   ├── ai/               # AI 클라이언트
│   │   ├── client.go     # 인터페이스
│   │   ├── openai.go     # OpenAI 구현
│   │   └── claude.go     # Claude 구현
│   ├── config/           # 설정 관리
│   ├── git/              # Git 인터페이스
│   └── prompt/           # 프롬프트 생성
├── .github/workflows/    # CI/CD
├── examples/             # 예시 및 데모
└── docs/                 # 문서
```

### 기술 스택
- **CLI 프레임워크**: Cobra
- **설정 관리**: Viper
- **AI SDK**: 
  - OpenAI: go-openai
  - Claude: 직접 HTTP 구현
- **UI**: promptui
- **빌드/배포**: GoReleaser, GitHub Actions

## 📈 통계

- **Go 파일 수**: 12개
- **총 코드 라인**: ~978 라인
- **커밋 수**: 3개
- **문서**: README, QUICKSTART, CONTRIBUTING, CHANGELOG
- **의존성**: 주요 외부 패키지 6개

## 🎯 사용 흐름

```
1. 사용자가 파일 수정
   ↓
2. git add . 로 staging
   ↓
3. commitgen 실행
   ↓
4. Git 저장소 확인
   ↓
5. Staged diff 추출
   ↓
6. 설정 로드 (API 키, 제공자)
   ↓
7. AI 클라이언트 생성
   ↓
8. 프롬프트 생성 (시스템 + 사용자)
   ↓
9. AI API 호출
   ↓
10. 커밋 메시지 생성
    ↓
11. 사용자에게 미리보기 제공
    ↓
12. 사용자 선택:
    - Yes: 커밋 실행
    - Edit: 수정 후 커밋
    - Regenerate: 다시 생성
    - Cancel: 취소
```

## 🚀 배포 전략

### GitHub Releases
- GoReleaser 자동 빌드
- 멀티 플랫폼 바이너리
- 체크섬 파일 제공
- 변경 이력 자동 생성

### CI/CD
- **Test Workflow**: PR 및 main 푸시 시 자동 테스트
- **Release Workflow**: 태그 푸시 시 자동 릴리즈
- 멀티 OS/Go 버전 테스트 매트릭스

## 💡 주요 설계 결정

1. **Go 언어 선택**
   - 크로스 컴파일 용이
   - 단일 바이너리 배포
   - 빠른 실행 속도

2. **Cobra CLI 프레임워크**
   - 표준 CLI 패턴
   - 서브커맨드 구조
   - 풍부한 기능

3. **AI 제공자 추상화**
   - 인터페이스 기반 설계
   - 쉬운 확장성
   - 제공자 전환 용이

4. **Claude HTTP 직접 구현**
   - SDK 복잡도 회피
   - 세밀한 제어
   - 의존성 최소화

## 📝 다음 단계 (향후 개선)

### v0.2.0
- [ ] 단위 테스트 추가
- [ ] 통합 테스트
- [ ] 설정 검증 강화

### v0.3.0
- [ ] Gemini 지원 추가
- [ ] 커밋 메시지 템플릿 커스터마이징
- [ ] 다국어 지원 (한국어, 일본어 등)

### v0.4.0
- [ ] 커밋 히스토리 분석
- [ ] 프로젝트별 설정
- [ ] 브랜치별 커밋 스타일

### v1.0.0
- [ ] Homebrew 배포
- [ ] 플러그인 시스템
- [ ] 웹 UI (선택사항)

## 🤝 기여 방법

1. 이슈 생성
2. Fork 및 브랜치 생성
3. 변경사항 커밋
4. Pull Request 제출

상세한 내용은 [CONTRIBUTING.md](CONTRIBUTING.md) 참조

## 📞 지원

- **이슈**: https://github.com/leehosu/commitgen/issues
- **토론**: https://github.com/leehosu/commitgen/discussions
- **이메일**: (저장소 프로필 참조)

## 🎉 완료!

이제 commitgen을 사용하여 더 나은 커밋 메시지를 자동으로 생성할 수 있습니다!

```bash
# 빠른 시작
git add .
commitgen
```

Happy Committing! 🚀
