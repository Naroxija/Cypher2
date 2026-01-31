package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"  
)
func main() {
    argumentsPassed:= os.Args[1:]


    fileName:= strings.Join(argumentsPassed," ")
    existingFiles, _ := os.ReadDir("./")

    existingFilesSlice:= []string{} 

    for _, file := range existingFiles {
        if strings.Contains(file.Name(), ".txt") {
            existingFilesSlice = append(existingFilesSlice, file.Name())
        }
    }    

    if IsValidFileName(existingFilesSlice, fileName) {
        var operation = Prompter()  
        operation = IsValidOperation(operation)     
        ExecuteGivenOperation(operation, fileName)
    }
    
}

func IsValidFileName(existingFiles []string, fileName string) bool {

    scanner := bufio.NewScanner(os.Stdin)
    var result string     
    fileName = strings.TrimSpace(fileName)
    var collectionId int = 1


    if DoesFileExistInDirectory(existingFiles, fileName) {
        fmt.Printf("The collection: %s does exist\n", fileName)
        return true         
    } else {
        fmt.Println("Usage: ./todotool")
        fmt.Println("You have passed non existing collection name. Please try one more time")  
        fmt.Println("Choose from existing collections: ")
        for _, file := range existingFiles {
            fmt.Printf("%d - %s\n", collectionId, file)
            collectionId++
        }                
        for {
            if scanner.Scan() {
                result = scanner.Text()
                result = strings.TrimSpace(result)
                collectionId = 1

                if DoesFileExistInDirectory(existingFiles, result) {
                    return true
                } else {
                    fmt.Println("Choose from existing collections: ")
                    for _, file := range existingFiles {
                        fmt.Printf("%d - %s\n", collectionId, file)
                        collectionId++
                    }                        
                }
            }
        }
    }
        
    return false

}

func DoesFileExistInDirectory(existingFiles []string, fileName string) bool{

    for _, file := range existingFiles {
        if file == fileName {
            return true
        }
    }    
        
    return false

}

func Prompter() string {
    
    fmt.Println("\nWelcome to the notes tool!\n")
    fmt.Println("Select operation:")    

    options:= []string{"1. Show notes.", "2. Add a note.", "3. Delete a note.", "4. Exit."}

    for i:=0; len(options)>i; i++ {
        fmt.Println(options[i])
    }

    scanner := bufio.NewScanner(os.Stdin)
    var operation string

    if scanner.Scan() { 
        operation = scanner.Text()
    }

    return operation
}

func ExecuteGivenOperation(operation string, fileName string) {

    //reader := bufio.NewReader(os.Stdin) 

    switch operation {
    case "1":
        fmt.Println("Notes")
        displayNotes(fileName)
    case "2":
        Write(fileName)
    case "3":
        fmt.Println("Delete a note.")
    case "4":
        fmt.Println("Exit.")                        
    }
    
}

func IsValidOperation(operation string) string{

    scanner := bufio.NewScanner(os.Stdin)
    var result string 

    operation = IsValidInput(operation)
    convertedOperation, _ := strconv.Atoi(operation) 

    if convertedOperation>=1 && convertedOperation<=4 {
        return operation
    } else {
        
        fmt.Printf("The operation %s is not between 1 and 4. Give it another try.\n", operation) 

        for {    
            if scanner.Scan() {
                anotherOperation:= scanner.Text()
                anotherOperation = strings.ReplaceAll(anotherOperation, " ", "")
                anotherConvertedOperation, _ := strconv.Atoi(anotherOperation) 

                if anotherConvertedOperation>=1 && anotherConvertedOperation<=4 {
                    result = strconv.Itoa(anotherConvertedOperation) 
                    break
                } else {
                    fmt.Printf("The operation %s is not between 1 and 4. Give it another try.\n", anotherOperation)                         
                }              
            }

        }
    }

    return result
}

func IsValidInput(operation string) string {

    operation = strings.TrimSpace(operation)

	if operation != "" && len(operation)>=1 {
		return operation
	}

	var notEmptyOperation string
	scanner := bufio.NewScanner(os.Stdin)

	for{ 
		fmt.Println("You are passing an empty input. Please try one more time")

		if scanner.Scan() {
			anotherOperation:= scanner.Text()
            anotherOperation = strings.TrimSpace(anotherOperation)

			if len(anotherOperation)>=1 && anotherOperation != "" {
				notEmptyOperation = anotherOperation
				break;
			}
		}
	}
	
	return notEmptyOperation    
}

////File operations

func loadNotes(filename string) []string{
	file, err :=os.Open(filename)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}
	return notes
}

/////File logic

func displayNotes(filename string) {
	notes := loadNotes(filename)
	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return
	}
	for i, note := range notes {
		fmt.Printf("%d: %s\n", i+1, note)
	}
}


//Writing to a file 

func Write(fileName string) {
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
    if err != nil {
        
    }
    defer file.Close()

    fmt.Println("Enter a note")
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan() {
        
        text:= "\n" + scanner.Text()
        msg := []byte(text)
        _, err = file.Write(msg)
        if err != nil {
            
        }

    }
    

    fmt.Println("Note appended.")
}
