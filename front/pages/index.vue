<template>
  <div>
    <v-card width="350px" class="my-10">
      <v-card-title class="headline">
        {{ now.format('YYYY年MM月DD日') }}の体重記録
      </v-card-title>
      <v-card-text>
        <p>体重 {{ weight }}kg</br>
        食事 {{ meal }}</p>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
import moment from 'moment'

export default {
  data: () => ({
    now: moment(),
    weight: '',
    meal: '',
    loading: false
  }),
  async created () {
    this.now = moment()
    this.loading = false
    this.snackbar = false
    this.text = ''
    const response1 = await this.$axios.$get('/user/weights?at=1666364400').catch((e) => {
        this.snackbar = true;
        this.text = 'エラーが発生しました。'
    });
    const response2 = await this.$axios.$get('/user/meals?at=1666191600').catch((e) => {
        this.snackbar = true;
        this.text = 'エラーが発生しました。'
    });
      console.log(response1)
    this.weight = response1[0].weights
      this.meal = response2[0].name
      this.snackbar = true;
    },
}
</script>