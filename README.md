### 🗃️ Archived Repository


**Development has moved to the official Portainer Terraform Provider:**

## 📦 Terraform Registry: [portainer/portainer](https://registry.terraform.io/providers/portainer/portainer/latest)
## 💻 GitHub: [github.com/portainer/terraform-provider-portainer](https://github.com/portainer/terraform-provider-portainer)

**Please update your configurations to use the official provider to:**
```hcl
terraform {
  required_providers {
    portainer = {
      source = "portainer/portainer"
    }
  }
}
```
