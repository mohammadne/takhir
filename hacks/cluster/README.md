# Cluster

```sh
# create .vault_pass file and store your password there
touch .vault_pass

# create proxy_address for your proxy role (if included)
ansible-vault encrypt_string --vault-password-file .vault_pass 'http://address:port' --name 'proxy_address'

# run playbook against your host
ansible-playbook playbook.yaml -i inventory.ini --vault-password-file=.vault_pass
```
