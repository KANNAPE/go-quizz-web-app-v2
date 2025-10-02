package frontend

import "embed"

//go:embed templates/*.html
var Tempates embed.FS

//go:embed static/*
var Static embed.FS
