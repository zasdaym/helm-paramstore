# helm-paramstore

This is a helm3 plugin to use values from AWS SSM Parameter Store on Helm values file.

## Install

```
helm plugin install https://github.com/zasdaym/helm-paramstore
```

## Usage

```
helm paramstore -f values-custom.yaml | helm upgrade release-name repo/chart --install --values -
```
