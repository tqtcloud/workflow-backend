# workflow-backend

jenkins工作流后端

### 模板划分：

- [x] argocd-deploy-template       argocd触发模板  

  

- [x] go-backend-template-build-argocd   golang构建部署模板

- [x] go-backend-template-build                 golang构建模板

- [x] go-backend-template-build-ssh          golang部署模板   基于ssh 

  

- [x] Java-backend-template-build-argocd      Java构建部署模板    基于argocd

- [ ] Java-backend-template-build-ssh       Java构建部署模板    基于ssh

- [x] Java-backend-template-build              Java构建模板

  

- [x] nodejs-template-build-deploy            nodejs构建部署模板     基于ssh部署

- [x] nodejs-template-build                         nodejs构建模板

**注意：需要调整所有的job参数**

> build-deploy 模板需要argocd的部分参数
>
> APP="$JOB_BASE_NAME"  获取当前job的名称
> TYPE="lebei"   k8s中名称空间后面部分
> ENV="sit"  k8s中环境部分



理想状态为：

- 开发环境提交一次请求，创建构建部署job
- 其他环境提交2次请求， 创建2个job，构建，部署
- 任务执行有通知，成功有通知

### 额外功能：
- [x] 可根据接口远程主机执行命令

### 存在问题：

1. 可能有部分情况没有考虑在内
   1. 生产环境 ssh 多台主机