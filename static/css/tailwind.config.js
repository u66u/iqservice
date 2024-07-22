import franken from 'franken-ui/shadcn-ui/preset-quick';



/** @type {import('tailwindcss').Config} */

export default {
	presets: [franken({ theme: "neutral" })],
	
	content: [
		"../templates/**/*.{html,js}",
		"../js/**/*.js",
	 ],
	theme: {
	  extend: {
		colors: {
		  'sunset-pink': '#ff7387',
		  'light-pink': '#ffd6dd',
		  'grey-1': '#7e6a6d',
		  'pinkorange': '#DF6C4F',
		  'redish': "#fc4445",
		  'pastelolive': '#A1BE95',
		  'salmonpink': '#F98866',
		  'pastel': '#fefae0',
		  'shrek': '#606c38',
		  'lightbrown': '#bc6c25',
		  'lightblack': '#252422',
		  'peach': '#eb5e28',
		  'chillblack': '#0e0e11',
		  'darkpink': '#c13d60',
		  'lightpurple': '#cda8cd',
		}
	  },
	},
	safelist: [
		{
			pattern: /^uk-/
		}
	],
	plugins: []

	
};
