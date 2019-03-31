

.PHONY: default
default: testimg manifests taint

.PHONY: testimg
testimg:
	$(MAKE) -C testimg

.PHONY: manifests
manifests:
	kubectl apply -f manifests/test-deployment.yaml

.PHONY: taint
taint:
	for pod in `kubectl get pods | grep testdep | awk '{print $$1}'`; do kubectl delete pods $$pod; done;
