## TRY-CATCH-FINALLY
--------------------

*note. Try-Catch-Finally Concept* 

	Try: code inside this block will be run first, if there's any errors, then use 'Throw()' to catch that error/exception.
	Catch: code inside this block will be run, if error occurs.
	Finally: regardless of the result, this code will be executed.

<pre><code>
func Throw(e Exception) {
	panic(e)
}
func (b *Block) Do(){
    b.Finally(){
		defer b.Finally()
    }
    // There will be a program inside this catch block to handle error
    if b.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				b.Catch(r)
			}
		}()
	}
    b.Try()	    

}
</code></pre>

<pre><code>
func main(){
   var block = Block{
		Try: func() {
			println("Try Block")
            Throw("this is error")
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
</code></pre>

In function main, variable <code>block</code> called *Do()*. 
In function do, block <code>Finally</code> is called first, because *program inside this block will be run regardless of the result, but will be run after <code>Try</code> if it's with no exception*

to make <code>Finally</code> block run the last, we use <code>defer</code>, make sure this block as a first defer, in which it will be run the last. [concept. first in last out].

<code>Catch</code> block will catch the error. When the error occurred, *recover()* allows program to manage behaviour of a panicking goroutines, pass the error message of panic. Use defer, to stop panic seq.

    Illustration. 
    [FinallyDefer <- CatchDefer <- Try]
    
    FinallyDefer stops the last. [ defer finally() ]
    CatchDefer stops before FinallyDefer. [ defer func(){recover...} ]
    Try runs first, stop after panic. Panic will exit the program but need to finish all the defer program. 




