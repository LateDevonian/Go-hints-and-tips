package hollowworld
//zaks testing workshop
import (
    "testing"

    "github.com/stretchr/testify/assert"
)

// Task 1: Run this test locally
// You can run this test in the editor or from the command line by running:
// $ go test

func TestHollowWorld(t *testing.T) {
    // Given
    name := "Jeremy"

    // When
    helloText := SayHello(name)

    // Then
    assert.Contains(t, helloText, name)
    assert.Contains(t, helloText, "Hello")
}

// Not working? Run it in browser: https://goplay.space/#drsSbqu10Uh

