# slot-machine-with-go

## Download and Install GO
https://go.dev/dl/

### Commands
#### To build & run in Windows
```
# create a build file
go build <filepath> 

# run the build file
./<built-file>.exe

# OR run directly which will compile and then run without creating a build file
go run <filepath>

# to create the project with export and import
go mod init <project-name>

# to run the project | (".") will search the mod file in the root
go run .

# or together
 go mod init <project-name> && go run .

```