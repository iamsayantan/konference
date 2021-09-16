<template>
  <div>
    Room Details
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
@Component({
  name: 'RoomPage',
  middleware: 'auth'
})
export default class RoomPage extends Vue {
  mounted(): void {
    console.log('Code', this.$route.params.code)
  }
}
</script>

<style scoped>

</style>
