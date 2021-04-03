<template>
  <v-layout
    column
    justify-center
    align-center
  >
    <div style="margin-top: 20px; width: 90%">
      <v-form @submit="handleLogin">
        <v-text-field
          v-model="loginRequest.email"
          outlined
          type="email"
          name="email"
          label="Email"
          required
        ></v-text-field>
        <v-text-field
          v-model="loginRequest.password"
          outlined
          type="password"
          name="password"
          label="Password"
          required
        ></v-text-field>
        <v-btn :loading="loading" color="teal" type="submit" dark large block>Login</v-btn>
      </v-form>
    </div>
    <div style="margin-top: 15px; color: teal">
      <v-btn text small color="teal" class="subtitle-1" @click.prevent="handleRedirectToRegister">Register</v-btn>
    </div>
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { LoginRequest } from "~/types/auth";

@Component({
  name: 'LoginPage',
  layout: 'auth',
  auth: 'guest',
  asyncData({ redirect, $auth }) {
    if ($auth.loggedIn) {
      return redirect('/')
    }
  }
})
export default class LoginPage extends Vue {
  loading: boolean = false;
  loginRequest = <LoginRequest>{};
  returnUrl: string|null = null;

  mounted(): void {
    this.returnUrl = this.$route.query['return_url'] ? this.$route.query['return_url'].toString() : null;
  }

  async handleLogin() {
    if (!this.loginRequest.email) {
      this.$toast.show('Email is required');
      return;
    }

    if (!this.loginRequest.password) {
      this.$toast.show('Password is required');
      return;
    }

    this.loading = true;
    try {
      await this.$auth.loginWith('local', {data: this.loginRequest});
    } catch (e) {
      if (e.response && e.response.data) {
        const errorResponse = e.response.data;
        this.$toast.error(errorResponse.message);
      }

      console.log('Error', e.response);
    }
    this.loading = false;
  }

  async handleRedirectToRegister() {
    await this.$router.push(`/auth/register?return_url=${this.returnUrl ? this.returnUrl : ''}`);
  }
}
</script>

<style scoped>

</style>
