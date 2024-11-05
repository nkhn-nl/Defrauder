#!/bin/bash

# Defining colors for output
GREEN="\e[32m"
BLUE="\e[34m"
RED="\e[31m"
NC="\e[0m" # No Color

# Print Banner
echo -e "${GREEN}"
figlet -f big "Setup Script"
echo -e "${NC}"

# Update and install dependencies
echo -e "${GREEN}Updating and installing required packages...${NC}"
sudo apt-get update

# Check and install Python3
if ! command -v python3 &> /dev/null; then
    echo -e "${RED}Python3 is not installed. Installing...${NC}"
    sudo apt-get install python3 -y
else
    echo -e "${GREEN}Python3 is already installed.${NC}"
fi

# Check and install required Python packages
echo -e "${GREEN}Checking for required Python packages...${NC}"
pip3 install dnspython

# Install figlet
if ! command -v figlet &> /dev/null; then
    echo -e "${RED}Figlet is not installed. Installing...${NC}"
    sudo apt-get install figlet -y
else
    echo -e "${GREEN}Figlet is already installed.${NC}"
fi

# Install lolcat
if ! command -v lolcat &> /dev/null; then
    echo -e "${RED}Lolcat is not installed. Installing...${NC}"
    sudo apt-get install lolcat -y
else
    echo -e "${GREEN}Lolcat is already installed.${NC}"
fi

# Check and install git
if ! command -v git &> /dev/null; then
    echo -e "${RED}Git is not installed. Installing...${NC}"
    sudo apt-get install git -y
else
    echo -e "${GREEN}Git is already installed.${NC}"
fi

# Define the current directory using pwd
CURRENT_DIR=$(pwd)

# Download and setup dnscan tool
if [ ! -f "$CURRENT_DIR/Tools/dnscan/dnscan.py" ]; then
    echo -e "${GREEN}Downloading dnscan tool...${NC}"
    
    # Create Tools directory if it doesn't exist
    mkdir -p "$CURRENT_DIR/Tools"

    # Clone dnscan repository to Tools directory
    git clone https://github.com/rbsec/dnscan "$CURRENT_DIR/Tools/dnscan"
    
    # Get the absolute path of dnscan.py
    DN_SCAN_PATH="$CURRENT_DIR/Tools/dnscan/dnscan.py"

    # Replace the old path in Defrauder.go with the new one
    sed -i "s|pwd_script = \".*\"|pwd_script = \"$DN_SCAN_PATH\"|" Defrauder.go

    echo -e "${GREEN}Updated Defrauder.go with the new dnscan path: $DN_SCAN_PATH${NC}"
else
    echo -e "${GREEN}dnscan is already downloaded.${NC}"
fi

# Ensure Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Go is not installed. Installing...${NC}"
    sudo apt-get install golang -y
else
    echo -e "${GREEN}Go is already installed.${NC}"
fi

# build bin
sudo go build Defrauder.go 
sudo mv Defrauder /usr/local/bin

# Final message
echo -e "${GREEN}Setup completed! You can now run the tool using:${NC}"
echo -e "${BLUE}Defrauder -d example.com -o output.txt ${NC}"
