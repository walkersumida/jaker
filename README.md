# Jaker(`Japanese` + `Faker`)
![badge](https://github.com/walkersumida/faker/actions/workflows/test.yml/badge.svg)

Jaker is a faker for Japanese.

Jakerは日本語のためのフェイカーです。

## Getting Started
### Download

```
go get -u github.com/walkersumida/jaker
```

### Usage

```golang
import 	"github.com/walkersumida/jaker"

profile := jaker.Profile

fmt.Println(profile.EnCompany)
// => "Suzuki Capital Partners, Inc."

fmt.Println(profile.JaKanjiFullName)
// => "鈴木 一郎"

fmt.Println(profile.EnFullName)
// => "Ichiro Suzuki"

fmt.Println(profile.Email)
// => "ichiro.suzuki@example.com"

fmt.Println(jaker.Uuid)
// => "add0de95-6d27-44de-ab7a-d45998bc6b05"

fmt.Println(jaker.Text("abcあいう", 10))
// => "abcあいうabcあ"
```
