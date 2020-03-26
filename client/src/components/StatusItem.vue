<template>
	<div>
		<span>{{ value }}</span>
		<p>{{ description }}</p>
	</div>
</template>
<style scoped>
	span {
		font-size: 32pt;
	}
	p {
		font-size: 16pt;
	}
</style>
<script>
	export default {
		props: ['value', 'description'],
		name: 'StatusItem'
	}
</script>