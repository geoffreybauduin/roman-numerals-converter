//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package roman

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test1to3 (t *testing.T) {
    assert.Equal(t, "I", IntToRoman(1), "1 should be I")
    assert.Equal(t, "II", IntToRoman(2), "2 should be II")
    assert.Equal(t, "III", IntToRoman(3), "3 should be III")
}

func Test5to8 (t *testing.T) {
    assert.Equal(t, "V", IntToRoman(5), "5 should be V")
    assert.Equal(t, "VI", IntToRoman(6), "6 should be VI")
    assert.Equal(t, "VII", IntToRoman(7), "7 should be VII")
    assert.Equal(t, "VIII", IntToRoman(8), "8 should be VIII")
}

func TestEntireValues (t *testing.T) {
    assert.Equal(t, "X", IntToRoman(10), "10 should be X")
    assert.Equal(t, "L", IntToRoman(50), "50 should be L")
    assert.Equal(t, "C", IntToRoman(100), "100 should be C")
    assert.Equal(t, "D", IntToRoman(500), "500 should be D")
    assert.Equal(t, "M", IntToRoman(1000), "1000 should be M")
    assert.Equal(t, "V̅", IntToRoman(5000), "5000 should be V̅")
}

func TestCombined (t *testing.T) {
    assert.Equal(t, "MMXVI", IntToRoman(2016), "2016 should be MMXVI")
}

func TestSubstractions (t *testing.T) {
    assert.Equal(t, "IV", IntToRoman(4), "4 should be IV")
    assert.Equal(t, "IX", IntToRoman(9), "9 should be IX")
    assert.Equal(t, "XL", IntToRoman(40), "40 should be XL")
    assert.Equal(t, "IL", IntToRoman(49), "49 should be IL")
    assert.Equal(t, "LIX", IntToRoman(59), "59 should be LIX")
    assert.Equal(t, "XDIV", IntToRoman(494), "494 should be XDIV")
    assert.Equal(t, "XMIV", IntToRoman(994), "994 should be XMIV")
    assert.Equal(t, "IM", IntToRoman(999), "999 should be IM")
    assert.Equal(t, "MDCXCIII", IntToRoman(1693), "1693 should be MDCXCIII")
    assert.Equal(t, "MV̅DCCXLIV", IntToRoman(4744), "4744 should be MV̅DCCXLIV")
    assert.Equal(t, "XX̅IV", IntToRoman(9994), "9994 should be XX̅IV")
    assert.Equal(t, "C̅M̅L̅X̅X̅X̅V̅MMCCXXXI", IntToRoman(987231), "987231 should be C̅M̅L̅X̅X̅X̅V̅MMCCXXXI")
}