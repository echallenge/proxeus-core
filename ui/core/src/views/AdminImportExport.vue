<template>
<div style="height:100%;">
  <vue-headful :title="$t('Data title', 'Proxeus - Data')"/>
  <top-nav :title="$t('Data')" />
  <import-export :chooser="true" :exportTypes="[]"/>
</div>
</template>

<script>
import TopNav from '../components/layout/TopNav'
import importExport from '../components/ImportExport'

export default {
  name: 'AdminImportExport',
  components: {
    TopNav,
    importExport
  }
}
</script>

<style lang="scss">
</style>
