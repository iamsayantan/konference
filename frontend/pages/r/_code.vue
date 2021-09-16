<template>
  <v-container fluid fill-height>
    <v-row justify="center" align="center">
      <v-col>
        <v-row justify="center" align="center">
          <v-row justify="center" align="center" style="background-color: #333; min-width: 100%; min-height: 100%"></v-row>
        </v-row>
      </v-col>
      <v-col>
        Join Button
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  name: 'RoomPage',
  middleware: 'auth'
})
export default class RoomPage extends Vue {
  hasVideoPermission: boolean = false;
  hasAudioPermission: boolean = false;

  meetingCode: string = '';

  mounted(): void {
    this.meetingCode = this.$route.params.code;
    this.checkCurrentPermissions();
  }

  async checkCurrentPermissions() {
    const devices = await navigator.mediaDevices.enumerateDevices();
    devices.forEach(device => {
      if (device.kind === 'audioinput' && device.label) {
        this.hasAudioPermission = true;
      }

      if (device.kind === 'videoinput' && device.label) {
        this.hasVideoPermission = true;
      }
    });
  }
}

</script>

<style scoped>

</style>
