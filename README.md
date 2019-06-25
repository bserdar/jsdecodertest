Small test to see if using a literal buffer for keys in JSON decoder
improves memory usage. It does.

This test builds a large JSON string, decodes it into an interface{},
runs GC and prints memory stats.

To use standard decoder:

``` go
import (
	//"github.com/bserdar/jsdecodertest/json"
	"encoding/json"
)

```


To use the modified decoder:
``` go
import (
	"github.com/bserdar/jsdecodertest/json"
	//"encoding/json"
)

```

The jsdecodertest/json package is a copy of the standard library
encoding/json, with a literal map in the decoder
state. decoderState.objectInterface uses the literal map to store
field keys in a map.

Using standard library:
Alloc:401839488 

Use modified decoder:
Alloc:353839856 

This simple test only works when decoding into an interface{}.
