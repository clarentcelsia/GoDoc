package utils

type Exception interface{}

type (
	Block struct {
		Try     func()
		Catch   func(e Exception)
		Finally func()
	}
)

// note. this panic will be executed last.
// defer doesn't affect.
func Throw(e Exception) {
	panic(e)
}

func (b *Block) Do() {
	if b.Finally != nil {
		// Check and run a statement inside this block
		// If exists, call finally() with a 'defer' to make it being executed the last.
		defer b.Finally()
	}
	if b.Catch != nil {
		// triggers if the error occurred (panic)
		defer func() {
			// recover the program, then catch the error
			if r := recover(); r != nil {
				b.Catch(r)
			}
		}()
	}
	b.Try()
}

func TryCatch() {
	var block = Block{
		Try: func() {
			println("Try Block")
		},
		Catch: func(e Exception) {
			println(e)
		},
		Finally: func() {
			println("Finally Block")
		},
	}

	block.Do()
}
