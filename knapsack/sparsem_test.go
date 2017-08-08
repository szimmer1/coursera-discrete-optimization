package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSparseMSet(t *testing.T) {
	m := &sparseM{}
	/*
		[ 3, 1, nil]
		[ 2, 0, 4]
	*/

	err := m.Set(1, 1, 0)
	assert.Nil(t, err)

	err = m.Set(0, 1, 1)
	assert.Nil(t, err)

	err = m.Set(1, 0, 2)
	assert.Nil(t, err)

	err = m.Set(0, 0, 3)
	assert.Nil(t, err)

	err = m.Set(1, 2, 4)
	assert.Nil(t, err)

	/*
	   [ 0,  ,  ]
	   [  , 1, 2]
	*/
	var n int
	n, err = m.Get(1, 2)
	assert.Equal(t, 4, n)

	n, err = m.Get(1, 1)
	assert.Equal(t, 0, n)

	n, err = m.Get(0, 0)
	assert.Equal(t, 3, n)

	n, err = m.Get(0, 4)
	assert.NotNil(t, err)

	n, err = m.Get(2, 0)
	assert.NotNil(t, err)

	Inorder(m.root, func(r *BTNode) {
		t.Logf("[%s]\n", SerializeTreeInorder(r))
	})
}
