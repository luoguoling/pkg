使用办法:
import (
	"fmt"
	"github.com/luoguoling/pkg/genid"
)
func main() {
	genid.Init("2021-10-28",22)
	fmt.Println(genid.GenID())
}
