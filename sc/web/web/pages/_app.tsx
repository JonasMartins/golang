import { persistor, store } from "@/app";
import Loader from "@/components/layout/Loader";
import { ColorScheme, ColorSchemeProvider, MantineProvider } from "@mantine/core";
import { AppProps } from "next/app";
import Head from "next/head";
import { useState } from "react";
import { Provider as ReduxProvider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";
import { Provider } from "urql";
import { CreateUrqlClient } from "@/main/config/clients";
import "@fontsource/comfortaa";

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
			<Provider value={CreateUrqlClient()}>
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
