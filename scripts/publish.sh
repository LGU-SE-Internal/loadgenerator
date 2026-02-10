#!/bin/bash
set -e

CHART_DIR="charts/loadgenerator"

ORG_NAME=$(echo $GITHUB_REPOSITORY | cut -d'/' -f1 | tr '[:upper:]' '[:lower:]')
REPO_NAME=$(echo $GITHUB_REPOSITORY | cut -d'/' -f2 | tr '[:upper:]' '[:lower:]')
REPO_URL="https://${ORG_NAME}.github.io/${REPO_NAME}"

echo "Target Repo URL: $REPO_URL"

helm dependency update $CHART_DIR

mkdir -p .deploy
helm package $CHART_DIR -d .deploy

cd .deploy
if [ -f index.yaml ]; then
    helm repo index . --url $REPO_URL --merge index.yaml
else
    helm repo index . --url $REPO_URL
fi
cd ..