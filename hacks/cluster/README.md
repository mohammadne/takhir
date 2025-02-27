# Cluster

## Setup

```sh
# create .vault_pass file and store your password there
echo 'strong-password' > .vault_pass

# create server_ip_address for your kind variables
ansible-vault encrypt_string --vault-password-file .vault_pass '195.248.243.80' --name 'server_ip_address'

# run playbook against your host
ansible-playbook playbook.yaml -i inventory.ini --vault-password-file=.vault_pass
```

## Usage

```sh
export KUBECONFIG=~/.kube/kind-zanbil.conf
kubectl get nodes
```
