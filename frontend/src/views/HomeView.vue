<template>
  <div class="min-h-screen bg-background">
    <div class="container mx-auto px-4 py-8">
      <div class="text-center mb-8">
        <h1 class="text-4xl font-bold text-foreground mb-4">
          进度看板
        </h1>
        <p class="text-muted-foreground text-lg">
          一个现代化的任务管理看板，帮助您高效管理项目进度
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
        <Card class="p-6">
          <CardHeader>
            <h3 class="text-xl font-semibold">任务管理</h3>
          </CardHeader>
          <CardContent>
            <p class="text-muted-foreground">
              创建、编辑和跟踪任务，支持拖拽排序和状态更新
            </p>
          </CardContent>
        </Card>

        <Card class="p-6">
          <CardHeader>
            <h3 class="text-xl font-semibold">进度跟踪</h3>
          </CardHeader>
          <CardContent>
            <p class="text-muted-foreground">
              实时查看项目进度，了解团队工作状态
            </p>
          </CardContent>
        </Card>

        <Card class="p-6">
          <CardHeader>
            <h3 class="text-xl font-semibold">协作工具</h3>
          </CardHeader>
          <CardContent>
            <p class="text-muted-foreground">
              团队成员可以协作管理任务，提高工作效率
            </p>
          </CardContent>
        </Card>
      </div>

      <div class="text-center">
        <Button @click="goToKanban" size="lg" class="px-8">
          开始使用看板
        </Button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import Card from '@/components/ui/Card.vue'
import CardHeader from '@/components/ui/CardHeader.vue'
import CardContent from '@/components/ui/CardContent.vue'
import Button from '@/components/ui/Button.vue'

const router = useRouter()

const goToKanban = () => {
  router.push('/kanban')
}
</script>
