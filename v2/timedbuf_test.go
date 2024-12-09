/*
 *   Copyright 2016 Charith Ellawala
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package timedbuf

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFlushOnFull(t *testing.T) {
	var bufferContents []string
	tb := New(5, 30*time.Second, func(items []string) { bufferContents = append(bufferContents, items...) })
	for i := 0; i < 8; i++ {
		tb.Put(fmt.Sprintf("item%d", i))
		if i < 5 {
			assert.Equal(t, 0, len(bufferContents))
		}
	}

	// timer hasn't fired yet, so the buffer should only contain 5 items
	assert.Equal(t, 5, len(bufferContents))
	// after closing, the buffer should be completely emptied
	tb.Close()
	assert.Equal(t, 8, len(bufferContents))
}

func TestFlushOnTimer(t *testing.T) {
	var bufferContents []string
	tb := New(5, 1*time.Second, func(items []string) { bufferContents = append(bufferContents, items...) })
	for i := 0; i < 3; i++ {
		tb.Put(fmt.Sprintf("item%d", i))
		assert.Equal(t, 0, len(bufferContents))
	}

	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, 3, len(bufferContents))
	tb.Close()
	assert.Equal(t, 3, len(bufferContents))
}
