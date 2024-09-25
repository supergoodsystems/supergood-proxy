# supergood-proxy

## Local Development

`make run-local`

The above make target with run a the supergood-proxy locally. The environment variables are found in _config/dev.yml. 

The supergood proxy also depends on ADMIN_CLIENT_KEY which can be supplied as either an environment variable, or directly to the _config/dev.yml in the remoteWorkerConfig struct. The admin client key is used to authenticate to the remote config worker.

As an example config:
```yml
remoteWorkerConfig:
  baseURL: "http://localhost:3001"
  fetchInterval: "60s"
  adminClientKey: "supergood-admin-key"
proxyConfig:
  port: "8080"
  healthCheckPort: "8081"
```

By default, the supergood-proxy will pull the remote config from a local instance of the supergood server. To override, update the baseURL in the above config. To point to production, set `baseURL: "https://api.supergood.ai"`

## Production Deployment

A dockerfile is provided in this repository.

The `ENV` environment variable will help determine which _config/* file to load.
