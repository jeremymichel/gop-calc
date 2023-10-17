# GOP Calculator

Single function exported to calculate the GOP (B-Frames) for a video with a specified target.

## Usage

First install the dependency

```shell
go get github.com/jeremymichel/gop-calc
```

To use the function

```go
import "github.com/jeremymichel/gop-calc"

func main() {
	gopInSecs, gopInFrames := gop_calc.Calculate(25, 48000, 1024, 2) // We target a max gop of 2 seconds
	// Will return 1.92 and 48 for this example
}
```


