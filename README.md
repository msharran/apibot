# ApiBot

A Postman for terminal users
---

# Installation (MacOS)

1. Download the binary from latest release.
2. Move the binary to `/usr/local/bin/apibot`

---
# Usage

## Initialize ApiBot with base url
```bash
apibot init "https://jsonplaceholder.typicode.com"
```
  For adding new profile pass `-p` or `--profile`. By default a profile named `default` is created for you.
```bash
apibot init "https://jsonplaceholder.typicode.com" --profile=dev
```
 *Note: The config file is created in `~/.apibot/config.yaml`*

 ## Simple GET Request
 
 Examples
 ```bash
 apibot get /comments/1
 ```

 ```bash
 apibot get /comments/1 -p=dev
 ```
 
 ```bash
 apibot get "/users?limit=10"
 ```

 ## Sending complex request
 
 1. Add your request in `config.yaml` under requests section with a unique request name. 
 
 ```yaml
 default:
  base_url: https://jsonplaceholder.typicode.com
  headers:
    Authorization: Bearer token
    Custom: Header
  requests:
    create_post:
      Endpoint: /posts
      Method: POST
      Body: >
        {
          title: 'foo',
          body: 'bar',
          userId: 1,
          id: 1
        }
 ```
 
 2. Run the command using the request name
  ```bash 
  apibot send create_post
  ```
 
 

