[Flux](https://github.com/fluxcd/flux2) is the GitOps Kubernetes operator.

Flux can deploy your app into k8s automatically or manually. Please [go through](https://fluxcd.io/docs/get-started/) the `Get started` before try the following demo.

## Prerequisite

* A Kubernetes cluster
  * You can install a k8s via [kind](https://github.com/kubernetes-sigs/kind), [k3d](https://github.com/rancher/k3d)
  * Or you can install k8s via [kubekey](https://github.com/kubesphere/kubekey)
* flux CLI
  * Install it via [hd](https://github.com/LinuxSuRen/http-downloader): `hd install flux2`

## Get started

Do the prerequistes check: `flux check --pre`

Run the bootstrap command:
```
flux bootstrap github \
  --owner=$GITHUB_USER \
  --repository=fleet-infra \
  --branch=main \
  --path=./clusters/my-cluster \
  --personal
```

Clone the git repository:
```
git clone https://github.com/$GITHUB_USER/fleet-infra
cd fleet-infra
```

Create Source:
```
flux create source git go-demo \
    --url https://github.com/devops-ws/learn-pipeline-go \
    --branch=kustomize \
    --interval=30s \
    --export > clusters/my-cluster/go-demo-source.yaml
```

Create a kustomization:
```
flux create kustomization go-demo \
    --source=go-demo \
    --path="./kustomize" \
    --prune=true \
    --validation=client \
    --interval=5m \
    --export > clusters/my-cluster/go-demo-kustomization.yaml
```

Save it to your git repository:
```
git add -A && git commit -m "Add go-demo Source and Kustomization"
git push
```
