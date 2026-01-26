# PDF Reader (Superpower Fork)

Fork of [github.com/ledongthuc/pdf](https://github.com/ledongthuc/pdf) with fixes for malformed PDF CMap handling.

## Why This Fork Exists

The upstream library fails to extract text from PDFs that have malformed CMap (character mapping) definitions. Specifically, some PDF generators create CMaps where:

- The **codespace range** is defined too narrowly (e.g., `<51>` to `<79>`)
- But the **bfrange mappings** include characters outside that range (e.g., `<20>` for space, `<30>` for digit 0)

The upstream library's `cmap.Decode()` function only looks up character mappings if the byte falls within a declared codespace range. Characters outside the codespace are replaced with the Unicode replacement character, resulting in garbled text output.

### In Simple Terms

**Without the patch**, extracting text from affected PDFs produces garbage:
```
��������� ������ - �����
���������� �����*:                    ��� ��/��
```

**With the patch**, you get the actual text:
```
Glyphosate Profile - Urine
Creatinine Value*:                    100 mg/dl
```

The fix allows the library to find character mappings even when the PDF's internal character map is technically malformed.

## Changes From Upstream

**File: `page.go`** - Modified `cmap.Decode()` function

Added a fallback that tries bfchar/bfrange lookups even when bytes are outside the declared codespace. This handles malformed CMaps that define mappings outside their declared codespace range.

```go
// After the codespace loop fails, try bfchar/bfrange lookups for single-byte
// codes even when outside the declared codespace. Some PDFs have malformed
// CMaps where the codespace range is too narrow but bfrange mappings exist.
```

## Install

```bash
go get -u github.com/superpowerdotcom/go-pdf-lib
```

## Usage

```golang
package main

import (
    "bytes"
    "fmt"

    pdf "github.com/superpowerdotcom/go-pdf-lib"
)

func main() {
    f, r, err := pdf.Open("./test.pdf")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    var buf bytes.Buffer
    b, err := r.GetPlainText()
    if err != nil {
        panic(err)
    }
    buf.ReadFrom(b)
    fmt.Println(buf.String())
}
```

## Read Text By Row (preserves layout)

```golang
package main

import (
    "fmt"

    pdf "github.com/superpowerdotcom/go-pdf-lib"
)

func main() {
    f, r, err := pdf.Open("./test.pdf")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    p := r.Page(1)
    rows, err := p.GetTextByRow()
    if err != nil {
        panic(err)
    }

    for _, row := range rows {
        fmt.Printf("Y=%.1f: ", row.Content[0].Y)
        for _, word := range row.Content {
            fmt.Print(word.S, " ")
        }
        fmt.Println()
    }
}
```

## Original Features

- Get plain text content (without format)
- Get Content (including all font and formatting information)
- Get text grouped by rows or columns

## License

BSD-style license (see LICENSE file)
