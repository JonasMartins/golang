import { AppProps } from "next/app";
import Head from "next/head";
import { useState } from "react";
import { MantineProvider, ColorSchemeProvider, ColorScheme } from "@mantine/core";
import { createClient, Provider } from "urql";
import { SERVER_URL } from "@/utils/consts";
import "@fontsource/comfortaa";

const client = createClient({
	url: SERVER_URL,
	fetchOptions: {
		credentials: "include",
	},
});

const App = (props: AppProps) => {
	const { Component, pageProps } = props;
	const [colorScheme, setColorScheme] = useState<ColorScheme>("light");
	const toggleColorScheme = (value?: ColorScheme) =>
		setColorScheme(value || (colorScheme === "dark" ? "light" : "dark"));

	return (
		<>
			<Head>
				<title>SC</title>
				<meta
					name="viewport"
					content="minimum-scale=1, initial-scale=1, width=device-width"
				/>
			</Head>
			<Provider value={client}>
				<ColorSchemeProvider
					colorScheme={colorScheme}
					toggleColorScheme={toggleColorScheme}
				>
					<MantineProvider
						withGlobalStyles
						withNormalizeCSS
						theme={{
							/** Put your mantine theme override here */
							fontFamily: "Verdana, sans-serif",
							fontFamilyMonospace: "Monaco, Courier, monospace",
							headings: { fontFamily: "Comfortaa, cursive" },
							colorScheme,
							breakpoints: {
								xs: 500,
								sm: 800,
								md: 1000,
								lg: 1200,
								xl: 1400,
							},
						}}
					>
						<Component {...pageProps} />
					</MantineProvider>
				</ColorSchemeProvider>
			</Provider>
		</>
	);
};

export default App;
