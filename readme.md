# Parameters
## PROXY_PASS (necessary)
Set the reverse url, for example, if you want visit https://xxx.com/api/xxx1 via http://localhost:8080/api/xxx1, you can just set this parameter
## PROXY_URL
If the proxy_pass url via a proxy, you need to set this parameter
## LISTEN
Set the address for the server

| Paramter Name | Set with command line | Set with environemnt variable | Default |
| ------------- | --------------------- | ----------------------------- | ------- |
| PROXY_PASS | --pass | STK_REVERSE_PASS | - |
| PROXY_URL | --proxy | STK_REVERSE_PROXY |  |
| LISTEN | --listen | STK_REVERSE_LISTEN | :8080 |

# Example
use proxy `http://127.0.0.1:7890` reverse to `https://www.google.com` listen on `:5000`
```bash
.\stk-proxy.exe --proxy http://127.0.0.1:7890 --pass https://www.google.com --listen :5000
```