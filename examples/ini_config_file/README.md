
```bash
> $ go run main.go -h

Usage
  -domain string
		
  -mongodb_host string
		 (default "example.com")
  -mongodb_port int
		 (default 999)
  -user_name string
		
  -user_passwd string
		
Environment variables:
 $EXAMPLE_DOMAIN string

 $EXAMPLE_USER_NAME string

 $EXAMPLE_USER_PASSWD string

  $EXAMPLE_MONGODB_HOST string
	 (default "example.com")
  $EXAMPLE_MONGODB_PORT int
	 (default "999")

Config file "./config.ini":

```ini
debug = true
Domain = example.com

[user]
name = ""
passwd = ""

[mongodb]
host = myhost
port = 909
```


exit status 2
```
