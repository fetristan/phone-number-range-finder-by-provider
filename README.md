# Phone Number Range Finder By Provider

## Overview
This Go project is designed to process a CSV file containing phone numbers and their associated providers. It identifies ranges of consecutive phone numbers for each provider and outputs the results into a new CSV file. The output includes the start and end of each range, the provider's name, and the length of the range.

## Features
- Reads phone numbers and providers from a CSV file.
- Identifies consecutive number ranges for each provider.
- Outputs the ranges to a new CSV file with range details.
- Handles large datasets efficiently.

## Requirements
- Go (version 1.15 or later) if you want to run the program from source.

## Installation
If you want to run the program from source:
- Clone the repository to your local machine:
  git clone https://github.com/fetristan/phone-number-range-finder-by-provider.git
  cd phone-number-range-finder-by-provider

Else, download the latest release from the [releases page](https://github.com/fetristan/phone-number-range-finder-by-provider/releases).

## Usage
1. Ensure your input file follows the format: number;provider

2. Run the program

The options are :
- `-i "./did.csv"` : Path for the phone numbers and providers
- `-o "./ranges_export.csv"` : Path for the ranges with details
- `-h` : display different usable options and quit

    2.1 From Source

        2.1.1 Place your input CSV file in the project directory as did.csv
        go run main.go [options]

    2.2 From binary

        2.2.1 Execute the binary file.
        phone-number-range-finder-by-provider [options]

3. Check the generated `ranges_export.csv` in the project directory for the output.

## Input File Format
The input CSV file should have the following format:
number;provider
number;provider
...

Example:
33123456789;orange
33987654321;sfr
...

## Output File Format
The output CSV file will have the following columns:
- Start: Starting number of the range.
- End: Ending number of the range.
- Provider: Provider name.
- Range Length: Number of phone numbers in the range.

## Contributing
Contributions to this project are welcome. Please ensure to update tests as appropriate.
