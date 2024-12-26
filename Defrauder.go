package main

import (
	"os"
	"fmt"
	"regexp"
	"os/exec"
	"strings"
	"io/ioutil"
)

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

	fmt.Println("[+] Running custom algorithm to alter characters at different positions. ")
	fmt.Println("[+] Live Check: Verifies which generated domains are currently active. ")
	fmt.Println("--------------------------------------------------------------------------")
}

func displayHelp() {
    fmt.Println("Usage:  Defrauder.go -d <domain> -o <output_file> -t <buffer_size>")
    fmt.Println("\nFlags:")
    fmt.Println("  -d  <domain>       Target domain to check for fakes.")
    fmt.Println("  -o  <output file>  Specify the output file for results.")
    fmt.Println("  -t  <buffer size>  Set the buffer size for concurrent checks (default is 50).")
    fmt.Println("  -fc <filter status>  filter response with specified status code (-fc 403,401)")
    fmt.Println("\nExample:")
    fmt.Println(" Defrauder.go -d example.com -o results.txt -t 60 -fc 403,401")
}

func generateTmpFile() error {
	// Create the .tmp directory if it doesn't exist
	cmd := exec.Command("mkdir", "-p", ".tmp")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error creating .tmp directory: %v", err)
	}

	return nil
}

// Function to remove the .tmp file using exec.Command()
func removeTmpFile() error {

	// Optionally, remove the .tmp directory if it's empty
	cmd8 := exec.Command("bash", "-c", "rm -r .tmp")
	cmd8.Run()

	return nil
}

