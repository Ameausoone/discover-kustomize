# Kustomize

---

## /me

Antoine Méausoone

---

## Kustomize

> `kustomize` lets you customize raw, template-free YAML
files for multiple purposes, leaving the original YAML
untouched and usable as is.

This tool is sponsored by [sig-cli]([KEP]), and
inspired by DAM.

---

* Generate Kubernetes resources file
* Pure-yaml
* Template-free
* Kustomize "understands" kubernetes
* "Overlay"

---

## Links

* https://kustomize.io/
* https://github.com/kubernetes-sigs/kustomize
* https://github.com/kubernetes-sigs/kustomize/blob/master/examples
* https://kubectl.docs.kubernetes.io/pages/examples/kustomize.html

---

## tldr

```shell script
tldr kustomize
```

---

## Cli

---

### Create

```shell script
# download resources
cd workdir

BASE=$(pwd) && \
CONTENT="https://raw.githubusercontent.com/Ameausoone/discover-kustomize/master/workbase" && \
curl -s -o "$BASE/#1" "$CONTENT/base/{deployment.yaml,configMap.yaml,service.yaml}" 

tail -n +1 *
``` 

---

```shell script
kustomize create --resources deployment.yaml,service.yaml
kustomize build .
kustomize create --autodetect
```

---

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - configMap.yaml
  - deployment.yaml
  - service.yaml
```

---

#### kubectl integration (BUT)

Since v1.14 the kustomize build system has been included in kubectl.

| kubectl version | kustomize version |
|---------|--------|
| v1.16.0 | [v2.0.3](https://github.com/kubernetes-sigs/kustomize/tree/v2.0.3) |
| v1.15.x | [v2.0.3](https://github.com/kubernetes-sigs/kustomize/tree/v2.0.3) |
| v1.14.x | [v2.0.3](https://github.com/kubernetes-sigs/kustomize/tree/v2.0.3) |

---

### Apply

```shell script
# Generate with kubectl 
kubectl apply -k . --dry-run -o yaml

# BUT
kustomize build . | kubectl apply -f - 
```

---

## Features

---

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

---

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

---

### ConfigmapGenerator
```shell script
kustomize edit add configmap my-configmap \ 
--from-literal=db_host=localhost
```

```yaml
#kustomization.yaml
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

---

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

---

## Overlays and merge

```shell script
cd iso
diff -u <(kustomize build base) <(kustomize build overlays/staging)
diff -u <(kustomize build overlays/staging) <(kustomize build overlays/production)
```
---

## Remote targets

```shell script
cd remoteTargets
kustomize build .
```

---

# Integration

---

## Skaffold

```shell script
cd skaffold
skaffold dev
```