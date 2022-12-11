# GoSca

Console App that allows to create projects scaffolding
by a given configuration

# Scaffolding Configuration

This configuration can be represented as json or yaml.
The members of the schema are described bellow.

## Dir Object

The **Dir Object** is the fundamental object, its composition
is the following:

```
Object Dir {
    String Name
    Array<String> Files? 
    Array<Dir> Dirs?
}
```

### Name

It's the current name of the directory. Its type is Literal String.

### Files

A String Array that represent the files inside the current directory.
*This attribute is optional*

### Dirs

A Dir Object Array that represent the subdirectories of the current directory.
*This attribute is optional*

## Scaffolding Object

This object owns a Dir Object as **Scaffolding** attribute.

````
Object Scaffolding {
    Dir Scaffolding {
        String Name
        Array<String> Files? 
        Array<Dir> Dirs?
    }
}
````

In this case, the **Name** attribute represents the actual name
of the project.

## Examples

### Json

````json
{
  "scaffolding": {
    "name": "my-go-project",
    "files": [
      "main.go",
      "go.mod"
    ],
    "dirs": [
      {
        "name": "utils",
        "files": [
          "utils.go"
        ]
      }
    ]
  }
}
````

### Yaml

````yml
scaffolding:
  name: "my-go-project"
  files:
    - main.go
    - go.mod
  dirs:
    - name: utils
      files:
        - utils.go
````

# Placeholder Usage

The configuration admits placeholders of the form `${PLACEHOLDER}`.
Once a placeholder is detected it will require manual action of the user for replace it.
This will be done by console input like this:
```
enter the placeholder: ...
```
*The 'placeholder' word will be the actual placeholder name.*

## Example

````yml
scaffolding:
  name: "${project-name}"
  files:
    - main.go
    - go.mod
  dirs:
    - name: "${project-name}_stuff"
      files:
        - "${project-name}_stuff_${stuff-name}.go"
````

# Command Line

The `gosca.exe` has its own help command *-h*.

## Create command
The `create` command takes two arguments: the scaffolding config and  the location of the project.

```
gosca create ./config.json ./example
```

In case of use a yaml as Scaffolding configuration, use the `--yaml -y` flag.

```
gosca create ./config.yml ./example -y
```
