import { GRAPHQL_WS_CLIENT, SERVER_URL } from "@/utils/consts";
import { multipartFetchExchange } from "@urql/exchange-multipart-fetch";
import {
	cacheExchange,
	Client,
	createClient,
	dedupExchange,
	Exchange,
	subscriptionExchange,
} from "urql";

// import { SubscriptionClient } from "subscriptions-transport-ws";
import { createClient as createWSClient } from "graphql-ws";

export const CreateUrqlClient = (): Client => {
	let exchanges: Exchange[] | undefined = [dedupExchange, cacheExchange, multipartFetchExchange];
	if (typeof window !== "undefined") {
		const wsClient = createWSClient({
			url: GRAPHQL_WS_CLIENT,
		});
		const subsExchange = subscriptionExchange({
			forwardSubscription: operation => ({
				subscribe: sink => ({
					unsubscribe: wsClient.subscribe(operation, sink),
				}),
			}),
		});

		/*
		const subscriptionClient = new SubscriptionClient(GRAPHQL_WS_CLIENT, { reconnect: true });
		const subsExchange = subscriptionExchange({
			forwardSubscription: operation => subscriptionClient.request(operation),
		});

		*/

		exchanges.push(subsExchange);
	}

	const client = createClient({
		url: SERVER_URL,
		fetchOptions: {
			credentials: "include",
		},
		exchanges,
	});

	return client;
};
