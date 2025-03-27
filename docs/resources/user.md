# 👤 **Resource Documentation: `portainer_user`**

# portainer_user
The `portainer_user` resource allows you to manage user accounts in Portainer.

## Example Usage

### Internal User

```hcl
resource "portainer_user" "your-user" {
  username = "youruser"
  password = "strong-password"
  role     = 1 # 1 = admin, 2 = standard user
}
```

### LDAP User
```hcl
resource "portainer_user" "your-user" {
  username  = "youruser"
  role      = 2
  ldap_user = true
}
```
## Lifecycle & Behavior

Users are recreated if any of the attributes change (e.g., username, password, role, ldap_user).

- To delete a user created via Terraform, simply run:
```hcl
terraform destroy
```

- To change a user's role, update the role field and re-apply:
```hcl
terraform apply
```

## Arguments Reference

| Name        | Type    | Required                  | Description                                                                 |
|-------------|---------|---------------------------|-----------------------------------------------------------------------------|
| `username`  | string  | ✅ yes                    | Username of the user.                                                       |
| `password`  | string  | ✅ yes                    | Required for non-LDAP users. Not allowed when `ldap_user = true`.           |
| `role`      | integer | 🚫 optional (default `2`) | Role of the user. `1` = admin, `2` = standard user.                         |
| `ldap_user` | bool    | 🚫 optional (default `false`) | Set to `true` if the user is managed by LDAP.                         |
| `team_id`   | integer | 🚫 optional               | Optional Portainer team ID. Can only be used when `role = 2` (standard user). |

> ⚠️ If `ldap_user = true`, the `password` must be omitted.  
> ⚠️ `team_id` is ignored for admin users (`role = 1`).

## Attributes Reference

| Name | Description              |
|------|--------------------------|
| `id` | ID of the Portainer user |
