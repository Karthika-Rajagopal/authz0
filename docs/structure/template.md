---
title: Template
layout: default
parent: Structure
nav_order: 1
---

## Template (YAML)

```yaml
name: template name
roles: # roles of this template
- name: role-name1
- name: role-name2
- name: role-name3
urls:
- url: https://authz0.Karthika-Rajagopal.com/your-url
  method: GET # Method
  contentType: "" # or json
  body: "" # HTTP Request Body
  allowRole: # Allowed role
  - role-name1  
  - role-name2
  denyRole: [] # Denied role
  alias: "main" # Alias of this URL
asserts: # assertions
- type: success-status
  value: "200,201,202,204"
credentials:
- rolename: User
  headers:
  - 'X-API-Key: 1234'
- rolename: Admin1
  headers:
  - 'X-API-Key: 5555'
```