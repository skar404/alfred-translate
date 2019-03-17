package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	InitEnv()
}

func TestDetectLangRun(t *testing.T) {
	lang := DetectLang("привет мир")

	assert.Equal(t, lang, "ru")
}

func TestDetectLangEng(t *testing.T) {
	lang := DetectLang("hello world")

	assert.Equal(t, lang, "en")
}
