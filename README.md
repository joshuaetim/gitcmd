# gitcmd

gitcmd is a CLI application to create, update and retrieve Github issues.

## Installation & Setup

Ensure you have the Go compiler installed on your system. Visit this link to install Go: [https://go.dev/doc/install](https://go.dev/doc/install)

After installing, head to the folder directory on your terminal and run the following command:

```bash
go build .
```

On initial run, you will be asked to provide your access token. Please refer to this link to get your access token: [https://github.com/settings/tokens](https://github.com/settings/tokens)

**Note**: When copying/pasting the token, add the prefix "token" to the copied token. For example:
```bash
token ghp_uf78HGu78u8oysweaaDFrWh2usG1B 
```

Additionally, you will be asked to set your default repository. This can always be changed at every run.

**Object file name**: The folder name is assumed as "**gitcmd**" which may also be the name of compiled file.

## Usage

### Retrieve Issue
key command: get

parameters:
- **n**: issue number (required)

**Example**:
```bash
./gitcmd get -n=25
```
(Remove the first dot for Windows terminal)

### Create Issue
key command: create

parameters:
- **title**: title of issue (required)
- **body**: body text of issue

**Example (without editor)**:

```bash
./gitcmd create -title="Test Issue" -body="Body of test issue"
```

**Example (with editor)**:
```bash
EDITOR=code ./gitcmd create
```
_code_ = Visual studio code. Replace with your preferred editor.

### Update Issue
key command: update

parameters: 
- **n**: issue number (required)
- **title**: title of issue 
- **body**: body of issue
- **state**: State of issue ("open" or "closed")


**Example (without editor)**:

```bash
./gitcmd update -n=1 -title="New Title" -body="New Body of test issue" -state="closed"
```

**Example (with editor)**:
```bash
EDITOR=code ./gitcmd update -n=1
```
_code_ = Visual studio code. Replace with your preferred editor.

## Notes
Suggestions for improvement are welcome. Ideas for this project came from the Go textbook ["The Go Programming Language"](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440)