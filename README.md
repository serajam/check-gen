# check-gen

Generates validation code for different types including nested structs, maps, slices.  
Returns errors as formatted []string: field_snake_case::is_required, field_snake_case::min_length_is  
See examples for more info  

## **Any ideas and help are welcome**

## Installation

```shell
go get -v github.com/serajam/check-gen
```

## Usage

It is possible to validate:

* `int`,  `*int`
* `int8`,  `*int8`
* `int16`,  `*int16`
* `int32`,  `*int32`
* `int64`,  `*int64`
* `[]int`,  `[]*int`
* `[]int8`,  `[]*int8`
* `[]int16`,  `[]*int16`
* `[]int32`,  `[]*int32`
* `[]int64`,  `[]*int64`
* `float32`,  `*float32`
* `float64`,  `*float64`
* `[]float32`,  `[]*float32`
* `[]float64`,  `[]*float64`
* `string`,  `*string`
* `*bool`
* `[]byte` , `[]*byte`
* `struct`,  `*struct`
* `[]struct`,  `[]*struct`
* `map[string]type`  
* `custom types`

1. Add tags `check:""` for every field that's needs checks. For example: `check:"required,min=1"` See available tags below
2. Run `check-gen` with path as last argument. Also, you can add `--copyright=<string>` or `-c <string>` for add a copyright to generated files
3. You should see `validation.go` with validation methods for every struct in every package that was specified fo/r validation generator
   Use `Validate()` method to check your struct data

## Tags

1. `required` - checks if field is set to non default value (not 0, not nil, not "", len > 0 for slices)
2. `min` - checks field min value
3. `max` - checks field max value
4. `len` - checks field len
5. `check` - add call to `Validate` method of nested struct
6. `deep` - checks every slice value calling `Validate` method for structs and generates validation for values of `int, string` types  specified by tags after `deep`. For example `required,min=1,deep,len=100` will check if slice has minimum len of 1 and also will check if every element of that slice has len exactly 100. You can use `check` with `deep`,  for example `required,min=1,deep,check` will check if slice maximum len of 1 and add call to `Validate` method for each of these structs.
7. `digit` will check if string contains only number
8. `word` will check if string contains only letters, numbers, punctuation and whitespaces
9. `uuid` will check if string is uuid
10. `password` will check if string is password with len check. For example: `min=8,password`
10. `phone` will check if string is satisfy phone format with len check. Checks only digits. Returns no errors with empty string, if not need this add `required` tag. For example: `phone=7~15`

## Examples

Please see `examples` for full list of examples
