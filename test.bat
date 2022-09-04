go test ./... -timeout 5s

@REM run specific test
@REM go test ./lexer -run TestTokenizer/for_loop -v