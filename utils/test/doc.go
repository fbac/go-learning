//	fbac test's test pkg
//
//	https://pkg.go.dev/testing
//
//	- Documentation links
//	https://www.golang-book.com/books/intro/12
//	https://go.dev/blog/subtests
//	https://go.dev/blog/fuzz-beta
//	https://go.dev/blog/appengine-dec2013
//
//	- Test Examples
//	https://go.dev/blog/examples
//
//	- Test key tips
//
//	files:
//	main_test.go: traditional unit test, simple func test
//	main_table_test.go: table tests using structs that contain many cases
//
//	* test files ends in '_test.go'.
//	* the test file has to be in the same package as the functions that it's going to test.
//	* func signatures must be 'func TestFuncName(t *testing.T)'.
//	* Run tests with 'go test' (ensure env is correctly configured).
//	* Return failures with Error/Fail or related methods to properly signal a test failure.
//	* Whenever possible, use autogenerated test templates, to reduce complexity and time.
//	* Inside the test functions run whatever code helps to test the associated functions.
//
//	- Document test via Examples
//
//	* Examples are tested by go test as a normal test would.
//	* They're published in godoc accordingly.
//

package testdocs
