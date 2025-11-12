# umimic 🤖

A Go-based API for creating personalized chatbots using OpenRouter, with support for message history and custom personalities.

## Features

- Support for multiple personalities via system messages
- Message history to maintain conversation context
- Integration with OpenRouter (access to various AI models)
- CORS configured for frontend compatibility
- Request validation using go-playground/validator

## Tech Stack

- Go 1.24
- Chi Router (HTTP framework)
- OpenRouter API
- Chi/CORS (middleware)

## Pricing

This API uses OpenRouter to access AI models. The default model is `openai/gpt-4o-mini`, which has the following approximate costs:

| Volume | Input Cost | Output Cost | Estimated Total* |
|--------|-----------|-------------|------------------|
| 1,000 messages | $0.15 | $0.60 | ~$0.75 |
| 10,000 messages | $1.50 | $6.00 | ~$7.50 |

*Based on average message length of 100 input tokens and 400 output tokens.

You can check current pricing and switch to other models at [openrouter.ai/docs/models](https://openrouter.ai/docs/models). OpenRouter supports various models with different price points, from budget-friendly to premium options.

## Local Development

### Prerequisites

- Go 1.24 or higher
- OpenRouter account (to obtain API key)

### Setup

1. Clone the repository
```bash
git clone https://github.com/xyztavo/umimic.git
cd umimic
```

2. Copy .env.example and fill in your credentials
```bash
cp .env.example .env
```

3. Add your environment variables to .env
```env
OPENROUTER_API_KEY=your_key_here
PORT=3000
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

4. Install dependencies
```bash
go mod download
```

5. Run the server
```bash
go run main.go
```

The server will start at `http://localhost:3000`

## Endpoints

### POST /api/message

Sends a message to the bot with optional history.

**Request Body:**
```json
{
  "message": "your message here",
  "history": [
    {
      "role": "system",
      "content": "bot personality here"
    },
    {
      "role": "user",
      "content": "previous message"
    },
    {
      "role": "assistant",
      "content": "previous response"
    }
  ]
}
```

**Response:**
```json
{
  "reply": "bot response"
}
```

### GET /api/healthz

Checks if the API is running.

**Response:**
```json
{
  "status": "ok"
}
```

## Deployment
### DisCloud

DisCloud offers a free plan with 256MB RAM and 1GB storage.

1. Create an account at [discloud.app](https://discloud.app)

2. Create a `discloud.config` file in the project root:
2.1 Register your subdomain on discloud

```env
ID=umimic
TYPE=site
MAIN=main.go
NAME=Umimic
RAM=256
VERSION=latest
```

3. Upload the via github the project with the branch


4. Add environment variables in the panel:
   - `OPENROUTER_API_KEY`
   - `PORT=8080` (DisCloud uses port 8080)
   - `ALLOWED_ORIGINS`

5. DisCloud will automatically build and deploy the application

6. Access via the provided URL like `your-app.discloud.app`

## Project Structure

```
.
├── config/          # Configuration and environment variables
├── mimic/           # Chatbot logic
├── models/          # Request/response structs
├── openrouter/      # OpenRouter API client
├── utils/           # Utilities (validation, etc.)
├── main.go          # API entrypoint
└── go.mod           # Dependencies
```

## How It Works

1. Frontend sends message + history (including personality as system message)
2. Backend appends the current message to history
3. Makes request to OpenRouter with full context
4. Returns the model's response

Personalities are managed on the frontend and sent as the first message in history with role "system".

## Troubleshooting

**CORS error:** Add your frontend origin to `ALLOWED_ORIGINS`

**API timeout:** OpenRouter can sometimes be slow; consider increasing timeout if necessary

**Rate limit:** OpenRouter has per-minute limits; monitor your usage

## Contributing

Feel free to open issues or pull requests.

## License

MIT License - see [LICENSE](LICENSE) file for details.

Made with ❤️ by ustav