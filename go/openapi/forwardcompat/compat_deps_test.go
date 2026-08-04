package forwardcompat

// Blank import anchors gopkg.in/validator.v2 in go.mod so go mod tidy keeps it.
// The OpenAPI compat test script generates code (withGoMod=false) into
// .compat-tmp/ that imports this package, and compiles it as part of the root
// module — so the root go.mod must provide it.
import _ "gopkg.in/validator.v2"
