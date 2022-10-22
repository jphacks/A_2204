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
    weight: 55.0,
    loading: false
  }),
  created () {
    this.now = moment()
    this.weight = 55.0
    this.meal = "ラーメン"
    this.loading = false
    this.snackbar = false
    this.text = ''
  },
}
</script>