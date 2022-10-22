<template>
  <v-row justify="center" align="center">
    <v-card width="350px" class="my-10">
      <v-card-title class="headline">
        体重記録
      </v-card-title>
      <v-card-text>
        <p>
          体重を記録
        </p>
        <v-menu
          ref="menu"
          v-model="menu"
          :close-on-content-click="false"
          :return-value.sync="date"
          transition="scale-transition"
          offset-y
          min-width="auto"
        >
          <template #activator="{ on, attrs }">
            <v-text-field
              v-model="date"
              label="日付を選択"
              prepend-icon="mdi-calendar"
              readonly
              required
              v-bind="attrs"
              v-on="on"
            ></v-text-field>
          </template>
          <v-date-picker
            v-model="date"
            no-title
            scrollable
          >
            <v-spacer></v-spacer>
            <v-btn
              text
              color="primary"
              @click="menu = false"
            >
              キャンセル
            </v-btn>
            <v-btn
              text
              color="primary"
              @click="$refs.menu.save(date)"
            >
              OK
            </v-btn>
          </v-date-picker>
        </v-menu>
        <v-text-field
            v-model="weights"
            required
            label="体重を入力"
            suffix="kg"
            placeholder="XX.X"
            prepend-icon="mdi-human"
        ></v-text-field>
        
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary" :loading="loading" @click="record()"> 記録 </v-btn>
      </v-card-actions>
    </v-card>
    <v-snackbar
      v-model="snackbar"
    >
      {{ text }}

      <template #action="{ attrs }">
        <v-btn
          color="pink"
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          閉じる
        </v-btn>
      </template>
    </v-snackbar>
  </v-row>
</template>

<script>
import moment from 'moment';

export default {
  name: 'WeightRecordPage',
  data: () => ({
    menu: false,
    loading: false,
    snackbar: false,
    text: "",
    date: "",
    weights: ""
  }),
  methods: {
    async record() {
      this.loading = true
      const response = await this.$axios.$post('/user/weights', {
        weights: Number(this.weights),
        at: moment(this.date).toISOString()
      }).catch((e) => {
        this.snackbar = true;
        this.text = 'エラーが発生しました。'
      });
      console.log(response)
      this.loading = false;
      this.snackbar = true;
      this.text = "記録しました!"
    }
  }
}
</script>
