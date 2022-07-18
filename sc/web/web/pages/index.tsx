import type { NextPage } from "next";
import dynamic from "next/dynamic";

const DynamicPage = dynamic(() => import("@/components/pages/dashboard"), {
	ssr: false,
});

const HomePage: NextPage = () => {
	return <DynamicPage />;
};

export default HomePage;
