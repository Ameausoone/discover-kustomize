# Kustomize 

## What is it ?

## Links

* https://github.com/kubernetes-sigs/kustomize/blob/master/examples
* https://kubectl.docs.kubernetes.io/pages/examples/kustomize.html

## Cli

``` shell script
tldr kustomize
```

# Usage 

## Cli

### Create
```shell script
# download resources
cd workdir
BASE=$(pwd)

CONTENT="https://raw.githubusercontent.com/Ameausoone/discover-kustomize/master/workbase"

curl -s -o "$BASE/#1" "$CONTENT/base\
/{deployment.yaml,configMap.yaml,service.yaml}"

kustomize create --resources deployment.yaml,service.yaml
kustomize create --autodetect
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
kubectl apply -k
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
diff -u <(kustomize build ../iso) <(kustomize build .)
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
diff -u <(kustomize build ../iso) <(kustomize build .)
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
# Then add a literal and
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

## Remote resources

# Integration

## Skaffold

## Cloud build 
