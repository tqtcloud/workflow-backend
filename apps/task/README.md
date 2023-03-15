# jenkins API 模块

## Create Job    Post

### 请求：

```json
{
    "jobName":"apijob0925",
    "create_by":"xxx",
    "env":"DEV",
    "gitUrl": "https://github.com/tqtcloud/workflow-backend.git",
    "branch": "sit",
    "description": "0925 test api job",
    "appName": "apijob0925",
    //buildeshell 此处可不填,应用模板的打包命令  postman提交的时候请删除
    //"buildeshell": "echo 1111",
    "folder": "test2",
    "templateName": "jobtemplate/job/go-backend-template"
}

```

### 字段解释

|    参数名    |  类型  |               描述               |     默认值     | 是否必须 |
| :----------: | :----: | :------------------------------: | :------------: | :------: |
|   jobName    | string |         jenkins job 名称         |    默认为空    |   must   |
|  create_by   | string |          任务发起创建人          |    默认为空    |   must   |
|     env      | string | Jenkins环境：dev,qa,uat,lpt,prod |    默认为空    |   must   |
|    gitUrl    | string |             Git 地址             |    默认为空    |   must   |
|    branch    | string |           Git 默认分支           |    默认为空    |   must   |
| description  | string |           job 描述信息           |    默认为空    |   must   |
|   appName    | string |             应用名称             |    默认为空    |   must   |
| buildeshell  | string |            打包shell             | 默认为template | optional |
|    folder    | string |          jenkins 文件夹          |    默认为空    |   must   |
| templateName | string |       基于的jenkins模板job       |    默认为空    |   must   |

> **注意：**
>
> templateName字段为job类型模板字段，按照现有思路作为一个固定参数由前端直接传递
>
> 例如：前端为一个下拉框，可以选择后端模板为：go，Java，nodejs等

### 正常返回

```json
{
    "code": 0,
    "data": {
        "id": "cco15vbehud2k23mhs10",
        "create_at": 1664094973994602,
        "update_at": 0,
        "update_by": "",
        "data": {
            "create_by": "tqt",
            "jobName": "apijob0925",
            "Env": "DEV",
            "gitUrl": "https://github.com/tqtcloud/workflow-backend.git",
            "branch": "sit",
            "appName": "apijob0925",
            "description": "0925 test api job",
            "buildeshell": "",
            "buildenv": "",
            "folder": "test2",
            "oldjobName": "",
            "newjobName": "",
            "templateName": "jobtemplate/job/go-backend-template"
        }
    }
}
```

## Get Job   Get

### 请求：

```bash
# 访问地址：/workflow-backend/api/v1/task/{env}/{folder}/{jobname}
例如：
http://127.0.0.1:8050/workflow-backend/api/v1/task/dev/test2/apideljob111
```

### 字段解释

| 参数名  |  类型  |               描述               |  默认值  | 是否必须 |
| :-----: | :----: | :------------------------------: | :------: | :------: |
| jobname | string |         jenkins job 名称         | 默认为空 |   must   |
|   env   | string | Jenkins环境：dev,qa,uat,lpt,prod | 默认为空 |   must   |
| folder  | string |          jenkins 文件夹          | 默认为空 |   must   |

> **注意：**
>
> 如果你不写放回 404 
>
> 删除 jenkins 根的job  {folder} 置 `root`
>
> ```bash
> http://127.0.0.1:8050/workflow-backend/api/v1/task/dev/root/apideljob111
> ```

### 正常返回

```json
{
    "code": 0,
    "data": {
        "id": "9cf81a77-eea0-41dc-8519-afeecfbf065d",
        "create_at": 1664359420914578,
        "update_at": 0,
        "update_by": "",
        "data": {
            "create_by": "",
            "jobName": "test2/job/apijob0927",
            "Env": "DEV",
            "gitUrl": "https://github.com/tqtcloud/workflow-backend.git",
            "branch": "sit",
            "appName": "apijob0927",
            "description": "api调用测试job，请勿启动运行",
            "buildeshell": "echo 1111",
            "buildenv": "",
            "folder": "test2",
            "oldjobName": "",
            "newjobName": "",
            "templateName": ""
        }
    }
}
```



## Delete  Job Get

### 请求：

```bash
# 访问地址：/workflow-backend/api/v1/task/{env}/{folder}/{jobname}
例如：
http://127.0.0.1:8050/workflow-backend/api/v1/task/dev/test2/apideljob111
```

### 字段解释

| 参数名  |  类型  |               描述               |  默认值  | 是否必须 |
| :-----: | :----: | :------------------------------: | :------: | :------: |
| jobname | string |         jenkins job 名称         | 默认为空 |   must   |
|   env   | string | Jenkins环境：dev,qa,uat,lpt,prod | 默认为空 |   must   |
| folder  | string |          jenkins 文件夹          | 默认为空 |   must   |

> **注意：**
>
> 如果你不写放回 404 
>
> 删除 jenkins 根的job  {folder} 置 `root`
>
> ```bash
> http://127.0.0.1:8050/workflow-backend/api/v1/task/dev/root/apideljob111
> ```

### 正常返回

