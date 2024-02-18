# config

This package is a way to define a standard configuration format when building tools that need to interact with Lagoon.

It uses [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) for storing the configuration in known locations.

The main configuration file will be stored in `$XDG_CONFIG_HOME` with the file location being in a `lagoon` directory, with the actual configuration file being named `config.yaml`.

The full path will then be `$XDG_CONFIG_HOME/lagoon/config.yml`. 

If `$XDG_CONFIG_HOME` is not set, then depending on the operating system, this location could be different. Review the specification to understand more about this.

# Components of Config

## Contexts

Contexts are where you can define one or more Lagoon targets. This context will contain information related to how clients can talk to the API or other components of Lagoon.

Contexts contain the following information
* name - this is a name to make it obvious to the user which context they are using, it does not have to match anything in Lagoon.
* api hostname - this is the hostname for the api, without the `/graphql` or any other paths, as clients should set this as required. (eg: https://api.lagoon.sh)
* token hostname - this is the hostname for the ssh based token endpoint. (eg: token.lagoon.sh)
* token port - this is the port for the ssh based token endpoint. (eg: 22)
* authentication hostname - this is the hostname for the Lagoon authentication endpoint. (eg: https://keycloak.lagoon.sh)

## Users

Users are where you can define one or more users. Users are assigned to a context, and contain information related to how clients will authenticate with a context.
Users are linked to a context, and they can be changed between contexts too. This way its possible to have multiple accounts to use without needing to define multiple contexts. You just update the context with a different user as needed.

Users contain the following information
* name - this is a name used only within the scope of the config. It does not have to be the actual users name
* ssh key - this is the ssh private key that will be used to authenticate with the ssh token endpoint to retreive a token. The public key should be attached to a user in Lagoon
* grant - this is where the users token information is stored. This will be updated automatically by clients when they interact with the configuration.

## Default Context

The default context is the context that clients will use when using the config package. Users can change their default context, but clients should also provide a way for the context to be provided in other methods. This package offers `GetDefaultContext` and `GetContext(name)` functions which return the same resource, clients can use these with possibly a flag to determine which is called.

## Features

Features are a way to enable or disable functionality using a `map[string]bool`. Features will default to `false` if requested but do not exist.
Features are global and context configurable, features on a context should override a global feature.

Config will not enforce any specific feature names. They are up to the consumer of the config to decide on what the feature will be called.
The recommendation though will be to ensure your tool adds a somewhat unique prefix when you're setting and retrieving features to prevent clashes.
* `cli` prefix for features in the Lagoon CLI
* `sync` prefix for features in Lagoon sync
* other prefix for custom tools.

Additionally, it is recommended that you avoid using features from other tools, ie, don't consume `cli` prefixed features in `sync`.

# Usage

To use config, you first need to create a `Config`, you can use the `LoadConfig` function to create the config. `LoadConfig` accepts a boolean parameter which can create the configuration directory structure and file, by default it will not create the configuration path and will return an error if the config does not exist.

```
createConfig := true
c, err := LoadConfig(createConfig)
if err != nil {
    panic(err)
}
```

To create a context, a user is required. Create a new user, and give this user a name. Depending on the tool that is using config, the `SSHKey` may or may not be required.

```
uConfig := UserConfig{
    Name:  "my-user",
    SSHKey: "/full/path/to/private_key",
}
err = c.NewUser(uConfig)
if err != nil {
    panic(err)
}
```

Now that a user exists, you need to create a new context. You have to populate all the fields or `config` will return an error.
Assign the user to this new context.

```
cUser := "my-user"
context := "my-lagoon"
cConfig := ContextConfig{
    Name: context,
	APIHostname: "https://api.lagoon.sh"
	TokenHost: "token.lagoon.sh"
	TokenPort: 22
	AuthenticationEndpoint: "https://keycloak.lagoon.sh"
}
err = c.NewContext(cConfig, cUser)
if err != nil {
    panic(err)
}
// then set the new context as the default
err = c.SetDefaultContext(cConfig.Name)
if err != nil {
    panic(err)
}
```

Optionally, you can leverage the `features` of Config for enabling or disabling functionality within your tool.
```
toolPrefix := "mytool" // this can be a const within your tool, hidden from the user.
feature := "my-feature"
state := true

// either globally set a feature
c.SetGlobalFeature(toolPrefix, feature, state)
// or as a specific context override
c.SetContextFeature(context, toolPrefix, feature, state)
```

After any changes made to config, contexts, users, or features, make sure to write the changes to the config file using `WriteConfig`.

```
c.WriteConfig()
```

Once you've got the user and context created, you can return the context for use. 
The context has attributes that clients can use to interact with Lagoon.

```
// source the default context for the config
lContext, err := c.GetDefaultContext()
if err != nil {
    panic(err)
}
// OR get a context by name
cName := "my-lagoon"
lContext, err := c.GetContext(cName)
if err != nil {
    panic(err)
}
```

Finally, you can get the user for the context (or any user by name). 
The user has attributes that clients can use to interact with Lagoon.

```
// get the users details for the context
lUser, err := c.GetUser(lContext.User)
if err != nil {
    panic(err)
}
```

To get the status of a feature, you can recall it using `GetFeature` and passing the required parameters. `GetFeature` will return an error if the context does not exist, otherwise if the feature is not found it will return `false`.
```
// retrieving a feature for use in your tool, must provide a context even if global.
featureState, err := c.GetFeature(context, toolPrefix, feature)
if err != nil {
    panic(err)
}
```

# Notes
Some clients may handle refreshing tokens themselves, if this is the case, once a token is refreshed you need to make sure to `WriteConfig`.
This `config` package does not handle token refresh at this point in time, but it may in the future.
