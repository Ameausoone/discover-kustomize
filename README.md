# Discover Kustomize 

## Run the prez

Install [demoit], then run `demoit`.

## What is it ?

> kustomize lets you customize raw, template-free YAML
  files for multiple purposes, leaving the original YAML
  untouched and usable as is.

This tool is sponsored by [sig-cli] ([KEP]), and
inspired by [DAM].

* Generate Kubernetes resources file
* Pure-yaml
* Template-free
* Kustomize "understands" kubernetes
* "Overlay"

## Links

* https://kustomize.io/
* https://github.com/kubernetes-sigs/kustomize
* https://github.com/kubernetes-sigs/kustomize/blob/master/examples
* https://kubectl.docs.kubernetes.io/pages/examples/kustomize.html

# Usage 

## Cli

### tldr

```shell script
tldr kustomize
```

### Create

```shell script
kustomize create --resources deployment.yaml,service.yaml
# OR 
kustomize create --autodetect
# Then
kustomize build . 
```

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - configMap.yaml
  - deployment.yaml
  - service.yaml
```

### Apply

```shell script
# Generate with kubectl 
kubectl apply -k . --dry-run -o yaml
# but
kustomize build . | kubectl apply -f - 
```

#### kubectl integration (BUT)

Since [v1.14][kubectl announcement] the kustomize build system has been included in kubectl.

| kubectl version | kustomize version |
|---------|--------|
| v1.16.0 | [v2.0.3](https://github.com/kubernetes-sigs/kustomize/tree/v2.0.3) |
| v1.15.x | [v2.0.3](https://github.com/kubernetes-sigs/kustomize/tree/v2.0.3) |
| v1.14.x | [v2.0.3](https://github.com/kubernetes-sigs/kustomize/tree/v2.0.3) |

## Features

### Prefix,Suffix
```shell script
kustomize edit set nameprefix acme-
```

```yaml
namePrefix: acme-
```

```shell script
diff -u <(kustomize build ../iso/base) <(kustomize build .)
```

### Annotation,Labels
```shell script
kustomize edit add annotation "application:front"
kustomize edit add label "env:dev"
```

```yaml
commonAnnotations:
  application: front
commonLabels:
  env: dev
```

```shell script
diff -u <(kustomize build ../iso/base) <(kustomize build .)
```

### ConfigmapGenerator
```shell script
kustomize edit add configmap my-configmap --from-literal=db_host=localhost
```

```yaml
configMapGenerator:
- literals:
  - db_host=localhost
  name: my-configmap
```

```shell script
kustomize build . | grep my-configmap 
# Then add a literal and run
kustomize build . | grep my-configmap 
```

### Image

```shell script
kustomize edit set image nginx:1.16
```

```yaml
images:
- name: nginx
  newTag: 1.16
```

```shell script
kustomize build . | grep nginx 
```

## Overlays and merge

```shell script
cd iso
diff -u <(kustomize build base) <(kustomize build overlays/staging)
diff -u <(kustomize build overlays/staging) <(kustomize build overlays/production)
```

## Remote targets

```shell script
cd remoteTargets
kustomize build .
```

# Integration

## Skaffold

```shell script
cd skaffold
skaffold dev
```

## Github Actions

```yaml
name: staging

on:
  push:
    paths:
      - 'staging/*'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v1
      - name: Get kubeconfig file from GKE
        uses: machine-learning-apps/gke-kubeconfig@master
        with:
          application_credentials: ${{ secrets.APPLICATION_CREDENTIALS }}
          project_id: ${{ secrets.PROJECT_ID }}
          location_zone: ${{ secrets.LOCATION_ZONE }}
          cluster_name: ${{ secrets.CLUSTER_NAME }}
      - name: Kubernetes toolset
        uses: stefanprodan/kube-tools@v1.0.0
        env:
          KUBECONFIG: '/github/workspace/.kube/config'
        with:
          kubectl: 1.16.2
          kustomize: 3.2.3
          command: |
            kustomize build staging | kubectl apply --dry-run -o yaml -f -
            kustomize build staging | kubectl apply -f -
```

## IntelliJ

[Plugin Jetbrains Kubernetes]



[sig-cli]: https://github.com/kubernetes/community/blob/master/sig-cli/README.md
[KEP]: https://github.com/kubernetes/enhancements/blob/master/keps/sig-cli/0008-kustomize.md
[DAM]: https://github.com/kubernetes-sigs/kustomize/blob/master/docs/glossary.md#declarative-application-management
[demoit]: https://github.com/dgageot/demoit
[Plugin Jetbrains Kubernetes]: https://plugins.jetbrains.com/plugin/10485-kubernetes/
