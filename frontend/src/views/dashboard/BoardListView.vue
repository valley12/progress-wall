<template>
  <div class="min-h-screen bg-background">
    <div class="container mx-auto px-4 py-8">
      <!-- 页面头部 -->
      <div class="flex justify-between items-center mb-8">
        <div>
          <h1 class="text-3xl font-bold text-foreground">我的看板</h1>
          <p class="text-muted-foreground mt-2">管理和组织您的项目看板</p>
        </div>
        <div class="flex items-center space-x-4">
          <!-- 视图模式切换 -->
          <div class="flex items-center bg-muted rounded-lg p-1">
            <button
              @click="viewMode = 'card'"
              :class="[
                'flex items-center space-x-2 px-3 py-2 rounded-md text-sm font-medium transition-colors',
                viewMode === 'card' 
                  ? 'bg-background text-foreground shadow-sm' 
                  : 'text-muted-foreground hover:text-foreground'
              ]"
            >
              <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
              </svg>
              <span>卡片</span>
            </button>
            <button
              @click="viewMode = 'list'"
              :class="[
                'flex items-center space-x-2 px-3 py-2 rounded-md text-sm font-medium transition-colors',
                viewMode === 'list' 
                  ? 'bg-background text-foreground shadow-sm' 
                  : 'text-muted-foreground hover:text-foreground'
              ]"
            >
              <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
              </svg>
              <span>列表</span>
            </button>
          </div>
          <Button @click="showCreateDialog = true" class="flex items-center space-x-2">
            <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            <span>创建看板</span>
          </Button>
        </div>
      </div>

      <!-- 加载状态 -->
      <div v-if="boardStore.loading" class="text-center py-12">
        <div class="mx-auto w-24 h-24 bg-muted rounded-full flex items-center justify-center mb-4">
          <svg class="w-8 h-8 text-muted-foreground animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-foreground mb-2">加载中...</h3>
        <p class="text-muted-foreground">正在获取看板列表</p>
      </div>

      <!-- 错误状态 -->
      <div v-else-if="boardStore.error" class="text-center py-12">
        <div class="mx-auto w-24 h-24 bg-red-100 rounded-full flex items-center justify-center mb-4">
          <svg class="w-12 h-12 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-foreground mb-2">加载失败</h3>
        <p class="text-muted-foreground mb-4">{{ boardStore.error }}</p>
        <Button @click="boardStore.fetchBoards()">
          重试
        </Button>
      </div>

      <!-- 看板列表 -->
      <div v-else-if="boardStore.boards.length > 0">
        <!-- 卡片视图 -->
        <div v-if="viewMode === 'card'" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <BoardCard
            v-for="board in boardStore.boards"
            :key="board.id"
            :board="board"
            :view-mode="'card'"
            @click="openBoard"
            @edit="editBoard"
            @delete="deleteBoard"
          />
        </div>
        
        <!-- 列表视图 -->
        <div v-else class="bg-background border border-border rounded-lg overflow-hidden">
          <div class="bg-muted/50 px-4 py-3 border-b border-border">
            <div class="flex items-center justify-between text-sm font-medium text-muted-foreground">
              <div class="flex items-center space-x-4">
                <span class="w-3"></span> <!-- 占位符，与列表项对齐 -->
                <span class="flex-1">看板名称</span>
              </div>
              <div class="flex items-center space-x-4">
                <span>创建时间</span>
                <span>更新时间</span>
                <span class="w-8"></span> <!-- 占位符，与操作按钮对齐 -->
              </div>
            </div>
          </div>
          <BoardCard
            v-for="board in boardStore.boards"
            :key="board.id"
            :board="board"
            :view-mode="'list'"
            @click="openBoard"
            @edit="editBoard"
            @delete="deleteBoard"
          />
        </div>
      </div>

      <!-- 空状态 -->
      <div v-else class="text-center py-12">
        <div class="mx-auto w-24 h-24 bg-muted rounded-full flex items-center justify-center mb-4">
          <svg class="w-12 h-12 text-muted-foreground" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17V7m0 10a2 2 0 01-2 2H5a2 2 0 01-2-2V7a2 2 0 012-2h2a2 2 0 012 2m0 10a2 2 0 002 2h2a2 2 0 002-2M9 7a2 2 0 012-2h2a2 2 0 012 2m0 10V7m0 10a2 2 0 002 2h2a2 2 0 002-2V7a2 2 0 00-2-2" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-foreground mb-2">还没有看板</h3>
        <p class="text-muted-foreground mb-6">创建您的第一个看板来开始管理任务</p>
        <Button @click="showCreateDialog = true" :disabled="boardStore.loading">
          {{ boardStore.loading ? '创建中...' : '创建看板' }}
        </Button>
      </div>

      <!-- 创建看板对话框 -->
      <CreateBoardDialog
        :is-open="showCreateDialog"
        :loading="boardStore.loading"
        @submit="createBoard"
        @cancel="cancelCreate"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useBoardStore } from '@/stores/board'
import type { Board } from '@/stores/board'
import Button from '@/components/ui/Button.vue'
import BoardCard from '@/components/features/BoardCard.vue'
import CreateBoardDialog from '@/components/features/CreateBoardDialog.vue'

const router = useRouter()
const boardStore = useBoardStore()

const showCreateDialog = ref(false)
const viewMode = ref<'card' | 'list'>('card')

// 组件挂载时获取看板列表
onMounted(async () => {
  await boardStore.fetchBoards()
})

const openBoard = (board: Board) => {
  // 跳转到具体的看板页面
  router.push(`/kanban/${board.id}`)
}

const editBoard = (board: Board) => {
  // 这里可以实现编辑看板的功能
  console.log('编辑看板:', board)
}

const deleteBoard = async (board: Board) => {
  if (confirm(`确定要删除看板 "${board.name}" 吗？`)) {
    const success = await boardStore.deleteBoard(board.id)
    if (!success) {
      alert('删除看板失败，请重试')
    }
  }
}

const createBoard = async (data: { name: string; description: string; color: string }) => {
  const result = await boardStore.addBoard(data)
  
  if (result) {
    showCreateDialog.value = false
  } else {
    alert('创建看板失败，请重试')
  }
}

const cancelCreate = () => {
  showCreateDialog.value = false
}
</script>
