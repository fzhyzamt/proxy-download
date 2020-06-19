package config

type Config struct {
	// 下载缓冲区大小 (Byte)
	BufferSize       int
	Port             int
	ExcludeHeaderKey map[string]bool
	IndexHtml        string
}
