com arg: add
  git commit -m "{{arg}}"

add:
  git add .

tag arg1 arg2:
  git tag -a {{arg1}} -m "{{arg2}}"

release:
  goreleaser release --clean

