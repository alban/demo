
cd /home/alban/go/src/github.com/coreos/coreos-kubernetes/single-node

vagrant up --provider virtualbox

$ export KUBECONFIG="${KUBECONFIG}:$(pwd)/kubeconfig"
export KUBECONFIG=/home/alban/go/src/github.com/coreos/coreos-kubernetes/single-node/kubeconfig
$ kubectl config use-context vagrant-single


-------------------

cd /home/alban/go/src/github.com/microservices-demo/microservices-demo/deploy/kubernetes

-------------------

cd /home/alban/git/kinvolk/kinvolk-demo/linuxcon-toronto
kubectl create -f scope.yaml
kubectl create -f wholeWeaveDemo.yaml

-------------------

sudo wget -O /var/tmp/scope https://git.io/scope
sudo chmod a+x /var/tmp/scope
sudo /var/tmp/scope launch --no-app
sudo /var/tmp/scope launch

sudo docker run albanc/scope:krnowak-tcd

-------------------

http://localhost:9101/
http://localhost:9102/
http://localhost:4040/

~/.ssh/config
  LocalForward 9101 10.3.0.31:80
  LocalForward 9102 10.3.0.32:80
  LocalForward 4040 10.3.0.61:4040

