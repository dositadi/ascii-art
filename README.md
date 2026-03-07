# ASCII-Art Generator

#### A robust command-line utility built in Go that transforms standard string input into large-scale graphical ASCII representations. This program handles letters (upper/lower), numbers, special characters, and complex newline sequences by mapping them to 8-line tall graphical blocks.


# Table of Contents

[1. Project Structure](#project-structure)

[2. Technical Architecture](#technical-architecture)

[3. Core Logic and Function](#core-logic--functions)

[4. Logic Flow](#logic-flow)

[5. Usage](#usage)

[6. Error Handling](#error-handling)

[7. License](#license)

# Project Structure
#### Based on your source code, the project is organized using a clean, modular architecture:

```plaintext
.
├── asset/                  # ASCII font templates (e.g., model.txt)
├── cmd/
│   └── main.go             # Application entry point
├── internal/
│   ├── config/             # Artist interface and Sketcher coordinator
│   └── arthandlers/        # Implementation of drawing and mapping logic
├── pkg/
│   ├── models/             # Custom error structs
│   └── utils/              # Character-to-line maps and error constants
└── README.md
```

# Technical Architecture
#### The project utilizes an Interface-Driven Design to separate the "drawing" logic from the "coordination" logic.

### 1. The Artist Interface
  #### Defined in artist.go, this interface forces any sketcher to implement:

* #### Start Line Detection: GetLowerCaseStartLine, GetUpperCaseStartLine, GetNumberStartLine, and GetSpecialCharStartLine.

* #### File Processing: ReadFileByLine to extract specific graphical blocks.

* #### Flow Control: CheckIfNewlineAndSplit and PrintArtOut.

# Core Logic & Functions
### 1. Intelligent Character Mapping (utils/)
#### The tool uses specialized maps to jump to the exact line in the template file where a character's representation begins:

* #### UpperCaseAlphabetsArt: Maps 'A'-'Z' (e.g., 'A' starts at line 299).

* #### LowerCaseAlphabetsArt: Maps 'a'-'z' (e.g., 'a' starts at line 587).

* #### NumbersArt: Maps digits '0'-'9'.

* #### SpecialCharactersArt: Maps 32+ symbols including space, brackets, and math operators.

### 2. The Sketching Engine (arthandlers/)
* #### DesignAllCharacters: Iterates through the input runes, determines their category, fetches the 8-line ASCII block for each, and stores them in a 3D slice ([][][]string).

* #### ReadFileByLine: Precisely scans the template file (defaulting to /asset/model.txt) to grab the exact 8 lines required for a character.

* #### CheckIfNewlineAndSplit: Handles literal \n characters in the input string, allowing for multi-line ASCII art generation.

### 3. Horizontal Rendering Logic
#### Because terminals print line-by-line, the program cannot simply print one character at a time.

* #### PrintArtOut: This function iterates through the 8 vertical rows of the font. For each row, it horizontally concatenates the corresponding row of every character in the word before moving to the next line.

# Logic Flow
#### 1. Entry: main.go initializes the App and calls Run().

#### 2. Setup: NewSketcher binds the Artist implementation to the user input.

#### 3. Processing:

* #### Input is split by newlines.

* #### Each character is mapped to a line number.

* #### 8-line blocks are fetched from the template.

* #### Output: The 3D store is rendered row-by-row to the terminal.

# Usage
#### Installation
#### Ensure you have Go 1.18+ installed:

```bash
git clone https://github.com/dositadi/ascii-art.git
cd ascii-art
```
### Running the Tool

```bash
# Standard input
go run main.go "Hello World"

# Input with newlines
go run main.go "First Line\nSecond Line"
```

### Example Output

```plaintext
 _    _          _   _               
| |  | |        | | | |              
| |__| |   ___  | | | |   ___        
|  __  |  / _ \ | | | |  / _ \       
| |  | | |  __/ | | | | | (_) |      
|_|  |_|  \___| |_| |_|  \___/
```

# Error Handling
#### The tool features a custom error system (pkg/models) that provides clear terminal feedback:

#### 1. SERVER_ERR: Issues reading the template file.

#### 2. ERR: Triggered when no input is provided.

#### 3. MULTI_INPUT_ERR: Triggered if more than one string argument is passed.

# License

### This project is licensed under the [MIT](LICENSE) License.