<template>
	<div>
		<span>Task</span>
		<span>Id: {{ id }}</span>
		<span>Tid: {{ tid }}</span>
	</div>
</template>
<style scoped>
</style>
<script>
	import axios from "axios"

	export default {
		data() {
			return {
				id: '',
				tid: '',
				task: '',
				values: {},
				error: false
			}
		},
		created () {
			axios.get('http://localhost:8080/tasks/' + tid)
		}
	}
</script>