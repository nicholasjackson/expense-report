variable "vault_k8s_cluster" {
  default = "dc1"
}

k8s_cluster "dc1" {
  driver = "k3s"

  nodes = 1

  network {
    name = "network.dc1"
  }
}

network "dc1" {
  subnet = "10.5.0.0/16"
}

module "vault" {
  source = "github.com/shipyard-run/blueprints//modules/kubernetes-vault?ref=800a8963ce10f341ae09b44d0412fae070057214"
}

output "KUBECONFIG" {
  value = k8s_config("dc1")
}

output "KUBE_CONFIG_PATH" {
  value = k8s_config("dc1")
}