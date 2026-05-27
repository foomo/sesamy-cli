# Google API setup

`sesamy-cli` talks to the GTM API v2. You need a Google Cloud **Service Account** with access to your GTM account.

## 1. Create the service account

In the [Google Cloud Console](https://console.cloud.google.com/iam-admin/serviceaccounts):

1. Pick (or create) a project.
2. **Create service account** → give it a name like `sesamy-cli`.
3. Skip the optional "grant access to project" step — GTM permissions are granted in the GTM UI, not via IAM.
4. **Keys** → **Add key** → **JSON**. Save the file somewhere safe.

## 2. Enable the GTM API

In the [API Library](https://console.cloud.google.com/apis/library/tagmanager.googleapis.com), enable **Tag Manager API**.

## 3. Add the service account to GTM

In [Tag Manager](https://tagmanager.google.com):

1. **Admin → User Management** (account level).
2. Add the service account's email (`<name>@<project>.iam.gserviceaccount.com`).
3. Grant **Publish** permission on the account, and **Publish** on each container you want the CLI to manage.

## 4. Point sesamy at the credentials

```yaml
googleApi:
  credentialsFile: ./google_service_account_creds.json
  requestQuota: 15
```

Or inline:

```yaml
googleApi:
  credentials: |
    {"type":"service_account", ... }
  requestQuota: 15
```

Use `credentialsFile` locally and inline `credentials` (from a secret) in CI.

## 5. Increase the quota (optional but recommended)

Default GTM API quota is **15 requests per minute**. For non-trivial configs:

1. Go to **APIs & Services → Tag Manager API → Quotas** in Cloud Console.
2. Find the per-minute request quota and request an increase. Google typically approves within a day or two.
3. Update `googleApi.requestQuota` to match.

## Troubleshooting

| Symptom | Fix |
|---------|-----|
| `403 PERMISSION_DENIED` | Service account email not added to the GTM account, or missing **Publish** rights on the container. |
| `429 Too Many Requests` | Request quota too low. Either lower `requestQuota` to be safe, or raise the quota in Cloud Console. |
| `401 invalid_grant` | Credentials JSON is malformed or the service account key was deleted. |
| `404 NOT_FOUND` on a container | `accountId` / `containerId` mismatch — double-check IDs in the GTM UI URL. |
