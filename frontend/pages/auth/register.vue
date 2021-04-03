<template>
  <v-layout
      column
      justify-center
      align-center
  >
    <div style="margin-top: 20px; width: 90%">
      <v-text-field
        v-model="registrationData.first_name"
        outlined
        name="first_name"
        label="First Name"
        required
      ></v-text-field>
      <v-text-field
          v-model="registrationData.last_name"
          outlined
          name="last_name"
          label="Last Name"
          required
      ></v-text-field>
      <v-text-field
        v-model="registrationData.email"
        type="email"
        outlined
        name="email"
        label="Email"
        required
      ></v-text-field>
      <v-text-field
          v-model="registrationData.password"
          outlined
          type="password"
          name="password"
          label="Password"
          required
      ></v-text-field>
      <v-btn :loading="loading" color="teal" dark large block @click="handleRegistration">Register</v-btn>
    </div>
    <div style="margin-top: 15px; color: teal">
      <v-btn text small color="teal" class="subtitle-1" @click.prevent="handleRedirectToLogin">Login</v-btn>
    </div>
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import {RegistrationRequest, RegistrationResponse} from "~/types/auth";

@Component({
  name: 'RegistrationPage',
  layout: 'auth'
})
export default class RegistrationPage extends Vue {
  loading: boolean = false;
  registrationData = <RegistrationRequest>{};
  returnUrl: string|null = null;

  mounted(): void {
    console.log('register mounted');
    this.returnUrl = this.$route.query['return_url'] ? this.$route.query['return_url'].toString() : null;
  }

  async handleRegistration() {
    const registrationData = this.registrationData;

    if (!registrationData.first_name) {
      this.$toast.show('First name is required');
      return;
    }

    if (!registrationData.last_name) {
      this.$toast.show('Last name is required');
      return;
    }

    if (!registrationData.email) {
      this.$toast.show('Email is required');
      return;
    }

    if (!registrationData.password) {
      this.$toast.show('Password is required');
      return;
    }


    this.loading = true;

    try {
      const registrationResponse = await this.$axios.post<RegistrationResponse>('/users/v1/register', this.registrationData);
      this.$toast.show(registrationResponse.data.message);
      this.handleRedirectToLogin()
    } catch (e) {
      if (e.response && e.response.data) {
        this.$toast.error(e.response.data.message);
      }
    }

    this.loading = false;
  }

  async handleRedirectToLogin() {
    await this.$router.push(`/auth/login?return_url=${this.returnUrl ? this.returnUrl : ''}`)
  }

}
</script>

<style scoped>

</style>
