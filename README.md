# fcfc: Factory of CF CLI

`fcfc` is alias factory for `cf`.

## Install

Download binary from [releases page](https://github.com/autopp/fcfc/releases) and place it in `$PATH`.

## Usage

Write settings YAML and place into `~/.fcfc.yml`.

E.g.
```yaml
commands:
  - name: devcf
    api: api.run.pivotal.io
    org: myorg
    space: myspace
    login-options: "-u autopp@example.com"
  - name: othercf
    api: pcf.example.com
    org: otherorg
    space: otherspace
    login-options: "-u autopp"
```

And evaluate output of `fcfc` in `.bashrc`/`.zshrc`.
```sh
eval "$(fcfc)"
```

You can also place settings into other path. If you do, please pass it as a parameter.
```sh
eval "$(fcfc /path/to/.fcfc.yml)"
```

Now, you can use `login-devcf`/`login-othercf` to login your target and `devcf`/`othercf` to execute command in specified target.
```
$ login-devcf
$ devcf push
```

## License

[Apache License 2.0](LICENSE)

## Author

[@AuToPP](https://twitter.com/AuToPP)
