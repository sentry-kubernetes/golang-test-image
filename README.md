Creating kubernetes secret:

Replace your-base64-encoded-sentry-dsn with your base64 encoded DSN:
```
echo -n “your-sentry-dsn” | base64 -w0
```

Apply secret:
```
kubectl apply -f golang-example-dsn.yaml
```

Change version in image and apply deployment:
```
kubectl apply -f golang-example.yaml
```

Notes:
Make sure you replace “your-sentry-dsn” with your real DSN.

Make sure you have access to your Kubernetes cluster configured and kubectl installed.

This code and instructions will allow you to generate a large number of different exceptions and send them to Sentry for benchmarking, and deploy this process to Kubernetes.

