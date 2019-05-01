package version

var (
	Code string
)

func init() {
	if Code == "" {
		Code = "Unknown version"
	} else if len(Code) > 8 {
		Code = Code[:8]
	}
}
