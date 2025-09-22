<template>
  <div class="min-h-screen bg-background">
    <div class="container mx-auto px-4 py-8">
      <div class="max-w-4xl mx-auto">
        <div class="mb-8">
          <h1 class="text-3xl font-bold text-foreground">个人资料</h1>
          <p class="text-muted-foreground mt-2">管理您的个人信息和账户设置</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <!-- 个人信息卡片 -->
          <div class="lg:col-span-1">
            <Card class="p-6">
              <div class="text-center">
                <Avatar
                  :src="userStore.currentUser?.avatar"
                  :fallback="userStore.currentUser?.name?.charAt(0) || 'U'"
                  class="h-24 w-24 mx-auto mb-4"
                />
                <h2 class="text-xl font-semibold text-foreground">
                  {{ userStore.currentUser?.name }}
                </h2>
                <p class="text-muted-foreground">
                  {{ userStore.currentUser?.email }}
                </p>
                <div class="mt-4">
                  <span
                    :class="[
                      'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                      userStore.currentUser?.role === 'admin'
                        ? 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200'
                        : 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200'
                    ]"
                  >
                    {{ userStore.currentUser?.role === 'admin' ? '管理员' : '普通用户' }}
                  </span>
                </div>
              </div>
            </Card>
          </div>

          <!-- 详细信息 -->
          <div class="lg:col-span-2">
            <Card class="p-6">
              <CardHeader>
                <h3 class="text-lg font-semibold">基本信息</h3>
              </CardHeader>
              <CardContent class="space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <label class="block text-sm font-medium text-foreground mb-2">
                      姓名
                    </label>
                    <input
                      v-model="profileForm.name"
                      type="text"
                      class="w-full px-3 py-2 border border-input rounded-md bg-background text-foreground focus:outline-none focus:ring-2 focus:ring-ring"
                    />
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-foreground mb-2">
                      邮箱
                    </label>
                    <input
                      v-model="profileForm.email"
                      type="email"
                      class="w-full px-3 py-2 border border-input rounded-md bg-background text-foreground focus:outline-none focus:ring-2 focus:ring-ring"
                    />
                  </div>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-foreground mb-2">
                    个人简介
                  </label>
                  <textarea
                    v-model="profileForm.bio"
                    rows="3"
                    class="w-full px-3 py-2 border border-input rounded-md bg-background text-foreground focus:outline-none focus:ring-2 focus:ring-ring"
                    placeholder="介绍一下自己..."
                  ></textarea>
                </div>

                <div class="flex justify-end space-x-4">
                  <Button @click="resetForm" variant="outline">
                    重置
                  </Button>
                  <Button @click="saveProfile">
                    保存更改
                  </Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import Card from '@/components/ui/Card.vue'
import CardHeader from '@/components/ui/CardHeader.vue'
import CardContent from '@/components/ui/CardContent.vue'
import Button from '@/components/ui/Button.vue'
import Avatar from '@/components/ui/Avatar.vue'

const userStore = useUserStore()

const profileForm = ref({
  name: '',
  email: '',
  bio: ''
})

const resetForm = () => {
  if (userStore.currentUser) {
    profileForm.value = {
      name: userStore.currentUser.name,
      email: userStore.currentUser.email,
      bio: ''
    }
  }
}

const saveProfile = () => {
  userStore.updateProfile({
    name: profileForm.value.name,
    email: profileForm.value.email
  })
  // 这里可以添加保存成功的提示
}

onMounted(() => {
  resetForm()
})
</script>
