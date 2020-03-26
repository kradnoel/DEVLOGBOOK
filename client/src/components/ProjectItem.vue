<template>
	<div>
		<h2>{{ title }}</h2>
		<p>{{ description }}</p>
		<big>{{ timespent }} hours</big>
	</div>
</template>
<style scoped>
	p {
		font-size: 14pt;
	}
</style>
<script>
	export default {
		props: ['title', 'description', 'timespent'],
		name: 'ProjectItem'
	}
</script>