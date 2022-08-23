import type { NextPage } from "next";
import dynamic from "next/dynamic";

const DynamicPage = dynamic(() => import("@/components/pages/appShell"), {
	ssr: false,
});

const AboutPage: NextPage = () => {
	return <DynamicPage />;
};

export default AboutPage;
