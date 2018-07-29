<template lang="html">
  <div class="page page--dashboard">
    <div class="container dashboard">
      <div class="dashboard__latest">
        {{ latest }}{{ point.unit }}
      </div>
      <div class="dashboard__average">
        Average: {{ avg }}{{ point.unit }}
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  async mounted () {
    try {
      await this.$store.dispatch('getPoints')
    } catch (err) {
      console.log(err)
    }
  },
  computed: {
    ...mapGetters(['points']),
    point () {
      if (!this.points.length) return {}
      const point = this.points.find(item => item.unit === '°C' && item.measurement === 'temp')
      if (!point) return {}
      return point
    },
    avg () {
      if (!this.point.data) return 0
      let sum = 0
      this.point.data.forEach(item => {
        sum += item.value
      })
      return (sum / this.point.data.length).toFixed(1)
    },
    latest () {
      if (!this.point.data) return 0
      let result = { datetime: 0 }
      this.point.data.forEach(item => {
        if (new Date(item.datetime) >= new Date(result.datetime)) {
          result = { ...item }
        }
      })
      return (result.value).toFixed(1)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/scss/main.scss';

.page {
  &--dashboard {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
}
.dashboard {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #fff;
  padding: 3rem 0;
  box-shadow: 0 .5rem 1rem rgba(0, 0, 0, 0.1);
  border-radius: 2rem;
  font-size: 1rem;

  &__latest {
    font-size: 13em;
  }
  &__average {
    font-size: 5em;
  }

  @media (max-width: 500px) {
    font-size: 0.7rem;
    width: 100%;
    margin-left: 1rem;
    margin-right: 1rem;
  }
}
</style>
