//go:generate go run ../../generate/tags/main.go -AWSSDKVersion=2 -ListTags -UpdateTags -ServiceTagsSlice
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package scheduler