```json
{
    "code": 0,
    "data": {
        "id": "9cf81a77-eea0-41dc-8519-afeecfbf065d",
        "create_at": 1664359420914578,
        "update_at": 0,
        "update_by": "",
        "data": {
            "create_by": "",
            "jobName": "test2/job/apijob0927",
            "Env": "DEV",
            "gitUrl": "https://github.com/tqtcloud/workflow-backend.git",
            "branch": "sit",
            "appName": "apijob0927",
            "description": "api调用测试job，请勿启动运行",
            "buildeshell": "echo 1111",
            "buildenv": "",
            "folder": "test2",
            "oldjobName": "",
            "newjobName": "",
            "templateName": ""
        }
    }
}
```

## Update  Job  Patch

### 请求：

```bash
# 访问地址：/workflow-backend/api/v1/task/{env}/{folder}/{jobname}
例如：
http://127.0.0.1:8050/workflow-backend/api/v1/task/dev/test2/apideljob111
```

```json
{
    "jobName":"apijob0927",
    "create_by":"tqt44444",
    "env":"test",
    "branch": "sit",
    "gitUrl": "https://github.com/tqtcloud/workflow-backend.git",
    "description": "api调用测试job，请勿启动运行44444put",
    "appName": "apijob0927",
    "buildeshell": "echo api调用测试job，请勿启动运行44444put",
    "folder": "test2",
    "templateName": "apijob0927"
}
```

### 字段解释

|    参数名    |  类型  |               描述               |     默认值     | 是否必须 |
| :----------: | :----: | :------------------------------: | :------------: | :------: |
|   jobName    | string |         jenkins job 名称         |    默认为空    |   must   |
|  create_by   | string |          任务发起创建人          |    默认为空    |   must   |
|     env      | string | Jenkins环境：dev,qa,uat,lpt,prod |    默认为空    |   must   |
|    gitUrl    | string |          修改后Git 地址          |    默认为空    |   must   |
|    branch    | string |        修改后Git 默认分支        |    默认为空    |   must   |
| description  | string |       修改后 job 描述信息        |    默认为空    |   must   |
|   appName    | string |         修改后 应用名称          |    默认为空    |   must   |
| buildeshell  | string |         修改后 打包shell         | 默认为template | optional |
|    folder    | string |          jenkins 文件夹          |    默认为空    |   must   |
| templateName | string |    基于的jenkins模板job，随意    |    默认为空    |   must   |

### 正常返回

```json
{
    "code": 0,
    "data": {
        "id": "0c2c022f-c03d-4e6d-a0e7-eac2fb1c96d6",
        "create_at": 1664456694821648,
        "update_at": 1664456694821648,
        "update_by": "",
        "data": {
            "create_by": "tqt44444",
            "jobName": "apijob0927",
            "Env": "TEST",
            "gitUrl": "https://github.com/tqtcloud/workflow-backend.git",
            "branch": "sit",
            "appName": "apijob0927",
            "description": "api调用测试job，请勿启动运行44444put",
            "buildeshell": "echo api调用测试job，请勿启动运行44444put",
            "buildenv": "",
            "folder": "test2",
            "oldjobName": "",
            "newjobName": "",
            "templateName": "apijob0927"
        }
    }
}
```

## Update  Job  Put

### 请求：

```bash
# 访问地址：/workflow-backend/api/v1/task/{env}/{folder}/{jobname}
例如：
http://127.0.0.1:8050/workflow-backend/api/v1/task/dev/test2/apideljob111
```

```json
{
    "jobName":"apijob0927",
    "create_by":"tqt44444",
    "env":"test",
    "branch": "sit",
    "gitUrl": "https://github.com/tqtcloud/workflow-backend.git",
    "description": "api调用测试job，请勿启动运行44444put",
    "appName": "apijob0927",
    "buildeshell": "echo api调用测试job，请勿启动运行44444put",
    "folder": "test2",
    "templateName": "apijob0927"
}
```

### 字段解释

|    参数名    |  类型  |               描述               |     默认值     | 是否必须 |
| :----------: | :----: | :------------------------------: | :------------: | :------: |
|   jobName    | string |         jenkins job 名称         |    默认为空    |   must   |
|  create_by   | string |          任务发起创建人          |    默认为空    |   must   |
|     env      | string | Jenkins环境：dev,qa,uat,lpt,prod |    默认为空    |   must   |
|    gitUrl    | string |          修改后Git 地址          |    默认为空    |   must   |
|    branch    | string |        修改后Git 默认分支        |    默认为空    |   must   |
| description  | string |       修改后 job 描述信息        |    默认为空    |   must   |
|   appName    | string |         修改后 应用名称          |    默认为空    |   must   |
| buildeshell  | string |         修改后 打包shell         | 默认为template | optional |
|    folder    | string |          jenkins 文件夹          |    默认为空    |   must   |
| templateName | string |    基于的jenkins模板job，随意    |    默认为空    |   must   |

### 正常返回

```json
{
    "code": 0,
    "data": {
        "id": "0c2c022f-c03d-4e6d-a0e7-eac2fb1c96d6",
        "create_at": 1664456694821648,
        "update_at": 1664456694821648,
        "update_by": "",
        "data": {
            "create_by": "tqt44444",
            "jobName": "apijob0927",
            "Env": "TEST",
            "gitUrl": "https://github.com/tqtcloud/workflow-backend.git",
            "branch": "sit",
            "appName": "apijob0927",
            "description": "api调用测试job，请勿启动运行44444put",
            "buildeshell": "echo api调用测试job，请勿启动运行44444put",
            "buildenv": "",
            "folder": "test2",
            "oldjobName": "",
            "newjobName": "",
            "templateName": "apijob0927"
        }
    }
}
```

## 
