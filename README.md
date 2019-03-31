# k8s learning

## testimg

testimg is a go web server that hosts a hello world image. Its purpose is to become a dockerfile, running the go binary exposed on port 8080, and deploy to kubernetes. 

## manifests

kubernetes manifests for deployment, and service.

### test-deployment.yaml

deployment and service for the testimg container.

### nginx-deployment.yaml

deployment and service for nginx. 


## Makefile

default action for make is to compile the program, build a docker image and add the resultant binary. Next it updates the test-deployment manifest using ```kubectl apply``` then terminates the existing pods. 
