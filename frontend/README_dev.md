# 开发人员指南

## 开发注意事项

### 📋 任务管理
- **及时更新TODO**: 在开发过程中，请及时在Github的TODO列表(https://github.com/users/valley12/projects/1 )中进行认领和更新，以防止重复开发
- **任务分解**: 将复杂功能分解为小的、可管理的任务

### 🏗️ 代码规范

#### Vue组件开发
- 使用Composition API和`<script setup>`语法
- 组件名使用PascalCase命名
- 单个文件不超过200行，超过时考虑拆分
- 保持组件的单一职责原则

#### TypeScript规范
- 为所有props、emits、composables定义类型
- 使用接口定义复杂数据结构
- 尽量避免使用`any`类型，优先使用具体类型

#### 样式规范
- 优先使用Tailwind CSS类名，避免自定义CSS
- 注意在各种设备上的兼容性（后面可能考虑套Electron来实现跨平台
- 注意CSS作用范围，避免滥用全局样式

#### 代码注释
- 保持中文注释

### 🧩 组件架构

#### 组件层次结构
后续可以继续增加，但注意更新本文档

```
src/
├── components/
│   ├── ui/              # shadcn/ui基础组件
│   │   ├── Avatar.vue
│   │   ├── Button.vue
│   │   ├── Card.vue
│   │   ├── CardContent.vue
│   │   ├── CardHeader.vue
│   │   ├── DropdownMenu.vue
│   │   └── DropdownMenuItem.vue
│   ├── layout/          # 布局组件
│   │   └── Navbar.vue
│   └── features/        # 功能组件
│       ├── BoardCard.vue        # 看板卡片组件
│       ├── CreateBoardDialog.vue # 创建看板对话框
│       ├── KanbanColumn.vue     # 看板列组件
│       └── TaskCard.vue         # 任务卡片组件
├── views/               # 页面组件
│   ├── HomeView.vue     # 首页（独立页面，唯一）
│   ├── dashboard/       # 仪表板相关页面
│   │   ├── BoardListView.vue    # 看板列表页面
│   │   └── KanbanView.vue       # 看板详情页面
│   └── user/            # 用户相关页面
│       ├── ProfileView.vue      # 用户资料页面
│       └── SettingsView.vue     # 用户设置页面
├── stores/              # Pinia状态管理
│   ├── board.ts         # 看板状态管理
│   ├── kanban.ts        # 看板任务状态管理
│   ├── user.ts          # 用户状态管理
│   └── index.ts         # Store统一导出
├── services/            # API服务层
│   ├── base-api.ts      # 基础API服务类
│   ├── board-api.ts     # 看板API服务
│   └── index.ts         # 服务统一导出
├── config/              # 配置文件
│   └── api.ts           # API配置
├── router/              # 路由配置
│   └── index.ts
├── lib/                 # 工具函数
│   └── utils.ts
└── assets/              # 静态资源
    └── vue.svg
```

#### 组件命名规范
- **UI组件**：`Button.vue`, `Card.vue`, `Avatar.vue` - 基础UI组件，可复用
- **功能组件**：`BoardCard.vue`, `TaskCard.vue`, `KanbanColumn.vue`, `CreateBoardDialog.vue` - 业务功能组件
- **布局组件**：`Navbar.vue` - 页面布局相关组件
- **页面组件**：`HomeView.vue`, `BoardListView.vue`, `KanbanView.vue`, `ProfileView.vue` - 路由页面组件

#### 新增组件说明
- **BoardCard.vue**: 看板卡片组件，支持卡片/列表两种显示模式
- **CreateBoardDialog.vue**: 创建看板对话框，包含颜色选择器功能
- **KanbanColumn.vue**: 看板列组件，用于显示任务列
- **TaskCard.vue**: 任务卡片组件，显示单个任务信息

### 🔄 状态管理

#### Pinia Store规范
- 每个功能模块一个store文件
- 使用Composition API风格
- 状态更新通过actions，避免直接修改state
- 复杂状态考虑使用computed

#### 当前Store结构
- **board.ts**: 看板状态管理，包含看板列表、创建、更新、删除等功能
- **kanban.ts**: 看板任务状态管理，包含任务CRUD操作
- **user.ts**: 用户状态管理，用户信息和设置
- **index.ts**: Store统一导出文件

### 🌐 API服务层

#### 服务架构
- **base-api.ts**: 基础API服务类，提供通用的HTTP请求方法
- **board-api.ts**: 看板相关API服务，继承自BaseApiService
- **index.ts**: 服务统一导出文件

#### API配置
- **config/api.ts**: API配置文件，包含端点配置和HTTP方法
- 支持同时定义请求路径和HTTP方法
- 提供工具函数构建API URL

#### 使用示例
```typescript
// 在组件中使用API服务
import { boardApiService } from '@/services/board-api'

const response = await boardApiService.getKanbanList()
```

### 🛣️ 路由配置

#### 当前路由结构
- **/** - HomeView.vue (首页)
- **/boards** - BoardListView.vue (看板列表页面)
- **/kanban/:boardId?** - KanbanView.vue (看板详情页面，支持可选参数)
- **/profile** - ProfileView.vue (用户资料页面)
- **/settings** - SettingsView.vue (用户设置页面)

#### 路由特点
- 使用动态路由参数支持看板ID
- 支持可选参数，提供向后兼容性
- 采用懒加载方式导入组件

#### 示例Store结构
```typescript
export const useFeatureStore = defineStore('feature', () => {
  // 状态
  const state = ref(initialValue)
  
  // 计算属性
  const computedValue = computed(() => {
    // 计算逻辑
  })
  
  // 方法
  const action = () => {
    // 业务逻辑
  }
  
  return {
    state,
    computedValue,
    action
  }
})
```

### 📦 依赖管理

#### 添加新依赖
```bash
# 生产依赖
pnpm add package-name

# 开发依赖
pnpm add -D package-name

# 更新依赖
pnpm update
```

### 🧪 测试策略

#### 测试工具
1. Postman - API测试和文档管理(群里有加入链接)
...

### 🔒 安全考虑

#### 数据安全
- 敏感信息不要硬编码
- 使用环境变量管理配置
- 输入验证和XSS防护

#### 代码安全
- 定期更新依赖包
- 使用ESLint进行代码检查
- 避免使用eval等危险函数

### 🔄 版本控制

#### Git工作流
- 使用feature分支开发新功能
- 提交信息使用约定式提交格式
- 合并前进行代码审查
- 保持主分支的稳定性
- 在发起合并前请确保代码中不含任何编译器报告的警告和错误

#### 提交信息规范
```
feat: 添加新功能
fix: 修复bug
docs: 更新文档
style: 代码格式调整
refactor: 代码重构
test: 添加测试
chore: 构建过程或辅助工具的变动
```

### 杂项

#### Github Project的使用

##### Label

说明几个已经用到/可能会用到的Label，这里只对部分进行补充

front: 前端需求
back: 后端需求
help wanted: 考虑到有些工作可能需要不止一个人，或者说发现工作量巨大/无法完成，可以打上这个标进行求助
bug: 用于标记代码中存在的错误或缺陷，表明这是一个需要被修复的问题。某个已经Done的task中包含的代码有bug将重新开一个修复bug的task（已经完成的任务不会回滚）
documentation: 指示此任务与改进或补充项目文档相关。

#### Estimate
Estimate (预估) 字段是一个自定义字段，其核心作用是帮助团队量化和评估完成一个任务（Issue）所需的工作量、复杂性或时间。

通过为每个任务设置 Estimate，团队可以计算出在一次迭代（Sprint）或一个特定时期内总共计划了多少工作量。这有助于判断计划是否现实，并能更准确地预测未来的工作吞吐量（即团队速率 Velocity）。

在 Project 的视图（如表格或看板）中，你可以对 Estimate 字段进行聚合计算。例如，你可以让看板的每一列（如 "To Do", "In Progress", "Done"）自动求和 (Sum) 该列所有卡片的 Estimate 值。

如：https://github.com/users/valley12/projects/1/views/1 中每一列都有 Estimate 总和

这是敏捷开发（尤其是 Scrum）中最流行的用法。故事点是一个相对值，它综合评估了任务的三个方面：

1. 工作量 (Volume of work): 需要做多少事？

2. 复杂性 (Complexity): 技术难度有多高？

3. 不确定性 (Uncertainty/Risk): 需求清晰吗？是否存在未知风险？

团队通常使用斐波那契数列（如 1, 2, 3, 5, 8, 13）来作为故事点的值，以体现随着任务变大，估算的不确定性也在增加。

示例:

Estimate: 1 = 一个非常简单、清晰的任务。

Estimate: 5 = 一个中等复杂度的任务。

Estimate: 13 = 一个非常复杂或充满不确定性的大任务，可能需要拆分成更小的任务。

也有用衣服尺码的(XS, S, M, L, XL)，但是暂定斐波那契。