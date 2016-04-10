//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package roman

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test1to3 (t *testing.T) {
    assert.Equal(t, IntToRoman(1), "I", "1 should be I")
    assert.Equal(t, IntToRoman(2), "II", "2 should be II")
    assert.Equal(t, IntToRoman(3), "III", "3 should be III")
}

func Test5to8 (t *testing.T) {
    assert.Equal(t, IntToRoman(5), "V", "5 should be V")
    assert.Equal(t, IntToRoman(6), "VI", "6 should be VI")
    assert.Equal(t, IntToRoman(7), "VII", "7 should be VII")
    assert.Equal(t, IntToRoman(8), "VIII", "8 should be VIII")
}

func TestEntireValues (t *testing.T) {
    assert.Equal(t, IntToRoman(10), "X", "10 should be X")
    assert.Equal(t, IntToRoman(50), "L", "50 should be L")
    assert.Equal(t, IntToRoman(100), "C", "100 should be C")
    assert.Equal(t, IntToRoman(500), "D", "500 should be D")
    assert.Equal(t, IntToRoman(1000), "M", "1000 should be M")
}

func TestCombined (t *testing.T) {
    assert.Equal(t, IntToRoman(2016), "MMXVI", "2016 should be MMXVI")
}

func TestSubstractions (t *testing.T) {
    assert.Equal(t, IntToRoman(4), "IV", "4 should be IV")
    assert.Equal(t, IntToRoman(9), "IX", "9 should be IX")
    assert.Equal(t, IntToRoman(40), "XL", "40 should be XL")
    assert.Equal(t, IntToRoman(49), "IL", "49 should be IL")
    assert.Equal(t, IntToRoman(59), "LIX", "59 should be LIX")
    assert.Equal(t, IntToRoman(994), "XMIV", "994 should be XMIV")
}