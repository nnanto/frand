# FRAND
frand generates fast random unsigned integers (non-cryptographic).

Uses [Marsaglia's XOR shift](https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf) similar to the frand function in go runtime

## Usage

```go
import (
	"fmt"
    "github.com/nnanto/frand"
)

func main() {
    x := frand.GetN(100)
    fmt.Printf("Random Number [0,100) :%v")
}

```