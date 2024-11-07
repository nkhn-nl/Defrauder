package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	"bufio"
	"os/exec"
	"strings"
	"unicode"
 	"strconv"
)

var wg sync.WaitGroup
var limitChan chan struct{} 

func initLimitChan(bufferSize int) {
	
	limitChan = make(chan struct{}, bufferSize)

}

func clearTerminal() {
	
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func banner() {
	cmd := exec.Command("bash", "-c", "figlet -f big Defrauder | lolcat ")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	}
	fmt.Println("[+] Generates domain variations by swapping characters ")
	fmt.Println("[+] Running custom algorithm to alter characters at different positions. ")
	fmt.Println("[+] Live Check: Verifies which generated domains are currently active. ")
	fmt.Println("---------------------------------------------------------------------------")
}

func displayHelp() {
    fmt.Println("Usage:  Defrauder.go -d <domain> -o <output_file> -t <buffer_size>")
    fmt.Println("\nFlags:")
    fmt.Println("  -d <domain>       Target domain to check for fakes.")
    fmt.Println("  -o <output_file>  Specify the output file for results.")
    fmt.Println("  -t <buffer_size>  Set the buffer size for concurrent checks (default is 32).")
    fmt.Println("  -h                This help text.")
    fmt.Println("\nExample:")
    fmt.Println(" Defrauder.go -d example.com -o results.txt -t 40")
}

func isDomainLive(domain string) {

    var wg sync.WaitGroup
	
    wg.Add(1)
    defer wg.Done()

    var pwd_script string
    pwd_script = "~/dev/Defrauder/Tools/dnscan/dnscan.py"
    cmd := exec.Command("bash", "-c", fmt.Sprintf("python3 %s -d %s -n >> .tmp/on_domain.txt", pwd_script, domain))
    cmd.Run()

}

func CreateTemporaryDirectory() {

	if err := os.MkdirAll(".tmp", os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
	}

}

func Rm_extra(){

	cmd8 := exec.Command("bash", "-c", "rm -r .tmp")
	cmd8.Run()

}

