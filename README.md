# go-run-reports

A usage tracking system for Xapien customers to enforce and monitor monthly report usage based on subscription plans. Built with Go and designed to run in AWS (e.g., EKS with DynamoDB integration).

## 🧱 Features

- Set subscription plans for organizations
- Enforce organization and per-user report limits
- Handle monthly top-ups
- Simulated notifications for reaching usage thresholds
- RESTful API architecture
- AWS-native design (DynamoDB, IAM, EKS)

## 📁 Project Structure
```bash
go-run-reports/
├── cmd/server           # Main app entrypoint
├── internal/api         # HTTP handlers
├── internal/service     # Business logic
├── internal/repository  # DynamoDB interface
├── internal/model       # Domain models
├── internal/config      # Config loader
├── internal/notifier    # Simulated notification
├── pkg/util             # Time/month utilities
```

## 🚀 Quick Start

### 1. Clone and Initialize

```bash
git clone https://github.com/yourname/go-run-reports.git
cd go-run-reports
go mod tidy

2. Run Locally

go run ./cmd/server

Server starts on http://localhost:8000

3. Environment Variables

You may use a .env file or your preferred config loader to set:
	•	AWS_REGION
	•	DYNAMODB_TABLE_NAME
	•	ENVIRONMENT=dev|staging|prod

4. Build Docker Image

docker build -t go-run-reports .

5. Deploy to EKS (via Helm)

helm install go-run-reports ./helm-chart \
  --set image.repository=your-repo/go-run-reports \
  --set image.tag=latest

🔌 REST API Endpoints

Method	Endpoint	Description
POST	/plan	Set organization plan
POST	/limit	Set per-user monthly limit
POST	/use	Register report usage
POST	/top-up	Top up monthly report credits

🧪 Testing

go test ./...

📘 Notes
	•	DynamoDB tables should have Org ID as primary key and optionally User ID as sort key
	•	Time-based logic resets usage at the start of each calendar month
	•	Notification logic is mocked for now — can be extended with SNS/SQS later

⸻