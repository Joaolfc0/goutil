package textscan_test

import (
	"testing"

	"github.com/Joaolfc0/goutil/dump"
	"github.com/Joaolfc0/goutil/strutil/textscan"
	"github.com/Joaolfc0/goutil/testutil/assert"
)

func TestParser_ParseText(t *testing.T) {
	p := textscan.NewParser(func(t textscan.Token) {
		dump.P(t)
	})

	err := p.ParseText(`
# comments 1
# comments 1.1
# comments 1.2
name = inhere

// comments 2
age = 28

/*
multi line
comments 3
*/
desc = '''
a multi
line string
'''
`)
	assert.NoErr(t, err)

}
