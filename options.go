package goose

import "io/fs"

func WithAllowMissing() OptionsFunc {
	return func(o *options) { o.allowMissing = true }
}

func WithNoVersioning() OptionsFunc {
	return func(o *options) { o.noVersioning = true }
}

func WithVerbose() OptionsFunc {
	return func(o *options) { o.verbose = true }
}

func WithTableName(name string) OptionsFunc {
	return func(o *options) { o.tableName = name }
}

func WithLogger(l Logger) OptionsFunc {
	return func(o *options) { o.logger = l }
}

func WithEmbededFS(filesystem fs.FS) OptionsFunc {
	return func(o *options) { o.filesystem = filesystem }
}

type OptionsFunc func(o *options)

type options struct {
	filesystem fs.FS

	allowMissing bool
	noVersioning bool
	verbose      bool
	tableName    string
	logger       Logger

	applyUpByOne bool
}

func withApplyUpByOne() OptionsFunc {
	return func(o *options) { o.applyUpByOne = true }
}
