## Example plugin for k3d

Go project to test the new k3d plugin system.

### Specification
Plugin specs are still in development, see [Discussion #741](https://github.com/rancher/k3d/discussions/741).

At the moment, a plugin is an executable file (either binary or a script) with a corresponding manifest that describes the content.
An example of manifest file is:
```yaml
# The name of the plugin
# This is used when listing plugins, executing them via a lifecyle hook or directly from CLI
name: pluginName
# Specify a version for the plugin
version: v0.0.0
# A short description for the plugin
description_short: This plugin does something
# A detailed description
description: |-
    This plugin is useful for doing something in a k3d cluster
```
