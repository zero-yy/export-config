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
  
  
## ConfLoader.cs
```c#
using UnityEngine;
using System.Collections.Generic;
using System;
using System.IO;
using Google.Protobuf;

namespace Config
{
    public static partial class ConfLoader
    {
        // 数据表名
        static Dictionary<string, object> tables = new Dictionary<string, object>();

        // todo.. 和版本兼容... 
        static string versionPath = string.Format("{0}_{1}_{2}", 1, 0, 0);

        // 数据存档目录，，热更数据用
        public static string GetDataPath()
        {
            return Application.persistentDataPath + "/cfgdata_" + versionPath + "/";
        }

        // 资源默认目录
        const string CfgPath = "Config/";
        const string NamespacePrefix = "Config.";

        public static List<string> TableNames = new List<string>();

        // 一次性函数
        static ConfLoader()
        {
            var path = GetDataPath(); 
            if (!Directory.Exists(path))
            {
                Directory.CreateDirectory(path);
            }

            AddTables();

            foreach (string name in TableNames)
            {
                Stream sm = null;
                if (File.Exists(path + name))
                    sm = new FileStream(path + name, FileMode.Open);
                else
                {
                    var ft = Resources.Load(CfgPath + name) as TextAsset;
                    sm = new MemoryStream(ft.bytes);
                }

                Type t = Type.GetType("Config." + name);
                cacheTable(t, sm);
                sm.Close();
            }
        }
        
        public static T GetTable<T>()
        {
            string name = typeof(T).Name;
            if (tables.ContainsKey(name))
                return (T) tables[name];
            else
                throw new Exception("no static config table:" + name);
        }

        public static void SaveTable(string name, byte[] data)
        {
            using (MemoryStream ms = new MemoryStream(data))
            {
                Type t = Type.GetType(NamespacePrefix  + name);
                cacheTable(t, ms);
                if (File.Exists(GetDataPath() + name))
                    File.Delete(GetDataPath() + name);
                using (FileStream fs = new FileStream(GetDataPath() + name, FileMode.Create))
                {
                    fs.Write(data, 0, data.Length);
                }
            }
        }

        static void cacheTable(Type t, Stream sm)
        {
            IMessage table = t.Assembly.CreateInstance(t.ToString()) as IMessage;
            table.MergeFrom(sm);

            tables[t.Name] = table;
        }

        public static void ReloadTable(string name)
        {
            Debug.LogFormat("Reload table {0}", name);

            Stream sm = null;
            if (File.Exists(GetDataPath() + name))
                sm = new FileStream(GetDataPath() + name, FileMode.Open);
            else
            {
                var ft = Resources.Load(CfgPath + name) as TextAsset;
                sm = new MemoryStream(ft.bytes);
            }

            Type t = Type.GetType(NamespacePrefix  + name);
            cacheTable(t, sm);
            sm.Close();
        }
    }
}
```