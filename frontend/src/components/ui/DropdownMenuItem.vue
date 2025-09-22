<template>
  <a
    :class="cn(
      'block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:text-gray-200 dark:hover:bg-gray-700 cursor-pointer',
      ($attrs as any).class
    )"
    v-bind="$attrs"
    @click="$emit('click')"
  >
    <slot />
  </a>
</template>

<script setup lang="ts">
import { cn } from "@/lib/utils"

defineEmits<{
  click: []
}>()
</script>