func alphabetMaker() {

	cmd := exec.Command("bash", "-c", "mkdir -p .tmp")
	cmd.Run()

	file, err := os.Create(".tmp/letters_output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < 100000; i++ {
		char := rune(i)

		if unicode.IsLetter(char) && unicode.Is(unicode.Latin, char) {
			_, err := writer.WriteString(fmt.Sprintf("%c\n", char))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing data to file:", err)
	}
	cmd1 := exec.Command("bash", "-c", "sort -u .tmp/letters_output.txt -o .tmp/sorted_alp.txt")
	cmd1.Run()
	wg.Done() 

}

func request(domain string, idx int, wg *sync.WaitGroup) {

    defer wg.Done()

    fileName := ".tmp/sorted_alp.txt"
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var alphabetArr []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        alphabetArr = append(alphabetArr, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    outputFile1 := ".tmp/domain_list.txt"
    file, err = os.OpenFile(outputFile1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    for _, char := range alphabetArr {
        newDomain := domain[:idx] + char + domain[idx+1:]
        

        writer.WriteString(newDomain + "\n")
    }
    writer.Flush() 

}

func splitDomain(domain string,base string, ext string) {

    var wg sync.WaitGroup

   
    domainLength := len(base) 

    for i := 0; i < domainLength; i++ {
        wg.Add(1)
        go request(domain, i, &wg)
    }

    wg.Wait() 
    
}

func GenerateVariations(baseWord string) []string {

	variations := map[rune][]string{
		'f': {"f", "fa", "fc", "ff"},
		'a': {"a", "aa", "ae", "o", "4", "@"},
		'c': {"c", "cc", "ck"},
		'e': {"e", "ee", "3"},
		'b': {"b", "bb", "p", "d", "9"},
		'o': {"o", "0", "oo"},
		'k': {"k", "kk", "q", "ck"},
		'd': {"d", "dd"},
		'g': {"g", "9"},
		'h': {"h", "4", "hh"},
		'i': {"i", "1", "l", "!", "|"},
		'l': {"l", "1", "|", "i"},
		'm': {"m", "nn", "n"},
		'n': {"n", "nn"},
		'p': {"p", "pp", "9"},
		'r': {"r", "rr"},
		's': {"s", "ss", "5", "$"},
		't': {"t", "tt", "7"},
		'u': {"u", "v", "ù"},
		'y': {"y", "ÿ"},
	}

	var splitWord [][]string
	for _, char := range baseWord {
		if v, found := variations[char]; found {
			splitWord = append(splitWord, v)
		} else {
			splitWord = append(splitWord, []string{string(char)})
		}
	}

	var results []string
	combine(splitWord, "", &results)
	return results
}

func combine(chars [][]string, current string, results *[]string) {
	if len(chars) == 0 {
		*results = append(*results, current)
		return
	}
	for _, char := range chars[0] {
		combine(chars[1:], current+char, results)
	}
}

func method2(SLD string, TLD string) {
    wordlist := GenerateVariations(SLD)

    outputFile := ".tmp/fake_domain_wordlist.txt"
    file, err := os.Create(outputFile)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    for _, word := range wordlist {
        tmp := word + "." + TLD
        writer.WriteString(tmp + "\n")
    }
    writer.Flush()
}

func check_live(stopChan chan struct{}){
    
		var innerWg sync.WaitGroup
	
		cmd:=exec.Command("bash","-c","cat .tmp/domain_list.txt .tmp/fake_domain_wordlist.txt > .tmp/combined.txt ")
		cmd.Run()

		file, err := os.Open(".tmp/combined.txt")
		if err != nil {
			fmt.Printf("failed to open file: %s", err)
		}
		defer file.Close()
	
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			innerWg.Add(1)
			go func(d string) {
				defer innerWg.Done()
				
				limitChan <- struct{}{} // Acquire a spot
				isDomainLive(d)
				<-limitChan // Release the spot
			}(line)

		}
		innerWg.Wait() 
	
		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading file: %s", err)
		}
		close(stopChan)
}
func showData(stopChan <-chan struct{}) {
		
	for {
		select {
		case <-stopChan:
			fmt.Println("Stopping data display...")
			return
		default:
			
			clearTerminal()
			banner()
		  
	        cmd0 := exec.Command("bash", "-c", `cat .tmp/on_domain.txt | sed 's/\x1b\[K//g' | grep -E "([0-9]{1,3}\.){3}[0-9]{1,3} - ([a-zA-Z0-9_-]+\.)+[a-zA-Z]{2,}" | sort -u -o .tmp/sorted_output_data.txt`)
	        cmd0.Run()
			
			output, err := os.ReadFile(".tmp/sorted_output_data.txt")
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
				
			outputLines := strings.Split(string(output), "\n")
				
			for _, line := range outputLines {
				if line != "" {
					fmt.Println(line)
				}
			}
				
			time.Sleep(5 * time.Second)
		}
	}
}

func main() {
    var domain, bs, outputFileName string
    var bufferSize int = 32 // default buffer size
    var mainWg sync.WaitGroup
    stopChan := make(chan struct{})
    
    Rm_extra()
    CreateTemporaryDirectory()
    
    
    flagString := os.Args[1:]
    if len(flagString) < 2 {
        fmt.Println("ENTER THE CORRECT COMMAND...")
        displayHelp()
        return
    }
    
    for i := 0; i < len(flagString); i++ {
        switch flagString[i] {
        case "-d":
            if i+1 < len(flagString) {
                domain = flagString[i+1]
                i++
            }
        case "-o":
            if i+1 < len(flagString) {
                outputFileName = flagString[i+1]
                i++
            }
        case "-t":
            if i+1 < len(flagString) {
                bs = flagString[i+1]
                var err error
                bufferSize, err = strconv.Atoi(bs)
                if err != nil {
                    fmt.Println("Invalid buffer size; using default value 32")
                    bufferSize = 32
                }
                i++
            }
        case "-h":
            displayHelp()
            return
        }
    }

    
    if domain == "" {
        fmt.Println("Error: Domain is required")
        displayHelp()
        return
    }

    
    initLimitChan(bufferSize)
    
    mainWg.Add(1)
    go func() {
        defer mainWg.Done()
        defer func() {
            if r := recover(); r != nil {
                // fmt.Printf("Recovered from panic in alphabetMaker: %v\n", r)
            }
        }()
        alphabetMaker()
    }()
    mainWg.Wait()
    go showData(stopChan)
    tldSize := -1
    for i := 0; i < len(domain); i++ {
        if domain[i] == '.' {
            tldSize = i
            break
        }
    }
    
    if tldSize == -1 {
        fmt.Println("Error: Invalid domain format. Domain must contain a TLD (e.g., example.com)")
        return
    }
        parts := strings.SplitN(domain, ".", 2)
        if len(parts) == 2 {
            base := parts[0]
            ext := parts[1]

            // - MET 1 -
            splitDomain(domain, base, ext)
            // - MET 2 -
            method2(base, ext)
            check_live(stopChan)


        } else {
            fmt.Println("Error: Invalid domain format. Domain must contain a TLD (e.g., example.com)")
            return
    	}

    // Handle output file if specified
    if outputFileName != "" {
        err := func() error {
            cmd := exec.Command("bash", "-c", fmt.Sprintf(`cat .tmp/on_domain.txt | sed 's/\x1b\[K//g' | grep -E "([0-9]{1,3}\.){3}[0-9]{1,3} - ([a-zA-Z0-9_-]+\.)+[a-zA-Z]{2,}" > %s`, outputFileName))
            if err := cmd.Run(); err != nil {
                return fmt.Errorf("error saving results: %v", err)
            }
            fmt.Printf("Results saved to %s\n", outputFileName)
            return nil
        }()

        if err != nil {
            fmt.Println(err)
        }
    }

    fmt.Println("\nDomain processing completed successfully!")
}
