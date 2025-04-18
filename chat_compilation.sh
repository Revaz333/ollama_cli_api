#!/bin/bash

if [ -z "$1" ]; then
  echo "❌ Использование: $0 \"Ваш вопрос к модели\""
  exit 1
fi

MSG="$1"
MODEL="avtobaza-v12.3"

ollama serve > /dev/null 2>&1 &
OLLAMA_PID=$!

sleep 3

curl -s -X POST http://localhost:11434/api/chat \
  -H "Content-Type: application/json" \
  -d "{
        \"model\": \"${MODEL}\",
        \"messages\": [
            {
                \"role\": \"user\",
                \"content\": \"${MSG}\"
            }
        ],
        \"stream\": false
    }"

ollama stop $MODEL
pkill ollama