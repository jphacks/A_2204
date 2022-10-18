<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar :clipped-left="clipped" fixed app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-spacer />
      <v-btn icon>
        <v-icon>mdi-chevron-left</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-chevron-right</v-icon>
      </v-btn>
      <v-spacer />
      <v-btn icon to="/user">
        <v-icon>mdi-account</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
    <FooterVue />
  </v-app>
</template>

<script>
export default {
  name: 'DefaultLayout',
  data() {
    return {
      clipped: true,
      drawer: true,
      items: [
        {
          icon: 'mdi-weight',
          title: '体重',
          to: '/user-weight',
        },
        {
          icon: 'mdi-paw',
          title: 'キャラクター',
          to: '/character',
        },
      ],
    }
  },
}
</script>
