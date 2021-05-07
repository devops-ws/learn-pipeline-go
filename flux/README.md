[Flux](https://github.com/fluxcd/flux) is the GitOps Kubernetes operator.

Flux can deploy your app into k8s automatically or manually. Please [go through](https://docs.fluxcd.io/en/latest/get-started/) the `Get started` before try the following demo.


## Get started

Firstly, you need to deploy [deploy.yaml](deploy.yaml) into your cluster.

Then, expose the port via: `kubectl port-forward deploy/go-demo 9898:80`

Test it via: `curl localhost:9898`