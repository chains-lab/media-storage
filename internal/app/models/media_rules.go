package models

type AllowedExtensionModel struct {
	Extension string `db:"extension"`
	MaxSize   int64  `db:"max_size"`
}

type MediaRules struct {
	Resource   string
	Category   string
	Extensions []AllowedExtensionModel
}
