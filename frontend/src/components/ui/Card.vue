<template>
  <div :class="cn('rounded-lg border bg-card text-card-foreground shadow-sm', typeof $attrs.class === 'string' ? $attrs.class : undefined)" v-bind="Object.fromEntries(Object.entries($attrs).filter(([k]) => k !== 'class'))">
    <slot />
  </div>
</template>

<script setup lang="ts">
import { cn } from "@/lib/utils"
</script>
