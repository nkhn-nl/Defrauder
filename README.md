# Defrauder - Fake Domain Detection Tool


**Defrauder** is a cybersecurity tool designed to identify fake or typo-squatting domains that could be used for phishing or scams. It generates domain variations based on a given target domain and checks if these domains are live.

## Features

1. **Domain Variation Generation**:
   - **Method 1**: Generates variations of the second-level domain (SLD) by replacing characters with visually or typographically similar characters (e.g., replacing 'o' with '0', 'i' with '1').

2. **Domain Scanning**:
   - For each generated domain variation, Defrauder checks if the domain is live using  (`httpx`).

3. **Concurrent Domain Checking**:
   - Defrauder supports multi-threaded domain checking with a configurable buffer size, allowing you to specify how many domain checks can run concurrently.

4. **Live Data Display**:
   - The tool continuously updates the terminal with the live domains it has discovered, displaying them in real-time as they are found.

5. **Customizable Output**:
   - You can specify an output file where the discovered live domains will be saved.

## Installation

### Prerequisites

- **Go** (version 1.15+)
- **Python3** (for running `dnscan.py`)
- **bash** (Unix/Linux systems)

Install dependencies using:

```
sudo apt-get install figlet lolcat
sudo apt-get install python3
```

### Clone the Repository

```
git clone https://github.com/Yashhackz7721/defrauder.git
cd defrauder
chmod +x setup.sh
sudo ./setup.sh
```

### Setup

Ensure you have the necessary permissions to execute scripts, and `httpx` is correctly set up.

## Usage

To run the tool, use the following command structure:

```
defrauder -d <domain> -o <output_file> -t <buffer_size> -fc 302,301
```

### Flags:

- `-d <domain>`: The target domain to check for fake domains.
- `-o <output_file>`: Specify the output file for saving results (optional).
- `-t <buffer_size>`: Set the buffer size for concurrent domain checks (optional, default is 32).
- `-fc <filter status>`: filter response with specified status code (-fc 403,401)
- `-h`: Displays help for usage.

### Example:

```
defrauder -d example.com -o results.txt -t 40
```

This command will generate domain variations for `example.com`, check which ones are live, and save the results to `defrauder_result.txt`.

## How It Works

1. **Alphabet Generation**:
   - Defrauder generates a set of characters (A-Z, a-z) that will be used to create domain variations.
   - Uses a custom algorithm to generate domain variations by altering characters in different positions of the SLD.
2. **Domain Splitting**:
   - It splits the input domain into the base (SLD) and extension (TLD). For example, `example.com` becomes `example` (SLD) and `com` (TLD).
3. **Fake Domain Generation**:
   - Defrauder generates possible fake domains by modifying characters in the SLD (e.g., `examp1e.com`, `exampl3.com`, etc.).
4. **Live Domain Checking**:
   - Each generated domain is checked to see if it's live using the `httpx` script.
5. **Real-time Updates**:
   - As live domains are found, they are displayed in real-time in the terminal.
6. **Final Results**:
   - After processing, the live domains are saved in the specified output file.

## Example Output

In the terminal, you will see a banner and the list of live domains:

```
 _____        __                     _
|  __ \      / _|                   | |
| |  | | ___| |_ _ __ __ _ _   _  __| | ___ _ __
| |  | |/ _ \  _| '__/ _` | | | |/ _` |/ _ \ '__|
| |__| |  __/ | | | | (_| | |_| | (_| |  __/ |
|_____/ \___|_| |_|  \__,_|\__,_|\__,_|\___|_|


[+] Running custom algorithm to alter characters at different positions.
[+] Live Check: Verifies which generated domains are currently active.
--------------------------------------------------------------------------

[-] Generation varation of the Doamin : example.com
[-] Total number of domain generated: 6192148
[-] Output will be saved at : result.txt
 [-] CHECKING FOR LIVE DOMAIN.
https://EXamplé.com [308]
https://EXamPle.com [200]
https://EXample.com [200]
https://EXaMple.com [200]
https://EXamplE.com [200]
https://EXAmple.com [200]
```

