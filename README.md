# Docker Auth Profiler

The goal of this project if to find out which **docker endpoints are being filtered by an authorization plugin**.

This code will **brute force all the possible docker api endpoints** and show which ones are **allowed**.

```bash
Usage: ./docker_auth_profiler /run/docker.sock
  -a	Print allowed and forbidden
  -e string
    	Indicate the error message fingerprint (default "authorization denied by plugin")
  -h	Print help
  -v string
    	Version of the docker API (default "v1.41")
```

## How I detect a docker authorization plugin
If you perform a regular docker action and find an error message such as this one: `docker: Error response from daemon: authorization denied by plugin authobot:latest: use of Privileged containers is not allowed.`. Then, there is a docker authorization plugin.

You could also **try to list the plugins** with:
```bash
docker plugin list

ID             NAME              DESCRIPTION                       ENABLED
ca99dd5a4e26   authobot:latest   Authorization plugin for Docker   true
```

## Update docker endpoints
The script `download_endpoints.py` will download and parse all the endpoints from https://docs.docker.com/engine/api/v1.40.yaml.

If you want to update the endpoints, start by **executing that script** then grab the generated json and update the `docker_auth_profiler.go` file.


## More Info
For more information about **what is** a docker authorization plugin and possible **bypasses** check **https://book.hacktricks.xyz/linux-unix/privilege-escalation/docker-breakout/authz-and-authn-docker-access-authorization-plugin**

For an example of a **simple and vulnerable docker auth plugin check https://github.com/carlospolop-forks/authobot**
