<template>
  <!-- 卡片样式 -->
  <Card 
    v-if="viewMode === 'card'"
    class="cursor-pointer hover:shadow-lg transition-shadow duration-200" 
    @click="handleClick"
  >
    <CardHeader>
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <div 
            class="w-4 h-4 rounded-full" 
            :style="{ backgroundColor: board.color }"
          ></div>
          <h3 class="text-lg font-semibold text-foreground">{{ board.name }}</h3>
        </div>
        <div @click.stop>
          <DropdownMenu>
            <template #trigger>
              <Button 
                variant="ghost" 
                size="sm" 
                class="h-8 w-8 p-0"
              >
                <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                </svg>
              </Button>
            </template>
            <DropdownMenuItem @click="handleEdit">
              编辑看板
            </DropdownMenuItem>
            <DropdownMenuItem @click="handleDelete" class="text-red-600">
              删除看板
            </DropdownMenuItem>
          </DropdownMenu>
        </div>
      </div>
    </CardHeader>
    <CardContent>
      <p class="text-sm text-muted-foreground mb-4">{{ board.description || '暂无描述' }}</p>
      <div class="flex items-center justify-between text-xs text-muted-foreground">
        <span>创建于 {{ formatDate(board.createdAt) }}</span>
        <span>更新于 {{ formatDate(board.updatedAt) }}</span>
      </div>
    </CardContent>
  </Card>

  <!-- 列表样式 -->
  <div 
    v-else
    class="cursor-pointer hover:bg-muted/50 transition-colors duration-200 border-b border-border py-4 px-4 flex items-center justify-between"
    @click="handleClick"
  >
    <div class="flex items-center space-x-4 flex-1">
      <div 
        class="w-3 h-3 rounded-full flex-shrink-0" 
        :style="{ backgroundColor: board.color }"
      ></div>
      <div class="flex-1 min-w-0">
        <h3 class="text-base font-semibold text-foreground truncate">{{ board.name }}</h3>
        <p class="text-sm text-muted-foreground truncate mt-1">{{ board.description || '暂无描述' }}</p>
      </div>
    </div>
    <div class="flex items-center space-x-4 text-xs text-muted-foreground flex-shrink-0">
      <span>创建于 {{ formatDate(board.createdAt) }}</span>
      <span>更新于 {{ formatDate(board.updatedAt) }}</span>
      <div @click.stop>
        <DropdownMenu>
          <template #trigger>
            <Button 
              variant="ghost" 
              size="sm" 
              class="h-8 w-8 p-0"
            >
              <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
              </svg>
            </Button>
          </template>
          <DropdownMenuItem @click="handleEdit">
            编辑看板
          </DropdownMenuItem>
          <DropdownMenuItem @click="handleDelete" class="text-red-600">
            删除看板
          </DropdownMenuItem>
        </DropdownMenu>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'
import type { Board } from '@/stores/board'
import Card from '@/components/ui/Card.vue'
import CardHeader from '@/components/ui/CardHeader.vue'
import CardContent from '@/components/ui/CardContent.vue'
import Button from '@/components/ui/Button.vue'
import DropdownMenu from '@/components/ui/DropdownMenu.vue'
import DropdownMenuItem from '@/components/ui/DropdownMenuItem.vue'

/**
 * 看板卡片组件
 * 支持卡片和列表两种显示模式
 * 提供看板的基本信息展示和操作功能
 */

interface Props {
  board: Board
  viewMode?: 'card' | 'list' // 显示模式：卡片或列表
}

const props = withDefaults(defineProps<Props>(), {
  viewMode: 'card'
})

const emit = defineEmits<{
  click: [board: Board]    // 点击看板事件
  edit: [board: Board]     // 编辑看板事件
  delete: [board: Board]   // 删除看板事件
}>()

/**
 * 处理看板点击事件
 * 触发父组件的点击回调
 */
const handleClick = () => {
  emit('click', props.board)
}

/**
 * 处理编辑看板事件
 * 触发父组件的编辑回调
 */
const handleEdit = () => {
  emit('edit', props.board)
}

/**
 * 处理删除看板事件
 * 触发父组件的删除回调
 */
const handleDelete = () => {
  emit('delete', props.board)
}

/**
 * 格式化日期显示
 * 使用中文本地化格式，显示年月日
 * @param date 要格式化的日期对象
 * @returns 格式化后的日期字符串
 */
const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('zh-CN', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date)
}
</script>
