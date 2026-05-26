package windivert

import (
	"strings"
	"testing"

	"github.com/ddkwork/bindgen/c2go"
)

func TestGenerate(t *testing.T) {
	c2go.Generate(t, []c2go.BindgenConfig{{
		HeadersDir:  "clone/include",
		OutputDir:   ".",
		PackageName: "windivert",
		HeaderOrder: []string{"windivert.h"},
		Predefined: `
#include <windows.h>
#define __in
#define __in_opt
#define __out
#define __out_opt
#define __inout
#define __inout_opt
`,
		BindDll: true,
		DllName: "WinDivert.dll",
		DllFuncFilter: func(name string) bool {
			return strings.HasPrefix(name, "WinDivert")
		},
	}})
}
