# Go testcontainers

- Golang
- PostgreSQL
- Testcontainers 
- GitLab (docker executor)

## GitLab Runner config.toml

```toml
concurrent = 1
check_interval = 0

[session_server]
  session_timeout = 1800

[[runners]]
  name = "###"
  url = "###"
  id = #
  token = "###"
  token_obtained_at = ###
  token_expires_at = ###
  executor = "docker"
  [runners.custom_build_dir]
  [runners.cache]
    [runners.cache.s3]
    [runners.cache.gcs]
    [runners.cache.azure]
  [runners.docker]
    tls_verify = false
    image = "docker:latest"
    privileged = true
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    disable_cache = false
    volumes = ["/cache"]
    shm_size = 0
```