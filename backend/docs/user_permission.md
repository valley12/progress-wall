# 用户权限设计

## 权限树

所有的权限都基于根权限'*'，当检查权限时，扫描器会检查所需权限是否在的用户拥有的权限的子集中，举几个例子：

全局管理员：'*'
用户管理员：'users.*'
Project A (project id=1) 项目管理员：project.1.*
Project B (project id=2) 项目管理员：project.2.*
Project A (project id=1) 项目成员：project.1.view; project.1.write; ...
