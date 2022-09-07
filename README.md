# pkg

Openware go packages

## Tagging submodules

Whenever you'd like to release a new version of a submodule(e.g. `ika`), simply run `git tag *module_name*/vX.X.X` & `git push --tags`
This will create a tag exlusively for this submodule so that it'd be available by `go get github.com/openware/pkg/*module_name*@vX.X.X` afterwards

## Release v0.1.3
