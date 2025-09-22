<template>
  <div class="min-h-screen bg-background">
    <div class="container mx-auto px-4 py-8">
      <div class="max-w-4xl mx-auto">
        <div class="mb-8">
          <h1 class="text-3xl font-bold text-foreground">设置</h1>
          <p class="text-muted-foreground mt-2">管理您的应用偏好和账户设置</p>
        </div>

        <div class="space-y-8">
          <!-- 外观设置 -->
          <Card class="p-6">
            <CardHeader>
              <h3 class="text-lg font-semibold">外观设置</h3>
              <p class="text-sm text-muted-foreground">自定义应用的外观和主题</p>
            </CardHeader>
            <CardContent class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-foreground">深色模式</label>
                  <p class="text-sm text-muted-foreground">启用深色主题</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input
                    v-model="settings.darkMode"
                    type="checkbox"
                    class="sr-only peer"
                  />
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
                </label>
              </div>
            </CardContent>
          </Card>

          <!-- 通知设置 -->
          <Card class="p-6">
            <CardHeader>
              <h3 class="text-lg font-semibold">通知设置</h3>
              <p class="text-sm text-muted-foreground">管理您接收的通知类型</p>
            </CardHeader>
            <CardContent class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-foreground">任务提醒</label>
                  <p class="text-sm text-muted-foreground">任务截止日期提醒</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input
                    v-model="settings.taskNotifications"
                    type="checkbox"
                    class="sr-only peer"
                  />
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
                </label>
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-foreground">邮件通知</label>
                  <p class="text-sm text-muted-foreground">通过邮件接收重要更新</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input
                    v-model="settings.emailNotifications"
                    type="checkbox"
                    class="sr-only peer"
                  />
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
                </label>
              </div>
            </CardContent>
          </Card>

          <!-- 账户设置 -->
          <Card class="p-6">
            <CardHeader>
              <h3 class="text-lg font-semibold">账户设置</h3>
              <p class="text-sm text-muted-foreground">管理您的账户安全</p>
            </CardHeader>
            <CardContent class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-foreground">修改密码</label>
                  <p class="text-sm text-muted-foreground">更新您的登录密码</p>
                </div>
                <Button variant="outline" size="sm">
                  修改密码
                </Button>
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <label class="text-sm font-medium text-foreground">删除账户</label>
                  <p class="text-sm text-muted-foreground">永久删除您的账户和所有数据</p>
                </div>
                <Button variant="destructive" size="sm">
                  删除账户
                </Button>
              </div>
            </CardContent>
          </Card>

          <!-- 保存按钮 -->
          <div class="flex justify-end">
            <Button @click="saveSettings" size="lg">
              保存设置
            </Button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Card from '@/components/ui/Card.vue'
import CardHeader from '@/components/ui/CardHeader.vue'
import CardContent from '@/components/ui/CardContent.vue'
import Button from '@/components/ui/Button.vue'

const settings = ref({
  darkMode: false,
  taskNotifications: true,
  emailNotifications: false
})

const saveSettings = () => {
  // 这里可以添加保存设置的逻辑
  console.log('Settings saved:', settings.value)
}

onMounted(() => {
  // 这里可以从localStorage或其他地方加载设置
})
</script>
