package probe

const (
	B float64 = 1 << (10 * iota)
	KB
	MB
)

var sizeMap = map[string]float64{
	"B":  B,
	"KB": KB,
	"MB": MB,
	"b":  B,
	"kb": KB,
	"mb": MB,
	"Kb": KB,
	"Mb": MB,
}

type Probe struct {
	Filepath string "图片路径"
}
