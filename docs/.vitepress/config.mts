import {defineConfig} from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
	title: 'sesamy-cli',
	description: 'CLI to keep you sane while working with Google Tag Manager.',
	lang: 'en-US',
	cleanUrls: true,
	lastUpdated: true,
	appearance: 'dark',
	ignoreDeadLinks: false,
	base: '/sesamy-cli/',
	sitemap: {
		hostname: 'https://foomo.github.io/sesamy-cli',
	},
	themeConfig: {
		// https://vitepress.dev/reference/default-theme-config
		logo: '/logo.png',
		outline: [2, 4],
		nav: [
			{text: 'Guide', link: '/guide/introduction', activeMatch: '/guide/'},
			{text: 'Commands', link: '/commands/', activeMatch: '/commands/'},
			{text: 'Providers', link: '/providers/', activeMatch: '/providers/'},
			{text: 'Reference', link: '/reference/configuration', activeMatch: '/reference/'},
		],
		sidebar: [
			{
				text: 'Guide',
				items: [
					{text: 'Introduction', link: '/guide/introduction'},
					{text: 'Installation', link: '/guide/installation'},
					{text: 'Quick start', link: '/guide/quick-start'},
					{text: 'Configuration', link: '/guide/configuration'},
					{text: 'Workflow', link: '/guide/workflow'},
					{text: 'sesamy-go integration', link: '/guide/sesamy-go'},
				],
			},
			{
				text: 'Commands',
				collapsed: true,
				items: [
					{text: 'Overview', link: '/commands/'},
					{text: 'provision', link: '/commands/provision'},
					{text: 'list', link: '/commands/list'},
					{text: 'diff', link: '/commands/diff'},
					{text: 'typescript', link: '/commands/typescript'},
					{text: 'tags', link: '/commands/tags'},
					{text: 'open', link: '/commands/open'},
					{text: 'config', link: '/commands/config'},
				],
			},
			{
				text: 'Providers',
				collapsed: true,
				items: [
					{text: 'Overview', link: '/providers/'},
					{text: 'Google Tag', link: '/providers/google-tag'},
					{text: 'Google Analytics', link: '/providers/google-analytics'},
					{text: 'Google Ads', link: '/providers/google-ads'},
					{text: 'Conversion Linker', link: '/providers/conversion-linker'},
					{text: 'Facebook', link: '/providers/facebook'},
					{text: 'Pinterest', link: '/providers/pinterest'},
					{text: 'Microsoft Ads', link: '/providers/microsoft-ads'},
					{text: 'Criteo', link: '/providers/criteo'},
					{text: 'Emarsys', link: '/providers/emarsys'},
					{text: 'Mixpanel', link: '/providers/mixpanel'},
					{text: 'Umami', link: '/providers/umami'},
					{text: 'Tracify', link: '/providers/tracify'},
					{text: 'Hotjar', link: '/providers/hotjar'},
					{text: 'Cookiebot', link: '/providers/cookiebot'},
				],
			},
			{
				text: 'Reference',
				collapsed: true,
				items: [
					{text: 'Configuration', link: '/reference/configuration'},
					{text: 'JSON schema', link: '/reference/schema'},
					{text: 'Google API setup', link: '/reference/google-api'},
					{text: 'CLI reference', link: '/reference/cli/sesamy'},
				],
			},
			{
				text: 'Contributing',
				collapsed: true,
				items: [
					{text: 'Guideline', link: '/CONTRIBUTING.md'},
					{text: 'Code of conduct', link: '/CODE_OF_CONDUCT.md'},
					{text: 'Security guidelines', link: '/SECURITY.md'},
				],
			},
		],
		socialLinks: [
			{icon: 'github', link: 'https://github.com/foomo/sesamy-cli'},
		],
		editLink: {
			pattern: 'https://github.com/foomo/sesamy-cli/edit/main/docs/:path',
		},
		search: {
			provider: 'local',
		},
		footer: {
			message: 'Made with ♥ <a href="https://www.foomo.org">foomo</a> by <a href="https://www.bestbytes.com">bestbytes</a>',
		},
	},
	markdown: {
		// https://github.com/vuejs/vitepress/discussions/3724
		theme: {
			light: 'catppuccin-latte',
			dark: 'catppuccin-frappe',
		},
	},
	head: [
		['meta', {name: 'theme-color', content: '#ffffff'}],
		['link', {rel: 'icon', href: '/sesamy-cli/logo.png'}],
		['meta', {name: 'author', content: 'foomo by bestbytes'}],
		// OpenGraph
		['meta', {property: 'og:title', content: 'foomo/sesamy-cli'}],
		[
			'meta',
			{
				property: 'og:image',
				content: 'https://github.com/foomo/sesamy-cli/blob/main/docs/public/logo.png?raw=true',
			},
		],
		[
			'meta',
			{
				property: 'og:description',
				content: 'CLI to provision, diff, and inspect Google Tag Manager server-side and web containers across analytics, advertising, and consent providers.',
			},
		],
		['meta', {name: 'twitter:card', content: 'summary_large_image'}],
		[
			'meta',
			{
				name: 'twitter:image',
				content: 'https://github.com/foomo/sesamy-cli/blob/main/docs/public/logo.png?raw=true',
			},
		],
		['meta', {name: 'viewport', content: 'width=device-width, initial-scale=1.0, viewport-fit=cover'}],
	],
})
