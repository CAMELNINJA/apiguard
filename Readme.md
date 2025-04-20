# 🛡️ API Guard — Secure, Self-hosted AI Gateway

### Simple & pluggable reverse proxy for OpenAI, Claude, Mistral, DeepSeek and other AI APIs.  
Built for indie hackers, dev teams, and privacy-conscious products.

![Go](https://img.shields.io/badge/Go-1.21+-blue)
![License](https://img.shields.io/github/license/you/apiguard)
![Status](https://img.shields.io/badge/status-WIP-orange)

---

## 🧩 Features

- 🧠 **Multi-provider AI proxy** — Easily route requests to OpenAI, Claude, Mistral, etc.
- 🔐 **API key & JWT authentication**
- 🚦 **Rate limiting** (per-user or per-route)
- 📊 **Structured logging** (with Zap)
- 🔁 **Dynamic routing** via `config.yaml`
- 🛠️ **Simple middleware-based extension model**
- 🚀 **Self-hostable** (single binary or Docker)

---

## 🔧 Example Use Cases

- 🧑‍💻 Build your own **“AI gateway”** for your SaaS to switch between model providers.
- 👨‍🔬 Internal tool for **testing multiple AI APIs with auth and limits**.
- 💸 Add **rate/billing control** before your LLM usage explodes.
- 🕵️ Run a **privacy-first proxy** in your infra for better observability.

---

## 🚀 Quick Start

1. **Clone & Build**
   ```bash
   git clone https://github.com/your-username/apiguard
   cd apiguard
   go run main.go
   ```

2. **Configure your routes**
   ```yaml
   routes:
     - name: openai
       match_prefix: /v1
       upstream: https://api.openai.com
       auth_required: true
       rate_limit:
         rps: 5
         burst: 10
   ```

3. **Make a request**
   ```bash
   curl -H "Authorization: Bearer YOUR_KEY" http://localhost:8080/v1/chat/completions
   ```

---

## 📦 Roadmap

| Feature                       | Status     |
|------------------------------|------------|
| Reverse proxy engine         | ✅ Done     |
| Config-driven routes         | ✅ Done     |
| API Key auth                 | ✅ Done     |
| JWT auth                     | ⏳ Soon     |
| Rate limiting (in-memory)    | ✅ Done     |
| Redis-backed limiter         | ⏳ Soon     |
| Metrics endpoint             | ⏳ Soon     |
| Minimal web dashboard        | 🧠 Idea     |
| Billing integration (webhook)| 🧠 Idea     |
| Provider-aware fallback      | 🧠 Idea     |
| Usage stats / analytics      | 🧠 Idea     |

---

## ⚙️ Tech Stack

- [Go](https://golang.org/) 1.21+
- [Zap](https://github.com/uber-go/zap) for logging
- `net/http` + `httputil.ReverseProxy`
- Chi for routing
- YAML config parsing via `gopkg.in/yaml.v3`

---

## 🧪 Why not Higress / Kong / KrakenD?

Because:
- They’re **too complex** for simple use cases
- Require K8s / Istio / Envoy / Lua / Wasm
- Not built for **AI-specific control** (tokens, rate, fallback)
- Not hacker-friendly out of the box

---

## ❤️ Who is this for?

| You are...                    | Use this if...                          |
|------------------------------|-----------------------------------------|
| 🧑 Indie hacker               | You need AI API control in side-project |
| 🛠️ SaaS builder               | You proxy traffic to multiple AI APIs   |
| 🕶️ Self-host advocate         | You want full transparency & logs       |
| 🧪 Internal infra engineer    | You want rate limit/auth proxy          |

---

## 🤝 Contribute

Pull requests, ideas, and issues welcome!  
This is a WIP and we’d love help refining the core.  
Focus: Simplicity, performance, extensibility.

---

## 📄 License

MIT © YourNameHere