func runPythonScript(domain string) error {
	pythonCode := `
import itertools


def generate_variations(domain, char_variations, replace_count, output_file):
    domain_chars = list(domain)
    domain_length = len(domain_chars)

    replace_positions_combinations = itertools.combinations(range(domain_length), replace_count)

    with open(output_file, 'w') as file:
        for positions in replace_positions_combinations:
            
            replacement_options = []
            for index, char in enumerate(domain_chars):
                if index in positions:
                    replacement_options.append(char_variations.get(char, [char]))
                else:
                    replacement_options.append([char])

            for variation in itertools.product(*replacement_options):
                file.write(''.join(variation) + '\n')

if __name__ == "__main__":
    
    char_variations = {
       'a': ['a','ａ','A','Ａ','ª','ᵃ','ₐ','ᴬ','á','Á','à','À','ă','Ă','ắ','Ắ','ằ','Ằ','ẵ','Ẵ','ẳ','Ẳ','â','Â','ấ','Ấ','ầ','Ầ','ẫ','Ẫ','ẩ','Ẩ','ǎ','Ǎ','å','Å','Å','ǻ','Ǻ','ä','ꞛ','Ä','Ꞛ','ǟ','Ǟ','ã','Ã','ȧ','Ȧ','ǡ','Ǡ','ą','Ą','ā','Ā','ả','Ả','ȁ','Ȁ','ȃ','Ȃ','ạ','Ạ','ặ','Ặ','ậ','Ậ','ḁ','Ḁ','ꜳ','Ꜳ','æ','Æ','ᴭ','ǽ','Ǽ','ǣ','Ǣ','ꜵ','Ꜵ','ꜷ','Ꜷ','ꜹ','Ꜹ','ꜻ','Ꜻ','ꜽ','Ꜽ','ẚ','ᴀ','ⱥ','Ⱥ','ᶏ','ᴁ','ᴂ','ᵆ','ꬱ','ɐ','Ɐ','ᵄ','ɑ','Ɑ','ᵅ','ꬰ','ᶐ','ɒ','Ɒ','ᶛ','ꭤ'],
		'b': ['b','ｂ','B','Ｂ','ᵇ','ᴮ','ḃ','Ḃ','ḅ','Ḅ','ḇ','Ḇ','ʙ','ƀ','Ƀ','ᴯ','ᴃ','ᵬ','ꞗ','Ꞗ','ᶀ','ɓ','Ɓ','ƃ','Ƃ','ꞵ','Ꞵ'],
		'c': ['c','ｃ','C','Ｃ','ᶜ','ć','Ć','ĉ','Ĉ','č','Č','ċ','Ċ','ç','Ç','ḉ','Ḉ','ᴄ','ȼ','Ȼ','ꞓ','Ꞓ','ꞔ','ƈ','Ƈ','ɕ','ᶝ','ↄ','Ↄ','ꜿ','Ꜿ'],
		'd': ['d','ｄ','D','Ｄ','ᵈ','ᴰ','ď','Ď','ḋ','Ḋ','ḑ','Ḑ','đ','Đ','ḍ','Ḍ','ḓ','Ḓ','ḏ','Ḏ','ð','Ð','ᶞ','ꝺ','Ꝺ','ȸ','ǳ','ʣ','ǲ','Ǳ','ǆ','ǅ','Ǆ','ʥ','ʤ','ᴅ','ᴆ','ᵭ','ᶁ','ɖ','Ɖ','ɗ','Ɗ','ᶑ','ƌ','Ƌ','ȡ','ꝱ','ẟ'],
		'e': ['e','ｅ','E','Ｅ','ᵉ','ₑ','ᴱ','é','É','è','È','ĕ','Ĕ','ê','Ê','ế','Ế','ề','Ề','ễ','Ễ','ể','Ể','ě','Ě','ë','Ë','ẽ','Ẽ','ė','Ė','ȩ','Ȩ','ḝ','Ḝ','ę','Ę','ē','Ē','ḗ','Ḗ','ḕ','Ḕ','ẻ','Ẻ','ȅ','Ȅ','ȇ','Ȇ','ẹ','Ẹ','ệ','Ệ','ḙ','Ḙ','ḛ','Ḛ','ᴇ','ꬲ','ꬳ','ɇ','Ɇ','ᶒ','ꬴ','ⱸ','ǝ','Ǝ','ᴲ','ⱻ','ə','Ə','ᵊ','ₔ','ᶕ','ɛ','Ɛ','ᵋ','ᶓ','ɘ','ɚ','ɜ','Ɜ','ᶟ','ᶔ','ᴈ','ᵌ','ɝ','ɞ','ʚ','ɤ'],
		'f': ['f','ｆ','F','Ｆ','ᶠ','ḟ','Ḟ','ꝼ','Ꝼ','ﬀ','ﬃ','ﬄ','ﬁ','ﬂ','ʩ','ꜰ','ꬵ','ꞙ','Ꞙ','ᵮ','ᶂ','ƒ','Ƒ','ⅎ','Ⅎ','ꟻ'],
		'g': ['g','ｇ','G','Ｇ','ᵍ','ᴳ','ǵ','Ǵ','ğ','Ğ','ĝ','Ĝ','ǧ','Ǧ','ġ','Ġ','ģ','Ģ','ḡ','Ḡ','ꞡ','Ꞡ','ᵹ','Ᵹ','ɡ','Ɡ','ᶢ','ꬶ','ɢ','ǥ','Ǥ','ᶃ','ɠ','Ɠ','ʛ','ᵷ','ꝿ','Ꝿ','ɣ','Ɣ','ˠ','ƣ','Ƣ'],
		'h': ['h','ｈ','H','Ｈ','ʰ','ₕ','ᴴ','ĥ','Ĥ','ȟ','Ȟ','ḧ','Ḧ','ḣ','Ḣ','ḩ','Ḩ','ħ','Ħ','ꟸ','ḥ','Ḥ','ḫ','Ḫ','ẖ','ʜ','ƕ','Ƕ','ꞕ','ɦ','Ɦ','ʱ','ⱨ','Ⱨ','ⱶ','Ⱶ','ꜧ','Ꜧ','ꭜ','ɧ'],
		'i': ['i','ｉ','I','Ｉ','ⁱ','ᵢ','ᴵ','í','Í','ì','Ì','ĭ','Ĭ','î','Î','ǐ','Ǐ','ï','Ï','ḯ','Ḯ','ĩ','Ĩ','İ','į','Į','ī','Ī','ỉ','Ỉ','ȉ','Ȉ','ȋ','Ȋ','ị','Ị','ḭ','Ḭ','ĳ','Ĳ','ı','ɪ','Ɪ','ᶦ','ꟾ','ꟷ','ᴉ','ᵎ','ɨ','Ɨ','ᶤ','ᵻ','ᶧ','ᶖ','ɩ','Ɩ'],
		'j': ['ᶥ','ᵼ','j','ｊ','J','Ｊ','ʲ','ⱼ','ᴶ','ĵ','Ĵ','ǰ','ȷ','ᴊ','ɉ','Ɉ','ʝ','Ʝ','ᶨ','ɟ','ᶡ','ʄ'],
		'k': ['k','ｋ','K','K','Ｋ','ᵏ','ₖ','ᴷ','ḱ','Ḱ','ǩ','Ǩ','ķ','Ķ','ꞣ','Ꞣ','ḳ','Ḳ','ḵ','Ḵ','ᴋ','ᶄ','ƙ','Ƙ','ⱪ','Ⱪ','ꝁ','Ꝁ','ꝃ','Ꝃ','ꝅ','Ꝅ','ʞ','Ʞ'],
		'l': ['l','ｌ','L','Ｌ','ˡ','ₗ','ᴸ','ĺ','Ĺ','ľ','Ľ','ļ','Ļ','ł','Ł','ḷ','Ḷ','ḹ','Ḹ','ḽ','Ḽ','ḻ','Ḻ','ŀ','Ŀ','ǉ','ǈ','Ǉ','ỻ','Ỻ','ʪ','ʫ','ʟ','ᶫ','ꝇ','Ꝇ','ᴌ','ꝉ','Ꝉ','ƚ','Ƚ','ⱡ','Ⱡ','ɫ','Ɫ','ꭞ','ꬸ','ꬹ','ɬ','Ɬ','ꬷ','ꭝ','ᶅ','ᶪ','ɭ','ᶩ','ꞎ','ȴ','ꝲ','ɮ','ꞁ','Ꞁ'],
		'm': ['m','ｍ','M','Ｍ','ᵐ','ₘ','ᴹ','ḿ','Ḿ','ṁ','Ṁ','ṃ','Ṃ','ᴍ','ᵯ','ᶆ','ɱ','Ɱ','ᶬ','ꬺ','ꟽ','ꟿ','ꝳ'],
		'n': ['n','ｎ','N','Ｎ','ⁿ','ₙ','ᴺ','ń','Ń','ǹ','Ǹ','ň','Ň','ñ','Ñ','ṅ','Ṅ','ņ','Ņ','ꞥ','Ꞥ','ṇ','Ṇ','ṋ','Ṋ','ṉ','Ṉ','ǌ','ǋ','Ǌ','ɴ','ᶰ','ᴻ','ᴎ','ᵰ','ɲ','Ɲ','ᶮ','ƞ','Ƞ','ꞑ','Ꞑ','ᶇ','ɳ','ᶯ','ȵ','ꬻ','ꝴ','ŋ','Ŋ','ᵑ','ꬼ'],
		'o': ['o','ｏ','O','Ｏ','º','ᵒ','ₒ','ᴼ','ó','Ó','ò','Ò','ŏ','Ŏ','ô','Ô','ố','Ố','ồ','Ồ','ỗ','Ỗ','ổ','Ổ','ǒ','Ǒ','ö','ꞝ','Ö','Ꞝ','ȫ','Ȫ','ő','Ő','õ','Õ','ṍ','Ṍ','ṏ','Ṏ','ȭ','Ȭ','ȯ','Ȯ','ȱ','Ȱ','ø','Ø','ǿ','Ǿ','ǫ','Ǫ','ǭ','Ǭ','ō','Ō','ṓ','Ṓ','ṑ','Ṑ','ỏ','Ỏ','ȍ','Ȍ','ȏ','Ȏ','ơ','Ơ','ớ','Ớ','ờ','Ờ','ỡ','Ỡ','ở','Ở','ợ','Ợ','ọ','Ọ','ộ','Ộ','œ','Œ','ꟹ','ꝏ','Ꝏ','ᴏ','ᴑ','ꬽ','ɶ','ᴔ','ꭁ','ꭂ','ꭀ','ꭃ','ꭄ','ᴓ','ꬾ','ɔ','Ɔ','ᵓ','ᴐ','ᴒ','ꬿ','ᶗ','ꭢ','ꝍ','Ꝍ','ᴖ','ᵔ','ᴗ','ᵕ','ⱺ','ɵ','Ɵ','ᶱ','ꝋ','Ꝋ','ɷ'],
		'p': ['p','ｐ','P','Ｐ','ᵖ','ₚ','ᴾ','ṕ','Ṕ','ṗ','Ṗ','ᴘ','ᵽ','Ᵽ','ꝑ','Ꝑ','ᵱ','ᶈ','ƥ','Ƥ','ꝓ','Ꝓ','ꝕ','Ꝕ','ꟼ'],
		'q': ['ⱷ','q','ｑ','Q','Ｑ','ȹ','ꝗ','Ꝗ','ꝙ','Ꝙ','ʠ','ɋ','Ɋ'],
		'r': ['r','ｒ','R','Ｒ','ʳ','ᵣ','ᴿ','ŕ','Ŕ','ř','Ř','ṙ','Ṙ','ŗ','Ŗ','ꞧ','Ꞧ','ȑ','Ȑ','ȓ','Ȓ','ṛ','Ṛ','ṝ','Ṝ','ṟ','Ṟ','ꞃ','Ꞃ','ꭅ','ʀ','Ʀ','ꭆ','ꝛ','Ꝛ','ᴙ','ɍ','Ɍ','ᵲ','ɹ','ʴ','ᴚ','ɺ','ᶉ','ɻ','ʵ','ⱹ','ɼ','ɽ','Ɽ','ꭉ','ɾ','ᵳ','ɿ','ꭇ','ꭈ','ꭊ','ꭋ','ꭌ','ʁ','ʶ','ꝵ','ꝶ'],
		's': ['s','ｓ','S','Ｓ','ˢ','ₛ','ś','Ś','ṥ','Ṥ','ŝ','Ŝ','š','Š','ṧ','Ṧ','ṡ','Ṡ','ş','Ş','ꞩ','Ꞩ','ṣ','Ṣ','ṩ','Ṩ','ș','Ș','ſ','ꞅ','Ꞅ','ẛ','ß','ẞ','ﬆ','ﬅ','ꜱ','ᵴ','ᶊ','ʂ','ᶳ','ȿ','Ȿ'],
		't': ['t','ｔ','T','Ｔ','ᵗ','ₜ','ᵀ','ť','Ť','ẗ','ṫ','Ṫ','ţ','Ţ','ṭ','Ṭ','ț','Ț','ṱ','Ṱ','ṯ','Ṯ','ꞇ','Ꞇ','ʨ','ᵺ','ƾ','ʦ','ʧ','ꜩ','Ꜩ','ᴛ','ŧ','Ŧ','ⱦ','Ⱦ','ᵵ','ƫ','ᶵ','ƭ','Ƭ','ʈ','Ʈ','ȶ','ꝷ','ʇ','Ʇ'],
		'u': ['ｕ','U','Ｕ','ᵘ','ᵤ','ᵁ','ú','Ú','ù','Ù','ŭ','Ŭ','û','Û','ǔ','Ǔ','ů','Ů','ü','ꞟ','Ü','Ꞟ','ǘ','Ǘ','ǜ','Ǜ','ǚ','Ǚ','ǖ','Ǖ','ű','Ű','ũ','Ũ','ṹ','Ṹ','ų','Ų','ū','Ū','ṻ','Ṻ','ủ','Ủ','ȕ','Ȕ','ȗ','Ȗ','ư','Ư','ứ','Ứ','ừ','Ừ','ữ','Ữ','ử','Ử','ự','Ự','ụ','Ụ','ṳ','Ṳ','ṷ','Ṷ','ṵ','Ṵ','ᴜ','ᶸ','ꭎ','ᴝ','ᵙ','ᴞ','ᵫ','ꭐ','ꭑ','ʉ','Ʉ','ᶶ','ꭏ','ᵾ','ᶙ','ꭒ','ꭟ','ɥ','Ɥ','ᶣ','ʮ','ʯ','ʊ','Ʊ','ᶷ','ᵿ'],
		'v': ['v','ｖ','V','Ｖ','ᵛ','ᵥ','ⱽ','ṽ','Ṽ','ṿ','Ṿ','ꝡ','Ꝡ','ᴠ','ꝟ','Ꝟ','ᶌ','ʋ','Ʋ','ᶹ','ⱱ','ⱴ','ỽ','Ỽ','ʌ','Ʌ','ᶺ'],
		'w': ['w','ｗ','W','Ｗ','ʷ','ᵂ','ẃ','Ẃ','ẁ','Ẁ','ŵ','Ŵ','ẘ','ẅ','Ẅ','ẇ','Ẇ','ẉ','Ẉ','ᴡ','ⱳ','Ⱳ','ʍ'],
		'x': ['x','ｘ','X','Ｘ','ˣ','ₓ','ẍ','Ẍ','ẋ','Ẋ','ᶍ','ꭖ','ꭗ','ꭘ','ꭙ','ꭓ','Ꭓ','ꭔ','ꭕ'],
		'y': ['y','ｙ','Y','Ｙ','ʸ','ý','Ý','ỳ','Ỳ','ŷ','Ŷ','ẙ','ÿ','Ÿ','ỹ','Ỹ','ẏ','Ẏ','ȳ','Ȳ','ỷ','Ỷ','ỵ','Ỵ','ʏ','ɏ','Ɏ','ƴ','Ƴ','ỿ','Ỿ','ꭚ'],
		'z': ['z','ｚ','Z','Ｚ','ᶻ','ź','Ź','ẑ','Ẑ','ž','Ž','ż','Ż','ẓ','Ẓ','ẕ','Ẕ','ƍ','ᴢ','ƶ','Ƶ','ᵶ','ᶎ','ȥ','Ȥ','ʐ','ᶼ','ʑ','ᶽ','ɀ','Ɀ','ⱬ','Ⱬ']
    }
    	
def process_domain(domain):
    # Separate the TLD
    if '.' in domain:
        name, tld = domain.rsplit('.', 1)  # Split from the last dot
        return name,tld
    else:
        return 

def append_tld_to_file(input_file, tld, output_file):
    try:
        with open(input_file, "r") as infile, open(output_file, "w") as outfile:
            for line in infile:
                line = line.strip()  
                if line:  
                    outfile.write(f"{line}.{tld}\n")
    except FileNotFoundError:
        print(f"Error: {input_file} not found.")
    except Exception as e:
        print(f"An error occurred: {e}")
        
domain = "`+domain+`".strip().lower()	
name,tld=process_domain(domain)

if len(name) >= 6:
    replace_count = 3
elif 3 < len(name) <= 6:
    replace_count = 2
else:
    replace_count = 1
output_file = ".tmp/var.txt"
output_file2 = ".tmp/var2.txt"

generate_variations(name, char_variations, replace_count, output_file)
append_tld_to_file(output_file,tld,output_file2)
	`

	tempFile, err := os.CreateTemp("", "*.py")
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(pythonCode)
	if err != nil {
		return fmt.Errorf("error writing to temp file: %w", err)
	}

	cmd := exec.Command("python3", tempFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing Python code: %w\nOutput: %s", err, output)
	}

	return nil
}
func count(){

	err := os.Remove(".tmp/var.txt")
	if err != nil {
		fmt.Printf("Error deleting .tmp/var.txt: %v\n", err)
	}

	cmd1 := exec.Command("bash", "-c", "cat .tmp/var2.txt | wc -l")
	output, err := cmd1.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("[-] Total number of domain generated:\033[33m %s\033[0m", string(output))
	}

}


