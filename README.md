# Pikkā

A simple `map[string]interface{}` deep picker from a string path.

## Example

Basic example when you unmarshal some json into a map and too lazy to make a struct for it:

```go
data := map[string]interface{}{}
json.Unmarshal(jsonBytes, &data)

fmt.Println(pikka.GetFromPath("foo.bar", data).String())
```

This will print the value of the field `bar` or an empty string.

## Installation

    go get github.com/maxperrimond/pikka

## Usage

### Path

The path can be formatted like:

 - `foo.bar` For nested values
 - `foo.someArray.1.bar` For nested array value
 - `foo.someArray.#.bar` For all nested array values

### Result

`GetFromPath` returns a result wrapper (`Value`) which provides some casting getters for some types and slice of it.

Note: If a field for the path is not found, it will return `nil`.
