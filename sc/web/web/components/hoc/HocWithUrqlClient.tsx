import { withUrqlClient } from "next-urql";
import { dedupExchange, cacheExchange, fetchExchange } from "@urql/core";
import { SERVER_URL } from "@/utils/consts";

export default function withUrqlClientDef<T>(Component: React.ComponentType<T>) {
	return withUrqlClient(ssrExchange => ({
		url: SERVER_URL,
		exchanges: [dedupExchange, cacheExchange, ssrExchange, fetchExchange],
	}))(Component);
}
