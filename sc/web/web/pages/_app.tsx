import { AppProps } from "next/app";
import Head from "next/head";
import { useState } from "react";
import { MantineProvider, ColorSchemeProvider, ColorScheme } from "@mantine/core";
import { withUrqlClient } from "next-urql";
import { SERVER_URL } from "@/utils/consts";
import "@fontsource/comfortaa";
import { CookiesProvider } from "react-cookie";

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
			<CookiesProvider>
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
			</CookiesProvider>
		</>
	);
};

export default withUrqlClient(
	() => ({
		url: SERVER_URL || "",
	}),
	{ ssr: false }
)(App);
