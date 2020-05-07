# Octopus

![Octopus](https://github.com/fallais/octopus/blob/master/assets/octopus_small.png)

**Octopus** is a proxy made with **Golang**.

## Usage

You can deploy it with a Docker container : `docker run -d -p 8080:8080 --name octopus fallais/octopus:latest`

## Configuration

Here is an example of configuration file.

```yaml
general:
  bind: :7001

security:
  allowed_ports:
    - 80
    - 443
    - 8080
  blacklist:
    - perdu.com
    - google.com
  whitelist:
    - perdu.com

cache:
  is_enabled: false
  type: local
  settings:
    path: /dir/dir/dir
```

## Credits

Icon made by Freepik from www.flaticon.com