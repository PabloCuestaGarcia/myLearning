# RBAC

Roles: permite definir políticas a un namespace concreto.
clusterRole: permite definir políticas a un cluster

roleBinding: Asigna un rol a un usuario o service account.
clusterRoleBinding: asigna un clusterRol a un usuario o serviceAccount

serviceAccount: en lugar de usuarios, podemos asignar esos permisos a Pods.

## Generating keys

```shell
openssl req -new -key developer.key -out developer.csr -subj "/C
```




