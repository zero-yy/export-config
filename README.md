# export-config
Parse excel config data to protobuf, then generate code(golang&amp;c#) and data(bytes&json).

## Install
`go get github.com/zero-yy/export-config`  
`sudo ln -s ~/go/bin/export-config /usr/local/bin/export-config`

## Usage
* prepare config.toml
```toml
######################
InputExcelPath      = "./test"
OutputProtoPath     = "./test"

OutputDataCSharpPath = "./test/csharp/conf_data"
OutputDataGoPath     = "./test/go/conf_data"
OutputCodeCSharpPath = "./test/csharp/gen/conf"
OutputCodeGoPath     = "./test/go/gen/conf"

ProtoPackage        = "Conf"
CSharpNamespace     = "conf"
GoPackage           = "conf"
GoPackageFullPath   = "github.com/zero-yy/export-config/test/go/gen/conf"
``` 
* gen code   
`export-config --config=config.toml`

* build gendata
`go build -o gendata {{.OutputCodeGoPath}}/cmd/gendata/main.go`

* gen data 
`gendata --config=config.toml`

## Convention over configuration
### excel format
> reference: ./test/test.xlsm
>  
>line 1: comment  
>line 2: column name  
>line 3: column type 
>line 4-n: data  
>col 1: comment

#### column name
* using decorator @id to indicate the column as id column
* if not found @id, then column with name "id" will be id column
* ary/0 ary/1 ary/x will merge to array. Decorator "/.*" will be regarded as array sign. 
#### column type
* int32
* int
* int64
* long
* 
* uint32
* uint
* uint64
* ulong
* 
* string
* 
* float
* float32
* double

#### data
* If id column have no data, then skip the row.  
  
  