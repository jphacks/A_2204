<template>
  <v-row justify="center" align="center">
    <v-card width="350px" class="my-10">
      <v-card-title class="headline">
        食事記録
      </v-card-title>
      <v-card-text>
        <p>
          食事を記録
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
            v-model="name"
            required
            label="食事を入力"
            placeholder=""
            prepend-icon="mdi-food-fork-drink"
        ></v-text-field>
        <v-text-field
            v-model="calorie"
            required
            label="カロリーを入力"
            suffix="kcal"
            placeholder="512"
            prepend-icon="mdi-calculator"
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
  name: 'FoodRecorcdPage',
  data: () => ({
    menu: false,
    loading: false,
    snackbar: false,
    text: "",
    date: "",
    name: "",
    calorie:""
  }),
  methods: {
    async record() {
      this.loading = true
      const response = await this.$axios.$post('/user/meals', {
        calorie: Number(this.calorie),
        name: this.name,
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
