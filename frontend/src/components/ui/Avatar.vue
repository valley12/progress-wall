<template>
  <span
    :class="cn(
      'relative flex h-10 w-10 shrink-0 overflow-hidden rounded-full',
      typeof $attrs.class === 'string' ? $attrs.class : undefined
    )"
    v-bind="Object.fromEntries(Object.entries($attrs).filter(([k]) => k !== 'class'))"
  >
    <img
      v-if="src"
      :src="src"
      :alt="alt"
      class="aspect-square h-full w-full object-cover"
    />
    <div
      v-else
      class="flex h-full w-full items-center justify-center rounded-full bg-muted"
    >
      <span class="text-sm font-medium">{{ fallback }}</span>
    </div>
  </span>
</template>

<script setup lang="ts">
import { cn } from "@/lib/utils"

interface Props {
  src?: string
  alt?: string
  fallback?: string
}

defineProps<Props>()
</script>
