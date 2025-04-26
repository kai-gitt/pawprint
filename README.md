## pawprint
a server for hosting markdown documentation


> [!WARNING]
> This project is work in progress, it should **not** be considered production ready

### Installation
- Get the latest release from [github releases](https://github.com/kai-gitt/pawprint) or build it yourself.
- Copy the example configuration from [./data/example_config.toml](https://github.com/kai-gitt/pawprint/blob/main/data/example_config.toml) or create one yourself.
```toml
[server]
# Network stack configuration
host = "0.0.0.0"
port = 8000
# You can generate your key by
# running something like `head -n 5 /dev/urandom | sha256sum`
key = "NOT_SECURE!_CHANGE_ME_TO_SOMETHING_SECURE_01234"
# If the server is running behind a proxy, like cloudflare or your own.
behind-proxy = false
# SSL configuration (unused if behind-proxy is set to true)
ssl = { key = "./data/ssl.pem", cert = "./data/ssl.cert" }
domain = "docs.kai.enterprises"

# Directories are relative to the server executable,
# you can set a directory to anything you want,
# as long as the running user has permissions to it.
[directories]
# Directory where all the documentation is
docs = "./docs"
# Directory for logs
logs = "./logs"
```
- Run the server executable, may require running `chmod +x pawprint` beforehand, to make the file executable.
- (optional) If you already have a running http server (like [apache](https://httpd.apache.org/docs/2.4/howto/reverse_proxy.html) or [nginx](https://docs.nginx.com/nginx/admin-guide/web-server/reverse-proxy/)), you should reverse-proxy pawprint with them.


<a name="top">Back to top</a>