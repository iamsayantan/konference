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
            <v-btn color="teal" dark large block outlined>Create Room</v-btn>
          </div>
        </v-layout>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  components: {
    Logo,
    VuetifyLogo,
  },
}
</script>
