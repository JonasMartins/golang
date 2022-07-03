import type { NextPage } from "next";
import dynamic from "next/dynamic";

const DynamicPage = dynamic(() => import("@/components/pages/login"), {
	ssr: false,
});

const LoginPage: NextPage = () => {
	return <DynamicPage />;
};

export default LoginPage;
