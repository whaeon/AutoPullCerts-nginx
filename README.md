# AutoPullCerts-nginx
Automatically pull certificates from dnspod

## Preparation
- Change the ownership of the default letsencrypt directory, such as `sudo chown -R "$USER":wheel /etc/letsencrypt` 

## TODO
- if the domain was not update, the program need to determine which to upload.