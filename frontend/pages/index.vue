<template>
  <v-container>
    <v-layout
        column
    >
      <v-row no-gutters>
        <div>
          <span class="display-3" style="color: #069587">Create Your Room</span>
          <div style="margin-top: 20px">
            <span class="subtitle-1">Invite you friends to your room to have a realtime audio video conferencing.</span>
          </div>
        </div>
        <v-layout
            column
            justify-center
            align-center
        >
          <div style="margin-top: 10px">
            <v-btn :loading="loading" color="teal" dark large block outlined @click="createRoom">Create Room</v-btn>
          </div>
        </v-layout>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { Room } from '~/types/room'

@Component({
  name: 'IndexPage',
  middleware: 'auth'
})
export default class IndexPage extends Vue{
  loading: boolean = false

  async createRoom() {
    this.loading = true
    try {
      const resp = await this.$axios.post<Room>('/rooms/v1')
      console.log('RoomDetails', resp.data)
    } catch (e) {
      console.error(e)
    }
    this.loading = false
  }
}
</script>