func check_live_with_filter(threads string, filterStatusCodes []string) {
	
	filterArg := strings.Join(filterStatusCodes, ",")

	
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cat .tmp/var2.txt | httpx -fc %s -sc -silent -threads %s -o .tmp/on_data.txt", filterArg, threads))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}

func removeColorCodes(content string) string {
	
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(content, "")

}

func move_output_file(file string) {
	
	content, err := ioutil.ReadFile(".tmp/on_data.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	
	cleanedContent := removeColorCodes(string(content))

	
	err = ioutil.WriteFile(".tmp/on_data.txt", []byte(cleanedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing cleaned content to file: %v\n", err)
		return
	}

	
	cmd := exec.Command("mv", ".tmp/on_data.txt", file)

	
	err = cmd.Run()
	if err != nil {
		
		fmt.Printf("Error moving file: %v\n", err)
	} else {
		fmt.Printf("File moved successfully to: %s\n", file)
	}
}

func main() {
        var Reset = "\033[0m"
        var Red = "\033[31m"

	var domain, output string
	var threads string
	threads = "50"
	output = "defrauder_result.txt"

	filterStatusCodes := []string{
		"301", "400", "401", "403", "404", "405", "408", "410",
		"413", "414", "429", "451", "500", "501", "502", "503",
		"504", "511", "444", "499",
	}

	flagString := os.Args[1:]
	if len(flagString) == 1 && flagString[0] == "-h" {
		displayHelp()
		return
	}

	if len(flagString) < 2 {
		fmt.Println(Red + "Fault: not enough parameters" + Reset + "\n")
		displayHelp()
		return
	}

	banner()

	for i := 0; i < len(flagString); i++ {
		switch flagString[i] {
		case "-d":
			if i+1 < len(flagString) {
				domain = flagString[i+1]
			}
		case "-o":
			if i+1 < len(flagString) {
				output = flagString[i+1]
			}
		case "-t":
			if i+1 < len(flagString) {
				threads = flagString[i+1]
			}
		case "-fc":
			if i+1 < len(flagString) {
				userCodes := strings.Split(flagString[i+1], ",")
				validCodes := filterValidStatusCodes(userCodes)
				filterStatusCodes = append(filterStatusCodes, validCodes...)
			}
		case "-h":
			displayHelp()
			return
		}
	}

	generateTmpFile()
	fmt.Printf("\n[-] Generating variations of the Domain: \033[92m%s\033[0m\n", domain)
	runPythonScript(domain)

	count()
	fmt.Printf("[-] Output will be saved at: %s\n", output)

	fmt.Printf("\033[36m[-] CHECKING FOR LIVE DOMAIN.\033[0m\n")
	check_live_with_filter(threads, filterStatusCodes)

	move_output_file(output)
	removeTmpFile()
}

func filterValidStatusCodes(codes []string) []string {
	var validCodes []string
	regex := regexp.MustCompile(`^\d{3}$`)
	for _, code := range codes {
		if regex.MatchString(code) {
			validCodes = append(validCodes, code)
		} else {
			fmt.Printf("[-] Invalid status code ignored: %s\n", code)
		}
	}
	return validCodes
}
