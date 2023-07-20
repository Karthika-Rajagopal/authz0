

Authz0 is an automated authorization test tool. Unauthorized access can be identified based on URLs and Roles & Credentials.

URLs and Roles are managed as YAML-based templates, which can be automatically created and added through authz0. You can also test based on multiple authentication headers and cookies with a template file created/generated once.

![authz0-2](https://user-images.githubusercontent.com/13212227/149650143-a34d8826-f272-4aca-b9a7-323de268cd52.jpg)

## 🛸 Key Features
* Generate scan template `$ authz0 new`
    * Include URLs
    * Include Roles
    * Include ZAP history (Select URLS > Save Selected Entiries as HAR)
    * Include Burp history (Select URLs > Save item)
    * Include HAR file
* Easy modify scan template (Role, URL) `$ authz0 setUrl` `$ authz0 setRole` `authz0 setCred`
* Scanning authorization(access-control) with template `$ authz0 scan`


