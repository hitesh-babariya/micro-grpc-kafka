# start-kind.ps1
# ----------------------------
# PowerShell script to start Kind cluster and deploy Go + Kafka services

# Configuration
$ClusterName = "go-kafka-grpc"
$KindConfig = ".\kind-config.yaml"  # Change path if needed

# Docker images
$Images = @(
    @{ Name = "user-service"; Path = ".\user-service\Dockerfile" },
    @{ Name = "order-service"; Path = ".\order-service\Dockerfile" }
)

# Kubernetes manifests
$K8sDir = ".\k8s"

# ----------------------------
# Step 1: Check if Kind cluster exists
$clusters = & kind get clusters
if ($clusters -contains $ClusterName) {
    Write-Host "Kind cluster '$ClusterName' already exists."
} else {
    Write-Host "Creating Kind cluster '$ClusterName'..."
    & kind create cluster --name $ClusterName --config $KindConfig --wait 60s
}

# Step 2: Build Docker images and load into Kind
foreach ($img in $Images) {
    Write-Host "Building Docker image: $($img.Name)"
    docker build -t "$($img.Name):latest" -f $img.Path .

    Write-Host "Loading image '$($img.Name):latest' into Kind cluster"
    & kind load docker-image "$($img.Name):latest" --name $ClusterName
}

# Step 3: Deploy Kafka (Zookeeper + Kafka)
Write-Host "Deploying Kafka..."
kubectl apply -f "$K8sDir\zookeeper.yaml"
kubectl apply -f "$K8sDir\kafka.yaml"

Write-Host "Waiting for Kafka pods to be ready..."
kubectl wait --for=condition=Ready pod -l app=zookeeper --timeout=120s
kubectl wait --for=condition=Ready pod -l app=my-kafka --timeout=120s

# Step 4: Deploy Go services
Write-Host "Deploying Go services..."
kubectl apply -f $K8sDir

# Step 5: Wait for service pods to be ready
Write-Host "Waiting for service pods to be ready..."
kubectl wait --for=condition=Ready pod -l app=user-service --timeout=120s
kubectl wait --for=condition=Ready pod -l app=order-service --timeout=120s

# Step 6: Show pods and logs
Write-Host "Pods status:"
kubectl get pods -o wide

Write-Host "Logs (last 50 lines):"
kubectl logs -l app=user-service --tail=50
kubectl logs -l app=order-service --tail=50

Write-Host "Deployment complete!"
