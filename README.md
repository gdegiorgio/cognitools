# 🌐 Cognito Tools




> ⚠️ **Early Stage Project**  
> Cognito Tools is under **active development** and not yet production-ready. Expect frequent changes, bugs, and missing features. Contributions are welcome as we build this together!


Cognito Tools is a cross-platform CLI tool for automating OAuth JWT generation with AWS Cognito User Pools. It simplifies the process of obtaining OAuth tokens for your AWS Cognito applications.


---

## ✨ Features

- 🔄 Automated OAuth token generation
- 🛠️ Written in Go for speed and portability
- 🌍 Cross-platform support (Linux, macOS, Windows)

---

## 🚀 Getting Started


#### Windows

```
powershell -c "irm https://raw.githubusercontent.com/gdegiorgio/cognitools/refs/heads/main/scripts/install.ps1 | iex"
```

#### MacOs & Linux

```bash
curl https://raw.githubusercontent.com/gdegiorgio/cognitools/refs/heads/main/scripts/install.sh | bash
```


### 🧪 Usage

```bash
Cognito Tools - CLI for AWS Cognito OAuth JWT Generation

Usage:
  cognitools [command]

Available Commands:
  generate    Generate OAuth JWT token
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Show current cognitools CLI version

Flags:
  -h, --help   help for cognitools
```

### 🧙 Contributing

We welcome contributions of all kinds!

- All PRs must follow Semantic Commit Messages
- Read our [Contributing guide](CONTRIBUTING.md) before submitting a PR
