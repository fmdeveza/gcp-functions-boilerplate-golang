# Google Cloud Function Template 
Simple boilerplate code for [Google Cloud Functions](https://cloud.google.com/functions) written in Go

## Running local

Just go run it and it serves at port `1323` by default:
```console
$ go run cmd/main.go
2025/01/07 18:56:36 Running HTTP server at port "1323"
Serving function: ""
```

Its now running...
```console
$ curl -X POST -H \
  "Authorization: Bearer $(gcloud auth print-identity-token)" \
  --header 'Content-Type: application/json' \
  --data-raw "{\"message\":{\"data\":\"`echo '{}' | base64`\"}}" \
  localhost:1323/
Done%
```

## Deploy

```console
$ gcloud functions deploy "$GCP_SERVICE_NAME" \
  --project="$GCP_PROJECT_ID" \
  --region="$GCP_REGION" \
  --source="$CI_PROJECT_DIR" \
  --runtime="go122" \
  --entry-point="Trigger" \
  --trigger-http
Deploying function (may take a while - up to 2 minutes)...
..
```
