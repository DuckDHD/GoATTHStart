# Go + Chi + Templ Full-Stack Starter

A **full-stack Go template** for building modern, **HTML-first SSR apps** with **HTMX** and **Alpine.js**.
Everything runs on a single lightweight Chi router — no SPA complexity.

## ✨ Features

* **Chi** for unified routing (SSR pages + JSON API)
* **Templ** for type-safe, component-based server rendering
* **HTMX + Alpine.js** for progressive, reactive UIs
* **TailwindCSS** for styling
* **GORM** for ORM and database access with auto migrations
* **Zap** for structured, high-performance logging
* **Viper** for flexible configuration management
* **Authentication & Authorization**: cookie sessions, JWT-based API tokens, roles, and permissions
* **Notifications** with real-time updates via **SSE**
* **Chat system** using **WebSockets**
* Clean, domain-driven structure for long-term scalability

---

## 🗂 Project Structure

```
/cmd
  /app           # main entrypoint
/internal
  /app           # configuration (Viper), logging (Zap), bootstrapping
  /domain        # core business logic (auth, chat, notifications, etc.)
  /http
    /routes      # Chi route definitions
    /views       # Templ templates for SSR
    /middleware  # auth, csrf, logging, etc.
  /repo          # database repositories using GORM
/assets          # TailwindCSS input, icons
/public          # Compiled CSS/JS, static files
/migrations      # SQL migrations (optional if you add Goose later)
```

---

## 🚀 Getting Started

### **Requirements**

* Go 1.22+
* PostgreSQL (or SQLite for local development)
* [TailwindCSS CLI](https://tailwindcss.com/docs/installation)

### **Setup**

```bash
# 1. Clone the repository
git clone https://github.com/yourusername/yourproject.git
cd yourproject

# 2. Create environment file
cp .env.example .env

# 3. Run DB migrations (optional) or let GORM auto migrate
make migrate-up

# 4. Start Tailwind in watch mode
make tailwind

# 5. Start the server
make run
```

Visit **[http://localhost:8080](http://localhost:8080)** to view your app.

---

## ⚙️ Configuration

Configuration is handled by **Viper**, supporting environment variables and `.env` files.

| Variable        | Default         | Description                         |
| --------------- | --------------- | ----------------------------------- |
| APP\_ENV        | dev             | Application environment (dev, prod) |
| PORT            | 8080            | Server port                         |
| DB\_URL         | postgres\://... | Database connection URL             |
| SESSION\_SECRET | change-me       | Secret for session cookies          |
| JWT\_SECRET     | change-me       | Secret for signing JWT tokens       |
| LOG\_LEVEL      | info            | Log level for Zap                   |

Example `.env`:

```
APP_ENV=dev
PORT=8080
DB_URL=postgres://postgres:postgres@localhost:5432/app?sslmode=disable
SESSION_SECRET=supersecret
JWT_SECRET=supersecretjwt
LOG_LEVEL=debug
```

---

## 🧪 Development

* **Run tests**

  ```bash
  make test
  ```

* **Lint & format**

  ```bash
  make lint
  make fmt
  ```

---

## 🔐 Security

* Secure cookies (`Secure`, `HttpOnly`, `SameSite`)
* CSRF protection for SSR forms and HTMX requests
* JWT for stateless API authentication
* Argon2id password hashing
* Rate limiting and account lockout hooks

---

## 🗺 Roadmap

* [ ] OAuth integration (Google, Microsoft)
* [ ] File uploads via S3-compatible storage
* [ ] Web push notifications
* [ ] Multi-tenant organization support
* [ ] Transition from GORM to Goose + sqlc for type-safe queries

---

## 🤝 Contributing

Contributions are welcome!
Please open an issue or submit a pull request following the [Contributing Guide](CONTRIBUTING.md).

---

## 📜 License

This project is open source and available under the [MIT License](https://github.com/DuckDHD/GoATTHStart/blob/main/LICENSE).
