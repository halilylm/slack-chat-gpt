
# Slack GPT bot
Slack bot talks with you via ChatGPT.

## Installation 

Clone the project

```bash
git clone https://github.com/halilylm/slack-chat-gpt
```

Generate `.env` file from `.env.example` and fill in fields.

```bash
cp .env.example .env
```

#### Run with docker

```bash
docker build -t slack-gpt-bot .
docker run slack-gpt-bot
```

#### Run without docker

```bash
go run *.go
```
## Demo
![Demo](ss.png)
