package decorator

import "fmt"

type Datasource interface {
	Read() string
	Write(value string)
}

type BaseDecorator struct {
	nextLayer Datasource
}

func (b *BaseDecorator) Read() string {
	return b.nextLayer.Read()
}

func (b *BaseDecorator) Write(value string) {
	b.nextLayer.Write(value)
}

type FileDatasource struct{}

func (f *FileDatasource) Read() string {
	return "Base"
}

func (f *FileDatasource) Write(value string) {
	fmt.Println(value)
}

type CompressionDecorator struct {
	BaseDecorator
}

func NewCompressionDecorator(datasource Datasource) *CompressionDecorator {
	return &CompressionDecorator{
		BaseDecorator{nextLayer: datasource},
	}
}

func (c *CompressionDecorator) Read() string {
	compressed := c.nextLayer.Read()
	return c.decompress(compressed)
}

func (c *CompressionDecorator) decompress(compressed string) string {
	return compressed + " - Decompressed"
}

func (c *CompressionDecorator) Write(value string) {
	compressed := c.compress(value)
	c.nextLayer.Write(compressed)
}

func (c *CompressionDecorator) compress(value string) string {
	return value + " - Compressed"
}

type EncryptionDecorator struct {
	BaseDecorator
}

func NewEncryptionDecorator(datasource Datasource) *EncryptionDecorator {
	return &EncryptionDecorator{
		BaseDecorator{nextLayer: datasource},
	}
}

func (e *EncryptionDecorator) Read() string {
	value := e.nextLayer.Read()
	return e.decrypt(value)
}

func (e *EncryptionDecorator) decrypt(value string) string {
	return value + " - Decrypted"
}

func (e *EncryptionDecorator) Write(value string) {
	encrypted := e.encrypt(value)
	e.nextLayer.Write(encrypted)
}

func (e *EncryptionDecorator) encrypt(value string) string {
	return value + " - Encrypted"
}
