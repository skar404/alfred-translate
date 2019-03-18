package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	Init()
}

func TestDetectLangRun(t *testing.T) {
	lang := DetectLang("привет мир")

	assert.Equal(t, lang, "ru")
}

func TestDetectLangEng(t *testing.T) {
	lang := DetectLang("hello world")

	assert.Equal(t, lang, "en")
}
