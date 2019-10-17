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
kustomize create --resources deployment.yaml,service.yaml
kustomize create --autodetect
kustomize build . 
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

### Edit
```shell script
kustomize edit set nameprefix 
```

## Features

### Prefix

```yaml

```

### Annotation

### Labels

### ConfigmapGenerator

### Image

 

```

```

prefix
annotation
labels

configmapgenerator (or secretgenerator)
image
overlays


real problems 

cloud build 