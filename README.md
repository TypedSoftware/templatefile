# `templatefile`

This is a standalone CLI to [the Terraform `templatefile` function][1].

> `templatefile` reads the file at the given path and renders its content as a template using a supplied set of template variables.
>
> […]
>
> The "vars" argument must be a map.

[Read the function docs on terraform.io →][1]

### Installation

[Download a binary from the releases section on GitHub →][2]

## Usage

```bash
printf '%s\n' 'Hello, ${name}!' > hello.tmpl
printf '%s\n' 'name: world' > hello.yml
templatefile hello.tmpl hello.yml
# Hello, world!
```

Both YAML and JSON files work. The root of the file must be an object.

## License

This code is based on the implementation in Terraform and [is available under the same license →](./LICENSE.txt)

  [1]:https://www.terraform.io/docs/configuration/functions/templatefile.html
  [2]:https://github.com/TypedSoftware/templatefile/releases
