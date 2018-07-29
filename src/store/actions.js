export default {
  async getPoints ({ commit }) {
    const response = await fetch('/api/temp')
      .then(response => {
        if (!response.ok) {
          throw new Error(response.statusText)
        }
        return response.json()
      })

    commit('SET_POINTS', response)
  }
}
