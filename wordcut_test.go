package gothaiwordcut

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPureThaiCut(t *testing.T) {
	segmenter := Wordcut()
	segmenter.LoadDefaultDict("")
	result := segmenter.Segment("ทดสอบการตัดคำภาษาไทย")
	assert.Equal(t, []string{"ทดสอบ", "การ", "ตัด", "คำ", "ภาษาไทย"}, result)
}

func TestMixEnglishThaiCut(t *testing.T) {
	segmenter := Wordcut()
	segmenter.LoadDefaultDict("")
	result := segmenter.Segment("มาลองตัดคำปนภาษา English กันนะ Alright เพื่อน")
	assert.Equal(t, []string{"มา", "ลอง", "ตัด", "คำ", "ปน", "ภาษา", "English", "กัน", "นะ", "Alright", "เพื่อน"}, result)
}

func TestDash(t *testing.T) {
	segmenter := Wordcut()
	segmenter.LoadDefaultDict("")
	result := segmenter.Segment("e-traning")
	assert.Equal(t, []string{"e", "-", "traning"}, result)
}

func TestDot(t *testing.T) {
	segmenter := Wordcut()
	segmenter.LoadDefaultDict("")
	result := segmenter.Segment("1.1 for traning.")
	assert.Equal(t, []string{"1", ".", "1", "for", "traning", "."}, result)
}
func TestThaiWithNumber(t *testing.T) {
	segmenter := Wordcut()
	segmenter.LoadDefaultDict("")
	result := segmenter.Segment("การติดตั้งโปรแกรม Bplus HRM v7.3")
	assert.Equal(t, []string{"การ", "ติดตั้ง", "โปรแกรม", "Bplus", "HRM", "v7", ".", "3"}, result)
}
func BenchmarkWordcut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		segmenter := Wordcut()
		segmenter.LoadDefaultDict("")
		dat, _ := ioutil.ReadFile("./dict/benchmark_text.txt")
		segmenter.Segment(string(dat))
	}
}
