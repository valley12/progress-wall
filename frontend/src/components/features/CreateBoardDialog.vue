<template>
  <!-- 创建看板对话框 -->
  <div v-if="isOpen" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-background rounded-lg p-6 w-full max-w-md mx-4">
      <h2 class="text-xl font-semibold mb-4">创建新看板</h2>
      <form @submit.prevent="handleSubmit">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-foreground mb-2">看板名称</label>
            <input
              v-model="form.name"
              type="text"
              required
              class="w-full px-3 py-2 border border-input rounded-md focus:outline-none focus:ring-2 focus:ring-ring"
              placeholder="输入看板名称"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-foreground mb-2">描述（可选）</label>
            <textarea
              v-model="form.description"
              class="w-full px-3 py-2 border border-input rounded-md focus:outline-none focus:ring-2 focus:ring-ring"
              rows="3"
              placeholder="输入看板描述"
            ></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-foreground mb-2">颜色</label>
            <div class="space-y-3">
              <!-- 预设颜色 -->
              <div>
                <div class="flex flex-wrap gap-2">
                  <button
                    v-for="color in presetColors"
                    :key="color"
                    type="button"
                    @click="form.color = color"
                    class="w-8 h-8 rounded-full border-2 transition-all hover:scale-110"
                    :class="form.color === color ? 'border-foreground shadow-md' : 'border-transparent hover:border-muted-foreground'"
                    :style="{ backgroundColor: color }"
                    :title="color"
                  ></button>
                </div>
              </div>
              
              <!-- 自定义颜色 -->
              <div>
                <div class="space-y-2">
                  <div class="flex items-center space-x-3">
                    <input
                      v-model="form.color"
                      type="color"
                      class="w-12 h-8 rounded border border-input cursor-pointer"
                      title="选择自定义颜色"
                    />
                    <input
                      v-model="form.color"
                      type="text"
                      class="flex-1 px-3 py-2 border rounded-md focus:outline-none focus:ring-2 text-sm font-mono"
                      :class="isValidColor(form.color) 
                        ? 'border-input focus:ring-ring' 
                        : 'border-red-500 focus:ring-red-500'"
                      placeholder="#000000"
                      pattern="^#[0-9A-Fa-f]{6}$"
                      title="输入十六进制颜色代码"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="flex justify-end space-x-2 mt-6">
          <Button type="button" variant="outline" @click="handleCancel">
            取消
          </Button>
          <Button type="submit" :disabled="loading">
            {{ loading ? '创建中...' : '创建' }}
          </Button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, watch } from 'vue'
import Button from '@/components/ui/Button.vue'

interface Props {
  isOpen: boolean
  loading?: boolean
}

interface Emits {
  (e: 'submit', data: { name: string; description: string; color: string }): void
  (e: 'cancel'): void
}

const props = withDefaults(defineProps<Props>(), {
  loading: false
})

const emit = defineEmits<Emits>()

const form = reactive({
  name: '',
  description: '',
  color: '#3B82F6'
})

const presetColors = [
  // 蓝色系
  '#3B82F6',
  // 绿色系
  '#10B981',
  // 红色系
  '#EF4444',
  // 黄色系
  '#F59E0B',
  // 紫色系
  '#8B5CF6',
  // 青色系
  '#06B6D4',
  // 橙色系
  '#F97316',
  // 粉色系
  '#EC4899',
  // 灰色系
  '#6B7280',
  // 青绿色系
  '#84CC16',
]

// 验证颜色格式
const isValidColor = (color: string): boolean => {
  if (!color) return false
  if (color.startsWith('#')) {
    const hex = color.slice(1)
    return /^[0-9A-Fa-f]{6}$/.test(hex)
  }
  return false
}

// 验证和格式化颜色
const validateAndFormatColor = (color: string): string => {
  // 如果颜色以 # 开头，确保格式正确
  if (color.startsWith('#')) {
    // 移除 # 并确保是6位十六进制
    const hex = color.slice(1).toUpperCase()
    if (/^[0-9A-F]{6}$/.test(hex)) {
      return `#${hex}`
    }
  }
  // 如果不是有效的十六进制颜色，返回默认颜色
  return '#3B82F6'
}

const handleSubmit = () => {
  if (form.name.trim()) {
    emit('submit', {
      name: form.name.trim(),
      description: form.description.trim(),
      color: validateAndFormatColor(form.color)
    })
  }
}

const handleCancel = () => {
  emit('cancel')
}

// 当对话框关闭时重置表单
watch(() => props.isOpen, (isOpen) => {
  if (!isOpen) {
    form.name = ''
    form.description = ''
    form.color = '#3B82F6'
  }
})
</script>
