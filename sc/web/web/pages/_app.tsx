import { AppProps } from "next/app";
import Head from "next/head";
import { useState } from "react";
import { MantineProvider, ColorSchemeProvider, ColorScheme } from "@mantine/core";
import { createClient, Provider, dedupExchange, cacheExchange } from "urql";
import { multipartFetchExchange } from "@urql/exchange-multipart-fetch";
import { SERVER_URL } from "@/utils/consts";
import { store, persistor } from "@/app";
import { Provider as ReduxProvider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";
import Loader from "@/components/layout/Loader";
import "@fontsource/comfortaa";

const client = createClient({
	url: SERVER_URL,
	fetchOptions: {
		credentials: "include",
	},
	exchanges: [dedupExchange, cacheExchange, multipartFetchExchange],
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
				<ReduxProvider store={store}>
					<PersistGate loading={<Loader />} persistor={persistor}>
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
									loader: "dots",
								}}
							>
								<Component {...pageProps} />
							</MantineProvider>
						</ColorSchemeProvider>
					</PersistGate>
				</ReduxProvider>
			</Provider>
		</>
	);
};

export default App;
