#!/bin/bash -xe
export SA_EMAIL=$(gcloud iam service-accounts list --filter="displayName:spinnaker-account" --format='value(email)')
export PROJECT=$(gcloud info --format='value(config.project)')
export BUCKET=$PROJECT-spinnaker-config

gcloud projects remove-iam-policy-binding $PROJECT --role roles/storage.admin --member serviceAccount:$SA_EMAIL || true
echo y | gcloud iam service-accounts delete $SA_EMAIL || true
echo y | gcloud container clusters delete spinnaker-tutorial --zone us-central1-f || true
echo y | gcloud source repos delete sample-app || true
gcloud beta pubsub subscriptions delete gcr-triggers || true
gsutil -m rm -r gs://$BUCKET || true
rm -f spinnaker-sa.json
rm -f spinnaker-config.yaml
rm -f helm-v2.10.0-linux-amd64.tar.gz
rm -rf linux-amd64
rm -f helm
