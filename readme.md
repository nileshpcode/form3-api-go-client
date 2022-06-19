## form3-api-go-client

#### Usage:
```
import "github.com/form3"

client := form3.NewClient(baseURL) // baseURL is the root URL for all invocations of the client to form3 API

// create an account
client.AccountInterface.Create()

// fetch an account
client.AccountInterface.Fetch()

// delete an account
client.AccountInterface.Delete()

```