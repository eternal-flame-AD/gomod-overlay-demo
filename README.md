# gomod-overlay-demo

## Description

This repo consisted of two seperate modules: `github.com/eternal-flame-AD/gomod-overlay-demo/server` and `github.com/eternal-flame-AD/gomod-overlay-demo/plugin` The plugin is designed to be built as a go plugin to be loaded by the server.

The server required a lower version of `github.com/gin-gonic/gin@1.1.4` while the plugin required a high version `github.com/gin-gonic/gin@v1.3.0`

## The Problem

As go mod follows a [Minimal Version Selection](https://github.com/golang/proposal/blob/master/design/24301-versioned-go.md), when the server requires a lower version of `gin` while the plugin requires a higher version. They will be built with different dependencies. Thus, a panic occurs when the plugin is loaded.

Since there is no way to tell go which main package the plugin package is built for, a workaround to this problem is to manually copy all common dependency requirements from the server `go.mod`, however it requires a lot of human intervention and a non-synchronized `go.mod` will cause subsequent builds to fail.

## The Proposal

I think we can add a `comply` directive to `go.mod`. The directive will tell go module system the package dependencies need to be in comply with another main package, as it will be built as a plugin for that package.
```
module github.com/eternal-flame-AD/gomod-overlay-demo/plugin

comply (
    github.com/eternal-flame-AD/gomod-overlay-demo/server@v1.0.0
)
```

The `comply` directive poses these effects:
- When executing `go mod tidy`, all applicable `require`, `exclude` and `replace` directives from given version of the foreign module are applied into the current `go.mod` (In this example, when executing `go mod tidy` on the plugin module, the version requirement on `github.com/gin-gonic/gin` will be changed to in comply with the server(`v1.1.4`))
- When downloading modules, abort with error if the calculated dependencies does not comply with the packages indicated in the `comply` directive