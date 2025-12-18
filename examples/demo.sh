#!/bin/bash

# commitgen 데모 스크립트

echo "🎬 commitgen 데모 시작!"
echo ""

# 1. 버전 확인
echo "📌 Step 1: 버전 확인"
commitgen version
echo ""

# 2. 설정 확인
echo "📌 Step 2: 현재 설정 확인"
commitgen config show
echo ""

# 3. API 키 설정 (예시 - 실제로는 실행하지 않음)
echo "📌 Step 3: API 키 설정 방법"
echo "  commitgen config set-key openai sk-your-api-key"
echo "  commitgen config set-provider openai"
echo ""

# 4. 도움말 확인
echo "📌 Step 4: 도움말 확인"
commitgen --help
echo ""

# 5. 사용 예시
echo "📌 Step 5: 기본 사용법"
echo "  1. 코드 변경 후:"
echo "     git add ."
echo ""
echo "  2. commitgen 실행:"
echo "     commitgen"
echo ""
echo "  3. AI가 생성한 커밋 메시지 확인 및 승인"
echo ""

# 6. 고급 옵션
echo "📌 Step 6: 고급 옵션"
echo "  - dry-run (커밋하지 않고 메시지만 생성):"
echo "    commitgen --dry-run"
echo ""
echo "  - 특정 제공자 사용:"
echo "    commitgen --provider claude"
echo ""
echo "  - git hooks 무시:"
echo "    commitgen --no-verify"
echo ""

echo "✨ 데모 완료! 실제로 사용하려면 API 키를 설정하세요."
