# Factorio Mod Installer

This program will download or update mods you specify in the config file to your headless Factorio install. It is just
a small binary, so just extract one of the releases into your $PATH, or anywhere, and run it.

## Run it

Install the Factorio headless server, by default this assumes `/opt/factorio` as the install path. Then create
a config file with the mods you want in it, it should look like the example below:

```yaml
username: yourFactorioNameHere
token: yourTokenHere

mods:
  - long-reach
  - Bottleneck
```

By default it looks for the config file in `/opt/factorio/mod-install.yaml` this can be overridden with `--config`

You can also specify `factorioPath` and give the absolute path to your install path if it is different 
than `/opt/factorio`.

## TODO

- Check for unwatched mods, and give the option to update them. Maybe `--unwatched`?
- Remove mods that aren't in the list. Something like `--remove`
- Support Windows

