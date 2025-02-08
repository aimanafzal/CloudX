# CloudX API

## Overview
CloudX API is a multi-cloud infrastructure management API that enables users to create and manage cloud resources across **GCP, AWS, and Azure**. This project provides RESTful endpoints to automate provisioning of compute instances, containerized applications, and NoSQL databases.

## Features
- 🚀 **Multi-cloud support** (GCP, AWS, Azure)
- 🏗️ **Compute instance management** (GCP Compute Engine, AWS EC2, Azure VMs)
- 📦 **Container orchestration** (Cloud Run, AWS ECS, Azure Container Apps)
- 🗄️ **Database integration** (Firestore, DynamoDB, CosmosDB)
- ⚡ **Fast & lightweight** (Built with Go for efficiency and concurrency)

## Tech Stack
- **Backend**: Golang
- **Infrastructure**: Terraform (for provisioning cloud resources)
- **Deployment**: Docker
- **API Framework**: Chi or Gin (lightweight Go web frameworks)

## Project Structure
```
cloudx-api/
├── cmd/
│   ├── main.go                # Entry point for the API service
│   ├── handlers/
│   │   ├── compute.go         # Compute instance API handlers
│   │   ├── container.go       # Container orchestration handlers
│   │   ├── database.go        # Firestore/DynamoDB handlers
│   ├── services/
│   │   ├── compute_service.go # Compute instance logic
│   │   ├── container_service.go # Container orchestration logic
│   │   ├── database_service.go # Firestore/DynamoDB logic
├── config/
│   ├── config.go              # Configuration loading
├── infra/
│   ├── main.tf                # Terraform scripts
├── Dockerfile                 # Docker configuration
├── go.mod                      # Go module dependencies
├── README.md                   # Documentation & setup instructions
```

## Setup Instructions
### Prerequisites
- **Go 1.21+** installed
- **Terraform** installed
- **Docker** installed
- Cloud provider authentication (GCP, AWS, Azure CLI configured)

### Installation
1. **Clone the repository**
   ```sh
   git clone git@github.com:aimanafzal/CloudX.git
   cd cloudx-api
   ```
2. **Install dependencies**
   ```sh
   go mod tidy
   ```
3. **Set up cloud provider credentials**
   - **For GCP:** Authenticate with `gcloud auth application-default login`
   - **For AWS:** Configure `aws configure`
   - **For Azure:** Authenticate with `az login`

### Running the API
```sh
go run cmd/main.go
```

### Deploying with Docker
```sh
docker build -t cloudx-api .
docker run -p 8080:8080 cloudx-api
```

### Infrastructure Setup (Terraform)
1. Navigate to the **infra** directory:
   ```sh
   cd infra
   ```
2. Initialize Terraform:
   ```sh
   terraform init
   ```
3. Apply Terraform configuration to provision resources:
   ```sh
   terraform apply -auto-approve
   ```

## API Endpoints
### Compute Engine (GCP) / EC2 (AWS)
- **Create instance:** `POST /api/compute`
- **Get instance:** `GET /api/compute/{id}`
- **Delete instance:** `DELETE /api/compute/{id}`

### Cloud Run (GCP) / ECS (AWS)
- **Deploy service:** `POST /api/container`
- **Get service:** `GET /api/container/{id}`
- **Delete service:** `DELETE /api/container/{id}`

### Firestore (GCP) / DynamoDB (AWS)
- **Create record:** `POST /api/database`
- **Get record:** `GET /api/database/{id}`
- **Delete record:** `DELETE /api/database/{id}`

## Contributing
1. Fork the repository
2. Create a feature branch (`git checkout -b feature-name`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the branch (`git push origin feature-name`)
5. Create a pull request

## License
MIT License

