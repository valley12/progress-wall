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
│   ├── layout/          # 布局组件
│   └── features/        # 功能组件
├── views/               # 页面组件
│   ├── HomeView.vue     # 首页（独立页面，唯一）
│   ├── dashboard/       # 仪表板相关页面
│   ├── user/            # 用户相关页面
│   └── auth/            # 认证相关页面
├── stores/              # Pinia状态管理
├── composables/         # 可复用逻辑
└── lib/                 # 工具函数
```

#### 组件命名规范
- UI组件：`Button.vue`, `Card.vue`
- 功能组件：`TaskCard.vue`, `KanbanColumn.vue`
- 页面组件：`HomeView.vue`, `ProfileView.vue`

### 🔄 状态管理

#### Pinia Store规范
- 每个功能模块一个store文件
- 使用Composition API风格
- 状态更新通过actions，避免直接修改state
- 复杂状态考虑使用computed

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
