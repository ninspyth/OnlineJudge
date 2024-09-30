

// Connection and channel for the sender to send tasks to queue

// func ExecuteCode(FileName string, c *gin.Context) (codeOutput string) {
//
// 	var dir string = "test/"
// 	trimmer := "." + SubmitRequestBody.Language
// 	var outputFile string = strings.Trim(FileName, trimmer)
//
// 	//compiling for c lang
// 	compileCommand := exec.Command("gcc", fmt.Sprintf("%s%s", dir, FileName), "-o", fmt.Sprintf("%s%s", dir, outputFile))
//
// 	//Buffer to store the error of the compilation
// 	var compileError bytes.Buffer
// 	compileCommand.Stderr = &compileError
//
// 	if err := compileCommand.Run(); err != nil {
// 		//if we get any compile-time error send the error to the client
// 		// fmt.Println("Error Compiling: \n", compileError.String())
// 		c.String(OK, "Compilation Failed: %s", compileError.String())
// 	}
//
// 	// fmt.Println("filename: %s, dir: %s, outputFile: %s", FileName, dir, outputFile)
// 	//Compile the binary
// 	binExec := exec.Command(fmt.Sprintf("%s./%s", dir, outputFile))
//
// 	//Buffer to store the outptu of the code
// 	var binOutput bytes.Buffer
// 	binExec.Stdout = &binOutput
//
// 	if err := binExec.Run(); err != nil {
// 		//if any runtime error we will send it to the client
// 		// fmt.Println("Execution Failed: ", err)
// 		c.String(OK, "Runtime Error: %s", binOutput.String())
// 		return
// 	}
//
// 	codeOutput = binOutput.String()
// 	return
// }

// //store the Request Body
// if err := json.NewDecoder(c.Request.Body).Decode(&SubmitRequestBody); err != nil {
// 	c.String(http.StatusBadRequest, "Incorrect Request parameter")
// 	return
// }
//
// //file to store the user code
// var FileName string = "user_code"
// var dir string = "test/"
//
// FileName = FileName + "." + SubmitRequestBody.Language
// var uCode string = dir + FileName
//
// UserCodeAsBytes := []byte(SubmitRequestBody.UserCode)
//
// //write the content of user code to the file
// UserCodeFile, err := os.Create(uCode)
// if err != nil {
// 	log.Fatal(err)
// }
//
// //close the file after use
// defer UserCodeFile.Close()
//
// //write the usercode into the file
// if err := os.WriteFile(uCode, UserCodeAsBytes, 0744); err != nil {
// 	log.Fatal(err)
// }
//
// //execute the code
// Output := ExecuteCode(FileName, c)
// c.String(OK, "Execution Succesful: \n%s", Output)
// fmt.Println("The Output of the code is: \n", Output)

//This is a test go File
