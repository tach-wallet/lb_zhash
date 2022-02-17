package zhash_test

import (
	"bytes"
	"fmt"
	zhash "github.com/tach-wallet/lb_zhash"
)

package zhash_test

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/tach-wallet/lb_zhash"
)

func unmarshalToml(d []byte, t interface{}) error {
	_, err := toml.Decode(string(d), t)
	return err
}

func Example_tomlUnmarshaller() {
	h := zhash.NewHash()
	h.SetUnmarshallerFunc(unmarshalToml)

	blob := []byte(`
key1 = "string"
key2 = 10
`)

	b := bytes.NewBuffer(blob)
	h.ReadHash(b)
	fmt.Println(h.String())

	//Output:
	//{
	//   "key1": "string",
	//   "key2": 10
	//}
}
