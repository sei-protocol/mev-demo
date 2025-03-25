# mev-demo
To build a plugin, run `make build`, which will produce `mev.so`.

To use a built `mev.so`, edit your `seid`'s `app.toml`. Replace `handler_plugin_path` under the `mev` section to your `mev.so`'s path, then restart `seid`.

# Important
In order for a Go plugin to work, the dependencies of the host program and the plugin should be the same. It is recommended that the plugin owner copy the entirety of `sei-chain`'s go.mod to their plugin's go.mod whenever `sei-chain` is upgraded, and rerun `go mod tidy && make build` for their plugin.