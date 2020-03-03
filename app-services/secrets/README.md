# Secrets example

#### Overview

This example demonstrates storing a secret to the secret store (Vault) and retrieving those secrets.

When running an application service in secure mode, secrets can be stored by making an HTTP `POST` call to `http://[host]:[port]/api/v1/secrets`.  If running in insecure mode, secrets can be configured in consul or in the config yaml file.

Application Services can retrieve their secrets from the underlying secret store using the GetSecrets() API in the SDK. 

`.GetSecrets(path string, keys ...string)` is used to retrieve secrets from the secret store. `path` describes the type or location of the secrets to retrieve. If specified it is appended to the base path from the secret store configuration. `keys` specifies the secrets which to retrieve. If no keys are provided then all the keys associated with the specified path will be returned.

#### Secure mode

If in secure mode, the secrets are stored and retrieved from Vault based on the SecretStore configuration values.

Setup the application config to use  Vault's secrets engine:

1. Run vault using the docker compose file in this directory. 

   ` docker-compose -f docker-compose.yml up`

2. If running this example natively (not in docker), then you will need to manually set the current Vault token in configuration.

   1. Find the root token for Vault in the logs of the container:

      `docker logs <container-id>

   ![image-20200224130525112](./root-token.png)

   2. Copy the token to the *'root_token'* field in the token file (*token.json*). 
   3. Change the *SecretStore.Tokenfile* section of the configuration toml to point to the token file.
   4. Copy the token to the *SecretStore.Authentication.AuthToken* section of the configuration.

3. Use docker exec to run commands on the running vault container instance.

   `docker exec -it <container_id> /bin/sh`

4. Enable the secrets engine using the Vault command line

   *Login into vault with the root token.*

```
/ # vault login
/ # vault secrets disable secret
Success! Disabled the secrets engine (if it existed) at: secret/
/ # vault secrets enable -version=1 -path=secret kv
Success! Enabled the kv secrets engine at: secret/
```

#### Insecure mode

When security is disabled, the secrets can be written to and are retrieved from the *Writable.InsecureSecrets* section of the configuration file. Insecure secrets and their paths can be configured as below.

```toml
 [Writable.InsecureSecrets]    
      [Writable.InsecureSecrets.NoPath]
        Path = ''
        [Writable.InsecureSecrets.NoPath.Secrets]
          username = 'nopath-user'
          password = 'nopath-pw'
      
      [Writable.InsecureSecrets.AWS]
        Path = 'aws'
        [Writable.InsecureSecrets.AWS.Secrets]
          username = 'aws-user'
          password = 'aws-pw'
      
      [Writable.InsecureSecrets.MongoDB]
        Path = 'mongodb'
        [Writable.InsecureSecrets.MongoDB.Secrets]
          username = ''
          password = ''
```

`NOTE: An empty path is a valid configuration for a secret's location  `

#### Run StoreSecrets and GetSecrets

##### StoreSecrets

When running in secure mode, secrets for can be stored in the secret store (Vault) by making an HTTP `POST` call to `http://[host]:[port]/api/v1/secrets`.  

*SecretsExample.postman_collection.json* contains Postman requests to store secrets to using the */secrets* API route. 

If in secure mode, execute the Store Secrets Postman requests to push secrets to Vault.

If running in insecure mode, configure the secrets in consul or in the config yaml file (as described [above](#Insecure mode)).

##### GetSecrets

*SecretsExample.postman_collection.json* contains a Postman request to get secrets by triggering a CoreData event to the application service. 

Execute the *CoreData Event Trigger (Random-Float-Device)* request and view the secrets in the application's console.

This trigger causes execution of the pipeline function that uses the SecretClient to get secrets.