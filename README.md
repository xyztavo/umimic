# umimic 🤖

Go API for creating personalized chatbots using OpenRouter with message history and custom personalities.

## Stack

Go 1.24 • Chi Router • OpenRouter API • CORS

## Pricing

Uses OpenRouter with `gpt-4o-mini` by default:
- ~$0.75 per 1K messages
- ~$7.50 per 10K messages

Check other models at [openrouter.ai/docs/models](https://openrouter.ai/docs/models)

## Quick Start

```bash
git clone https://github.com/xyztavo/umimic.git
cd umimic
cp .env.example .env
```

Edit `.env`:
```env
OPENROUTER_API_KEY=your_key
PORT=3000
ALLOWED_ORIGINS=http://localhost:3000
```

Run:
```bash
go mod download
go run main.go
```

## API

**POST /api/message**
```json
{
  "message": "your message",
  "history": [
    { "role": "system", "content": "personality" },
    { "role": "user", "content": "previous msg" },
    { "role": "assistant", "content": "previous response" }
  ]
}
```

Response: `{ "reply": "bot response" }`

**GET /api/healthz** - Health check

## Deploy

### DisCloud
1. Create `discloud.config`:
1.1. Create subdomain which is the same as ID
```
ID=umimic
TYPE=site
MAIN=main.go
RAM=256
VERSION=latest
```
2. Connect GitHub at [discloud.app](https://discloud.app)
3. Add env vars: `OPENROUTER_API_KEY`, `PORT=8080`, `ALLOWED_ORIGINS`

## How It Works

Frontend → Backend (+ history) → OpenRouter → Response

Personalities are sent as first message with `role: "system"` in history.

## License

MIT License - see [LICENSE](LICENSE) file for details.

Made with ❤️ by ustav