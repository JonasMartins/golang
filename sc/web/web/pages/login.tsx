import type { NextPage } from "next";
import dynamic from "next/dynamic";
import { run } from "@/utils/aux/test";

const DynamicPage = dynamic(() => import("@/components/pages/login"), {
	ssr: false,
});

const LoginPage: NextPage = () => {
	run();
	return <DynamicPage />;
};

export default LoginPage;
