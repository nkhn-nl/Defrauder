#!/bin/bash
GREEN="\e[32m"
BLUE="\e[34m"
NC="\e[0m" # No Color

# Banner
echo "Setting up Defrauder environment..."

# Check for necessary tools
echo "[+] Checking for required dependencies..."

# Install Python and pip if not already installed
if ! command -v python3 &> /dev/null; then
    echo "[-] Python3 not found. Installing..."
    sudo apt update && sudo apt install -y python3
else
    echo "[+] Python3 is already installed."
fi

if ! command -v pip3 &> /dev/null; then
    echo "[-] pip3 not found. Installing..."
    sudo apt install -y python3-pip
else
    echo "[+] pip3 is already installed."
fi

# Install necessary Go tools
echo "[+] Checking for Go installation..."
if ! command -v go &> /dev/null; then
    echo "[-] Go not found. Please install Go and try again."
    exit 1
fi

# Install httpx using Go
echo "[+] Installing httpx..."
go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest
if [ $? -eq 0 ]; then
    echo "[+] httpx installed successfully."
else
    echo "[-] Failed to install httpx. Please check your Go environment."
    exit 1
fi

# Install figlet and lolcat for banner
echo "[+] Installing figlet and lolcat..."
if ! command -v figlet &> /dev/null; then
    sudo apt install -y figlet
fi

if ! command -v lolcat &> /dev/null; then
    sudo apt install -y lolcat
fi

# Final message
sudo go build Defrauder.go 
sudo mv Defrauder /usr/local/bin

# Final message
echo -e "${GREEN}Setup completed! You can now run the tool using:${NC}"
echo -e "${BLUE}Defrauder -d example.com -o output.txt ${NC}"
