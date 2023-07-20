---
title: New command
layout: default
parent: Usage
nav_order: 1
---

```
Usage:
  authz0 new <filename> [flags]

Flags:
      --assert-fail-regex string       Set fail regex assert
      --assert-fail-size ints          Set fail size assert (support duplicate flag)
      --assert-fail-size-margin int    Set approximation range of fail size assert
      --assert-fail-status ints        Set fail status assert (support duplicate flag)
      --assert-success-status string   Set success status assert
  -h, --help                           help for new
      --include-burp string            Include Burp log file (XML)
      --include-har string             Include HAR file
      --include-roles string           Include Roles from the file
      --include-urls string            Include URLs from the file
      --include-zap string             Include ZAP log file (HAR)
  -n, --name string                    Template name
```

```
authz0 new admin.yaml -n test-admin --include-urls ./urls.txt --assert-fail-regex "permission denied"
```