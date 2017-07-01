# Logzio hook for Logrus <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:" />
Send Logrus logs to Logzio

## Getting Started

### Get Logzio token
1. Go to Logzio website
2. Sign in with your Logzio account
3. Click the top menu gear icon (Account)
4. The Logzio token is given in the account page

### Initialize Logzio hook
```go
package yourpackagename

import (
        "github.com/sirupsen/logrus"
        "github.com/bshuster-repo/logruzio"
)

const LOGZIO_TOKEN = "fjdhslGJHSDHG23edg"

func main() {
        ctx := logrus.Fields{
                "ID": "12adebacd8",
                "Version": "1.0.0-dev",
        }
        hook, err := logruzio.New(LOGZIO_TOKEN, "YourAppName", ctx)
        if err != nil {
                logrus.Fatal(err)
        }
        logrus.AddHook(hook)
        logrus.Info("Lets go!")
}
```

**NOTE**: Set `LOGZIO_TOKEN` to the Logzio token as mentioned in `Get Logzio token`.
