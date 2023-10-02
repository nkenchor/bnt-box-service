package dto

type BoxConfig struct {
	PrefixStart    string `json:"prefix_start"`
	PrefixEnd   string `json:"prefix_end"`
	PrefixType string `json:"prefix_type"`
	IndentRef string `json:"indent_ref"`
}
